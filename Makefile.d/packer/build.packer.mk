#
# Makefile.d/packer/build.packer.mk
# (c) Sam Caldwell.  See LICENSE.txt
#

build/packer: build/packer/setup \
			  build/packer/windows \
			  build/packer/linux \
			  build/packer/darwin
	  @echo "\033[32m>ok $@\033[0m"
