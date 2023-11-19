#!/bin/bash -e

for i in $(seq 1 "${1}"); do
  echo "start segment $1"
  find_collisions -Segment "${i}" -Seed 10000000 &
done
