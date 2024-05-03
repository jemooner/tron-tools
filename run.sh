#!/bin/bash

docker rm -f tron-tools
docker run --name tron-tools -itd --restart=unless-stopped -v /etc/localtime:/etc/localtime -v /etc/timezone:/etc/timezone -v $(pwd)/project/:/data -p 9091:9091 tron-tools:v1.0

docker logs -f tron-tools
