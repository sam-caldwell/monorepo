#
# Makefile.d/packer/build.packer.windows.mk
# (c) Sam Caldwell.  See LICENSE.txt
#
build/packer/windows/: build/packer/windows/10 \
					   build/packer/windows/11
	@echo "\033[32m>ok $@\033[0m"