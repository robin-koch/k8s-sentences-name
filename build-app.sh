#!/bin/bash

docker build -t "$docker_username/name:latest" -t "$docker_username/name:1.0-${GITHUB_SHA::4}" ./app