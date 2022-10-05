#!/bin/bash

docker stop rest_api
docker rm rest_api
docker rmi rest-api

docker build --tag rest-api --file ./Dockerfile ..

