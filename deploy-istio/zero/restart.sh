#! /bin/bash
services='email domain user-v1 user-v2 user-v3 gateway'

for service in $services
do
  kubectl rollout restart deploy -n zero $service
done
