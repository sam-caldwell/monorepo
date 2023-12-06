FROM opsys/ubuntu:latest AS base

RUN apt-get update -y --fix-missing && \
    apt-get install python3-full python3-pip -y && \
    pip3 install pip --break-system-packages
