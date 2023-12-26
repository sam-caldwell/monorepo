# containers/service/ceph/Dockerfile
# (c) 2023 Sam Caldwell.  See License.txt.
#
# This is a multistage builder for ceph which builds
# the ceph software from source and produces containers
# built from the monorepo's approved base images.
#
ARG CEPH_VERSION=18.2.1
ARG BASE_UBUNTU_VERSION=22.04

FROM opsys/ubuntu:${BASE_UBUNTU_VERSION} AS base
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update -y --fix-missing && \
    apt-get upgrade -y

FROM base AS ceph-builder
ARG CEPH_VERSION
RUN apt-get install -y build-essential git cmake nodejs python3-routes curl wget openssl libssl-dev \
                       jq libzstd-dev libre2-dev libre2-9 zlib1g-dev cryptsetup libthrift-dev ccache
WORKDIR /opt/
COPY containers/services/ceph/ceph_${CEPH_VERSION}.orig.tar.gz /opt/
RUN tar -xvzf ceph_${CEPH_VERSION}.orig.tar.gz
RUN rm ceph_${CEPH_VERSION}.orig.tar.gz
WORKDIR /opt/ceph-${CEPH_VERSION}
RUN ./install-deps.sh
# we should probably set CMAKE_CROSSCOMPILING=false
RUN ./do_cmake.sh -Wno-dev \
                  -DCMAKE_BUILD_TYPE=RelWithDebInfo \
                  -DENABLE_GIT_VERSION=OFF \
                  -DCEPH_GIT_VER=${CEPH_VERSION} \
                  -DCEPH_GIT_NICE_VER=production \
                  -DCMAKE_CXX_STANDARD=11 \
                  -DCMAKE_INSTALL_LIBDIR=/opt/ceph-${CEPH_VERSION}/build/ \
                  -DHAVE_UNALIGNED_ACCESS=true \
                  -DWITH_CCACHE=ON \
                  -DWITH_MANPAGE=ON \
                  -DWITH_BABELTRACE=ON \
                  -DWITH_MGR_DASHBOARD_FRONTEND=ON \
                  -DWITH_RBD=ON \
                  -DWITH_KRBD=ON \
                  -DWITH_RADOSGW=ON
RUN cd build && ninja
WORKDIR /opt/
RUN rm -rf /opt/ceph-${CEPH_VERSION}

FROM base AS ceph-artifacts
ARG CEPH_VERSION
WORKDIR /opt/
COPY --from=ceph-builder /opt/ceph-${CEPH_VERSION}/build/bin/* /usr/local/bin/




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
