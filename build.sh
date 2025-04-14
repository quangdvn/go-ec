#!/bin/bash

# Docker build

make docker_build

echo "127.0.0.1 go-ec.com" >> /etc/hosts

