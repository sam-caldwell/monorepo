clean/docker:
	for i in $$(docker images | awk '{print $$3}' | grep -v IMAGE); do docker rmi -f $$i; done

build/docker: build/docker/opsys \
			  build/docker/language

build/docker/opsys:
	@docker build --tag opsys/ubuntu:20.04 --file opsys/ubuntu/20.04/Dockerfile .
	@docker build --tag opsys/ubuntu:22.04 --file opsys/ubuntu/22.04/Dockerfile .
	@docker build --tag opsys/ubuntu:23.04 --file opsys/ubuntu/23.04/Dockerfile .
	@docker tag opsys/ubuntu:23.04 opsys/ubuntu:latest
	@docker images


build/docker/language:
	@docker build --tag language/python:3 --file language/python/3/Dockerfile .
	@docker build --tag language/node:20.10.0 --file language/node/20.10.0/Dockerfile .
