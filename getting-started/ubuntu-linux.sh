#!/bin/bash -e

command -v git ||{
  echo "git must be installed first"
  echo "run sudo apt-get install git -y"
  echo "Then clone this repo...how did you get this far?"
  exit 1
}

sudo apt-get update -y --fix-missing

sudo apt-get update -y --fix-missing

sudo apt-get install build-essential curl nodejs npm python3 python3-pip shellcheck yamllint flake8 -y

flake8 --version
yamllint --version
shellcheck --version

# shellcheck disable=SC2016
echo 'export PATH=${PATH}:${HOME}/.local/bin' > ${HOME}/.bashrc

# shellcheck disable=SC1090
source "${HOME}/.bashrc"

pip3 install virtualenv
source "${HOME}/.bashrc"
virtualenv --version

npm --version
node --version

npm install

curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash

# shellcheck disable=SC1090
source "${HOME}/.bashrc"

nvm install "$(nvm ls-remote | tail -n 1 | awk '{print $1}')"

npm install npm

npm --version
node --version

npm install -g snyk