#
# Makefile.d/_config/100.has_yum.package-manager.mk
# (c) Sam Caldwell.  See LICENSE.txt
#
# linux: redhat/centos/fedora yum
#
HAS_YUM=$(shell if command -v yum &>/dev/null; then \
          echo 'yes';\
        else \
          echo 'no'; \
        fi;\
        )
has_yum:
	@echo "$(HAS_YUM)"
