# containers/Makefile.d/build.mk
# (c) 2023 Sam Caldwell.  See License.txt
#

# --------------------------------------------------------------------------------------------------------
# Build container images
#
build/containers: build/containers/opsys \
			  build/containers/language \
			  build/containers/services
	@docker images
	@color -green -lf -reset ">Completed $@"

# --------------------------------------------------------------------------------------------------------
# Build operating system images
#

build/containers/opsys: build/containers/opsys/ubuntu/20.04 \
						build/containers/opsys/ubuntu/22.04 \
						build/containers/opsys/ubuntu/23.04 \
						build/containers/opsys/ubuntu/latest
	@color -green -lf -reset ">Completed $@"

build/containers/opsys/ubuntu/20.04:
	@color -blue -lf -reset ">Start $@"
	@docker build --tag opsys/ubuntu:20.04 \
                  --file containers/opsys/ubuntu/20.04/Dockerfile . || exit 1
	@color -green -lf -reset ">>>Completed $@"

build/containers/opsys/ubuntu/22.04:
	@color -blue -lf -reset ">Start $@"
	@docker build --tag opsys/ubuntu:22.04 \
                  --file containers/opsys/ubuntu/22.04/Dockerfile . || exit 1
	@color -green -lf -reset ">>>Completed $@"

build/containers/opsys/ubuntu/23.04:
	@color -blue -lf -reset ">Start $@"
	@docker build --tag opsys/ubuntu:23.04 \
                  --file containers/opsys/ubuntu/23.04/Dockerfile . || exit 1
	@color -green -lf -reset ">>>Completed $@"

build/containers/opsys/ubuntu/latest:
	@color -blue -lf -reset ">Start $@"
	@docker tag opsys/ubuntu:23.04 opsys/ubuntu:latest || exit 1
	@color -green -lf -reset ">>>Completed $@"

# --------------------------------------------------------------------------------------------------------
# Build language-specific images
#

build/containers/language: build/containers/language/python/3 \
  						   build/containers/language/node/20.10.0 \
  						   build/containers/language/java/21
	@color -green -lf -reset ">Completed $@"

build/containers/language/python/3:
	@color -blue -lf -reset ">Start $@"
	@docker build --tag language/python:3 --file containers/language/python/3/Dockerfile .
	@color -green -lf -reset ">>>Completed $@"

build/containers/language/node/20.10.0:
	@color -blue -lf -reset ">Start $@"
	@docker build --tag language/node:20.10.0 --file containers/language/node/20.10.0/Dockerfile .
	@color -green -lf -reset ">>>Completed $@"

build/containers/language/java/21:
	@color -blue -lf -reset ">Start $@"
	@docker build --tag language/java:21 --file containers/language/java/21/Dockerfile .
	@color -green -lf -reset ">>>Completed $@"

# --------------------------------------------------------------------------------------------------------
# Build service-specific images
#

build/containers/services: build/containers/services/dns-server \
						   build/containers/services/postgresql \
						   build/containers/services/jenkins/21
	@color -green -lf -reset ">Completed $@"

build/containers/services/dns-server:
	@color -blue -lf -reset ">Start $@"
	@docker build --compress --tag services/dns-server:latest  --file containers/services/dns-server/Dockerfile .
	@color -green -lf -reset ">>>Completed $@"

build/containers/services/postgresql:
	@color -blue -lf -reset ">Start $@"
	@docker build --compress --tag services/postgresql:latest  --file containers/services/postgresql/Dockerfile .
	@color -green -lf -reset ">>>Completed $@"

build/containers/services/jenkins/21:
	@color -yellow -lf -reset ">>>Disabled $@"${RESET}
#	@color -blue -lf -reset ">Start $@"
#	@docker build --tag services/jenkins:21 --file containers/services/jenkins/21/Dockerfile .
#	@color -green -lf -reset ">>>Completed $@"
