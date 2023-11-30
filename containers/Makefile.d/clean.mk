# containers/Makefile.d/clean.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
clean/containers: stop/containers
	@echo ${GREEN}"$@ >> removing containers"
	@for i in $$(docker images | awk '{print $$3}' | grep -v IMAGE); do docker rmi -f $$i; done
	@echo ${GREEN}">>>Completed $@"${RESET}

stop/containers:
	@echo ${GREEN}"$@ >> stopping containers"
	@for i in $$(docker ps | awk '{print $$1}' | grep -v CONTAINER); do docker kill $$i; done
	@echo ${GREEN}">>>Completed $@"${RESET}
