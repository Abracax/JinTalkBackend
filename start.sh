#!/bin/bash

docker build -t golang:latest .
docker run -it -p 80:2333 -d golang:latest go run main.go