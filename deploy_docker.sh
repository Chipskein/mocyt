#!/bin/bash
DOCKER_USERNAME=chipskein
docker build -t $DOCKER_USERNAME/mocyt .
docker push $DOCKER_USERNAME/mocyt
