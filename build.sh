#!/bin/bash

docker rmi tron-tools:v1.0

docker build . -t tron-tools:v1.0
