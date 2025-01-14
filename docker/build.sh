#!/bin/bash 

docker build -t go-debug .
docker tag godebug immnan/go-debug:latest
docker push immnan/godebug:latest
docker images
