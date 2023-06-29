#
# Identify detected packet managers.
#

PACKAGE_MANAGER:=$(shell /bin/bash -c '\
		[ "$(HAS_BREW)_" == "yes_" ] && { \
			echo "brew"; \
			exit 0;\
		}; \
		[ "$(HAS_APT)_" == "yes_" ] && { \
			echo "apt"; \
		  	exit 0;\
		}; \
		[ "$(HAS_YUM)_" == "yes_" ] && { \
			echo "yum"; \
		  	exit 0;\
		}; \
		[ "$(HAS_DPKG)_" == "yes_" ] && { \
	  		echo "dpkg"; \
		  	exit 0;\
		}; \
		[ "$(HAS_RPM)_" == "yes_" ] && { \
	  		echo "rpm"; \
		  	exit 0;\
		}; \
		[ "$(HAS_WINGET)_" == "yes_" ] && { \
	  		echo "winget"; \
		  	exit 0;\
		}; \
		[ "$(HAS_CHOCO)_" == "yes_" ] && { \
	  		echo "choco"; \
		  	exit 0;\
		};\
		exit 0')

check/package-managers:
	@echo $(PACKAGE_MANAGER)

