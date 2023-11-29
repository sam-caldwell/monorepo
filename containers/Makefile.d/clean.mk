# containers/Makefile.d/clean.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
clean/containers:
	for i in $$(docker images | awk '{print $$3}' | grep -v IMAGE); do docker rmi -f $$i; done
