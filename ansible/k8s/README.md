Ansible/K8s
===========

## Description
This project defines ansible kubernetes cluster automation.  The goal of this project is to replace the current
automation for the kubernetes cluster operated locally and to make the solution reasonably portable to other
environments over time.


## Usage:
1. Setup linux hosts
2. Ensure key-based ssh for each host
3. Ensure hostnames can be resolved locally (DHCP-managed DNS)
4. Run ansible/core
5. Run ansible/k8s
