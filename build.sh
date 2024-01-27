#! /bin/bash

harborurl="harbor.davidops.info"
services='gateway user domain email'

for service in $services
do
  docker build -f gateway/Dockerfile -t $harborurl/zero/$service:latest .
  docker push $harborurl/zero/$service:latest
done







