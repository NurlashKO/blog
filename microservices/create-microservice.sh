#!/bin/bash

usage() {
  echo "Usage: $(basename $0) [--gen-resources] <your mcs name>

This script creates:
- \"scratch\" mcs with specified <your mcs name> in microservices folder.
- Github Actions workflow which builds and pushes your mcs to a Container Registry

Review locally, then push to github for CI pipelines to apply your changes.
"
}

create_microservice() {
  local mcs=$1
  local mcs_dir="$SCRIPT_DIR/$mcs"
  # If git tree is clean then it is okay to commit.
  local SHOULD_GIT_COMMIT="true"
  git diff --quiet || SHOULD_GIT_COMMIT="false"

  if [ -d "$mcs_dir" ]
  then
    echo "Directory $mcs_dir already exist. Skipping creation."
    exit 1
  fi
  # Copy $scratch with <your microservice name>
  cp -r "$SCRIPT_DIR/$SCRATCH_MCS" "$mcs_dir"
  # Replace all entries of $scratch with <your microservice name>
  grep -rl "$SCRATCH_MCS" "$mcs_dir" | xargs sed -i "" -e "s/$SCRATCH_MCS/$mcs/g"
  # Move CI workflow config to .github folder
  mv "${mcs_dir}/${SCRATCH_MCS}-ci.yml" "${SCRIPT_DIR}/../.github/workflows/${mcs}-ci.yml"
  # Prompt description
  read -p "Write 3-5 word description: " description
  # Add status badge
  ex "$SCRIPT_DIR/../README.md" <<eof
3 insert
[![$mcs CI](https://github.com/NurlashKO/blog/actions/workflows/$mcs-ci.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/$mcs-ci.yml)|$description
.
xit
eof
  # Add to git tracking
  git add "$mcs_dir" "${SCRIPT_DIR}/../.github/workflows/${mcs}-ci.yml" "$SCRIPT_DIR/../README.md" && \
  [[ $SHOULD_GIT_COMMIT == "true" ]] && git commit -m "Added $mcs microservice"
}


SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]:-$0}"; )" &> /dev/null && pwd 2> /dev/null; )";
SCRATCH_MCS="_scratch"

if (($# == 1)); then
  create_microservice $1
  exit 0
else
  usage
  exit 1
fi
