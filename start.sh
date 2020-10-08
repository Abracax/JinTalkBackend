#!/bin/bash

docker build -t jintalk:latest .
docker run -it -p 80:2333 -d jintalk:latest go run main.go