# containers/Makefile.d/clean.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
clean/containers: stop/containers
	@color -green -lf -reset "$@ >> removing containers"
	@for i in $$(docker images | awk '{print $$3}' | grep -v IMAGE); do docker rmi -f $$i; done
	@color -green -lf -reset ">>>Completed $@"

stop/containers:
	@color -green -lf -reset "$@ >> stopping containers"
	@for i in $$(docker ps | awk '{print $$1}' | grep -v CONTAINER); do docker kill $$i; done
	@color -green -lf -reset ">>>Completed $@"
