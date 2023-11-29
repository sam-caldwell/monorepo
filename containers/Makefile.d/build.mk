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
	@echo ${GREEN}">Completed $@"${RESET}

# --------------------------------------------------------------------------------------------------------
# Build operating system images
#

build/containers/opsys: build/containers/opsys/ubuntu/20.04 \
						build/containers/opsys/ubuntu/22.04 \
						build/containers/opsys/ubuntu/23.04 \
						build/containers/opsys/ubuntu/latest
	@echo ${GREEN}">>Completed $@"${RESET}

build/containers/opsys/ubuntu/20.04:
	@echo ${GREEN}">>>Start $@"${RESET}
	@docker build --tag opsys/ubuntu:20.04 --file containers/opsys/ubuntu/20.04/Dockerfile . || exit 1
	@echo ${GREEN}">>>Completed $@"${RESET}

build/containers/opsys/ubuntu/22.04:
	@echo ${GREEN}">>>Start $@"${RESET}
	@docker build --tag opsys/ubuntu:22.04 --file containers/opsys/ubuntu/22.04/Dockerfile . || exit 1
	@echo ${GREEN}">>>Completed $@"${RESET}

build/containers/opsys/ubuntu/23.04:
	@echo ${GREEN}">>>Start $@"${RESET}
	@docker build --tag opsys/ubuntu:23.04 --file containers/opsys/ubuntu/23.04/Dockerfile . || exit 1
	@echo ${GREEN}">>>Completed $@"${RESET}

build/containers/opsys/ubuntu/latest:
	@echo ${GREEN}">>>Start $@"${RESET}
	@docker tag opsys/ubuntu:23.04 containers/opsys/ubuntu:latest || exit 1
	@echo ${GREEN}">>>Completed $@"${RESET}

# --------------------------------------------------------------------------------------------------------
# Build language-specific images
#

build/containers/language: build/containers/language/python/3 \
  						   build/containers/language/node/20.10.0 \
  						   build/containers/language/java/21
	@echo ${GREEN}">>Completed $@"${RESET}

build/containers/language/python/3:
	@echo ${GREEN}">>>Start $@"${RESET}
	@docker build --tag language/python:3 --file containers/language/python/3/Dockerfile .
	@echo ${GREEN}">>>Completed $@"${RESET}

build/containers/language/node/20.10.0:
	@echo ${GREEN}">>>Start $@"${RESET}
	@docker build --tag language/node:20.10.0 --file containers/language/node/20.10.0/Dockerfile .
	@echo ${GREEN}">>>Completed $@"${RESET}

build/containers/language/java/21:
	@echo ${GREEN}">>>Start $@"${RESET}
	@docker build --tag language/java:21 --file containers/language/java/21/Dockerfile .
	@echo ${GREEN}">>>Completed $@"${RESET}

# --------------------------------------------------------------------------------------------------------
# Build service-specific images
#

build/containers/services: build/containers/services/dns-server \
						   build/containers/services/postgresql \
						   build/containers/services/jenkins/21
	@echo ${GREEN}">>Completed $@"${RESET}

build/containers/services/dns-server:
	@echo ${GREEN}">>>Start $@"${RESET}
	@docker build --compress --tag services/dns-server:latest  --file containers/services/dns-server/Dockerfile .
	@echo ${GREEN}">>>Completed $@"${RESET}

build/containers/services/postgresql:
	@echo ${GREEN}">>>Start $@"${RESET}
	@docker build --compress --tag services/postgresql:latest  --file containers/services/postgresql/Dockerfile .
	@echo ${GREEN}">>>Completed $@"${RESET}

build/containers/services/jenkins/21:
	@echo ${YELLOW}">>>Disabled $@"${RESET}
#	@echo ${GREEN}">>>Start $@"${RESET}
#	@docker build --tag services/jenkins:21 --file containers/services/jenkins/21/Dockerfile .
#	@echo ${GREEN}">>>Completed $@"${RESET}
