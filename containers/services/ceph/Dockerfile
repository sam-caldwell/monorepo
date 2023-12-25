# containers/service/ceph/Dockerfile
# (c) 2023 Sam Caldwell.  See License.txt.
#
# This is a multistage builder for ceph which builds
# the ceph software from source and produces containers
# built from the monorepo's approved base images.
#
ARG CEPH_VERSION=18.2.1
ARG BASE_UBUNTU_VERSION=22.04

FROM opsys/ubuntu:${BASE_UBUNTU_VERSION} as base
RUN apt-get update -y --fix-missing && \
    apt-get upgrade -y

FROM base as ceph-builder
ARG CEPH_VERSION
RUN apt-get install -y build-essential git cmake nodejs python3-routes curl wget
WORKDIR /opt/
RUN wget https://download.ceph.com/tarballs/ceph_${CEPH_VERSION}.orig.tar.gz && \
    tar -xvsf ceph_${CEPH_VERSION}.orig.tar.gz && \
    mv ceph_${CEPH_VERSION} ceph
WORKDIR /opt/ceph

#RUN rm -rf .git
#RUN ./install-deps.sh
#RUN rm -rf build || true
#RUN ./do_cmake.sh -Wno-dev \
#                  -DCMAKE_BUILD_TYPE=RelWithDebInfo \
#                  -DENABLE_GIT_VERSION=OFF \
#                  -DCEPH_GIT_VER=v18.2.1 \
#                  -DCEPH_GIT_NICE_VER=production && \
#    cd build && \
#    ninja




# Multi-stage Dockerfile for Ceph
#
## Stage 1: Build Ceph Base Image
#FROM ceph/daemon AS ceph-base
#
## Stage 2: Monitor Container
#FROM ceph-base AS monitor
#LABEL stage=monitor
#
## Set environment variables
#ENV MON_NAME=ceph-mon
#ENV CEPH_NETWORK=172.18.0.0/16
#
## Run Ceph monitor
#ENTRYPOINT ["mon"]
#CMD ["mon"]
#
## Stage 3: OSD Container
#FROM ceph-base AS osd
#LABEL stage=osd
#
## Set environment variables
#ENV OSD_TYPE=directory
#ENV OSD_NAME=ceph-osd
#
## Run Ceph OSD
#ENTRYPOINT ["osd"]
#CMD ["osd"]
#
## Stage 4: Manager Container
#FROM ceph-base AS mgr
#LABEL stage=mgr
#
## Run Ceph manager
#ENTRYPOINT ["mgr"]
#CMD ["mgr"]
#
## Stage 5: Client Container
#FROM ceph-base AS client
#LABEL stage=client
#
## Run Ceph client
#ENTRYPOINT ["demo"]
#CMD ["demo"]

# Example Usage:


# Note: This is a simplified example, and you may need to customize it based on your specific requirements and Ceph configuration.
