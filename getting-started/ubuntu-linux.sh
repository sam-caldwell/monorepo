#!/bin/bash -e

command -v git ||{
  echo "git must be installed first"
  echo "run sudo apt-get install git -y"
  echo "Then clone this repo...how did you get this far?"
  exit 1
}

[[ "$(whoami)" == "root" ]] && {
  echo "cannot run as root"
  exit 1
}

create_directories(){
  mkdir -p "${HOME}/go/bin"
  mkdir -p "${HOME}/.local/bin"
}
append_bashrc_paths(){
  # shellcheck disable=SC2016
  echo 'export PATH=${PATH}:${HOME}/.local/bin/:${HOME}/go/bin' > "${HOME}/.bashrc"
}
update_apt(){
  sudo apt-get update -y --fix-missing
  sudo apt-get install build-essential curl nodejs npm python3 python3-pip shellcheck yamllint flake8 -y
}
install_virtualenv(){
  command -v virtualenv || pip3 install virtualenv
  source "${HOME}/.bashrc"
}
install_nvm(){
  curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
  LATEST_NODE_VERSION="$(nvm ls-remote | tail -n 1 | awk '{print $1}')"
  nvm install "${LATEST_NODE_VERSION}"
}

create_directories
append_bashrc_paths
update_apt
# shellcheck disable=SC1090
source "${HOME}/.bashrc"
flake8 --version
yamllint --version
shellcheck --version
# shellcheck disable=SC1090
source "${HOME}/.bashrc"
install_virtualenv
npm --version
node --version
virtualenv --version
npm install
# shellcheck disable=SC1090
source "${HOME}/.bashrc"
install_nvm
# shellcheck disable=SC1090
#source "${HOME}/.bashrc"
#npm --version
#node --version
#command -v snyk || npm install snyk
#pip3 install -r ./requirements.txt
#go mod tidy
#go install honnef.co/go/tools/cmd/staticcheck@latest

