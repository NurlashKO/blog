#!/bin/bash

PROVISIONER_MARK="__provisioner-managed__"

DOCKER_DEFAULT_ARGS() {
  name="$1";
  echo "-detach --log-driver json-file --log-opt tag=\"{{.ImageName}}|{{.Name}}|{{.ImageFullID}}|{{.FullID}}\" --network internal --hostname ${name} --name ${name} --label ${PROVISIONER_MARK}"
}

