#!/bin/sh

docker build -t ersaraven/todo:latest .
docker image push ersaraven/todo:latest