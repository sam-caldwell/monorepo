#
# Makefile.d/_config/100.has_dpkg.package-manager.mk
# (c) Sam Caldwell.  See LICENSE.txt
#
# linux: redhat/centos/fedora/sles rpm package manager
#
HAS_RPM=$(shell if command -v rpm &>/dev/null; then \
                  echo 'yes';\
                else \
                  echo 'no'; \
                fi;\
                )
has_rpm:
	@echo "$(HAS_RPM)"
