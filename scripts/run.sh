#!/usr/bin/env bash

docker rm -fv br-parser

docker run -it -p 8888:8080 --name br-parser br-parser-base