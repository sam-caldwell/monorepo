#!/bin/bash -e

remote="$1"
url="$2"

echo "\033[34m>starting $0\033[0m"
make security || {
  echo "\033[31m>Lint/security check failed $0\033[0m"
  exit 1
}
echo "\033[32m>ok $0\033[0m"
exit 0
