#!/bin/bash

# Define array of container names
#containers=("relayer1" "relayer2" "relayer3" "relayer4" "relayer5" "relayer6" "relayer7" "relayer8" "relayer9" "relayer10")
containers=("relayer1-norm" "relayer2-tor" "relayer3-proxy" "relayer4-proxy")
# Loop through containers
for container in "${containers[@]}"
do
  # Enter the container
  docker exec -it $container bash -c "cat /log-step.txt" >> ./step-log/log-step-$container.txt
done
