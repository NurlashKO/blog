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
  # Add to git tracking
  git add "$mcs_dir" "${SCRIPT_DIR}/../.github/workflows/${mcs}-ci.yml"
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
