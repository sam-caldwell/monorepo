check/package-managers:
	#
	# PACKAGE MANAGERS
	#
	# mac/linux package managers
	HAS_BREW=$(shell command -v brew &> /dev/null && echo yes)
	# linux: debian/ubuntu
	HAS_APT=$(shell command -v apt-get &> /dev/null && echo yes)
	HAS_DPKG=$(shell command -v dpkg &> /dev/null && echo yes)

	# linux: redhat/centos/fedora
	HAS_YUM=$(shell command -v yum &> /dev/null && echo yes)

	# linux: redhat/centos/fedora/sles
	HAS_RPM=$(shell command -v rpm &> /dev/null && echo yes)
	#
	# POSIX OPERATING SYSTEMS
	#