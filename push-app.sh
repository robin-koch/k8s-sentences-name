#!/bin/bash

docker push "$docker_username/name:1.0-${GITHUB_SHA::4}" 
docker push "$docker_username/name:latest" &
wait
