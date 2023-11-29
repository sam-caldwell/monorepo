FROM opsys/ubuntu:latest AS base

RUN apt-get update -y --fix-missing && \
    apt-get upgrade -y

FROM base AS java

RUN apt-get update -y && \
    apt-get install curl wget xz-utils -y

WORKDIR /tmp
RUN wget https://download.java.net/java/GA/jdk21.0.1/415e3f918a1f4062a0074a2794853d0d/12/GPL/openjdk-21.0.1_linux-x64_bin.tar.gz && \
    tar -xvzf openjdk-21.0.1_linux-x64_bin.tar.gz  && \
    rm -rf openjdk-21.0.1_linux-x64_bin.tar.gz && \
    mv jdk-21.0.1 /usr/local/ && \
    cd /usr/local/jdk-21.0.1

ENV JAVA_HOME=/usr/local/jdk-21.0.1
