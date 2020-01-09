#!/bin/bash
[[ -z "${GITHUB_SHA}" ]] && Tag='local' || Tag="${GITHUB_SHA::4}"
[[ -z "${docker_username}" ]] && DockerRepo='' || DockerRepo="${docker_username}/"
docker build -t "${DockerRepo}name:latest" -t "${DockerRepo}name:1.0-$Tag" ./app