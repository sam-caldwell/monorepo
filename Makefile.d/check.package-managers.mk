#
# PACKAGE MANAGERS
#
# mac/linux package managers
HAS_BREW=$(shell go run cmd/tools/hasExecutable/main.go brew)
# linux: debian/ubuntu
HAS_APT=$(shell go run cmd/tools/hasExecutable/main.go apt-get)
HAS_DPKG=$(shell go run cmd/tools/hasExecutable/main.go dpkg)

# linux: redhat/centos/fedora
HAS_YUM=$(shell go run cmd/tools/hasExecutable/main.go yum)

# linux: redhat/centos/fedora/sles
HAS_RPM=$(shell go run cmd/tools/hasExecutable/main.go rpm)

# windows:
HAS_WINGET=$(shell go run cmd/tools/hasExecutable/main.go winget)
HAS_CHOCO=$(shell go run cmd/tools/hasExecutable/main.go choco)

#
# POSIX OPERATING SYSTEMS
#
check/package-managers:
	[[ "$${HAS_BREW}_" =="yes_" ]] && {echo "brew" && exit 0;}; \
	[[ "$${HAS_APT}_" =="yes_" ]] && {echo "apt" && exit 0;}; \
	[[ "$${HAS_YUM}_" =="yes_" ]] && {echo "yum" && exit 0;}; \
	[[ "$${HAS_DPKG}_" =="yes_" ]] && {echo "dpkg" && exit 0;}; \
	[[ "$${HAS_RPM}_" =="yes_" ]] && {echo "rpm" && exit 0;}; \
	[[ "$${HAS_WINGET}_" =="yes_" ]] && {echo "winget" && exit 0;}; \
	[[ "$${HAS_CHOCO}_" =="yes_" ]] && {echo "choco" && exit 0;}; \


check/package-managers/list:
	@echo "\033[32m detected package managers: \033[0m"
	@echo "\033[34m   macos/linux: \033[0m"
	@echo "\033[34m\t     HAS_BREW: $(HAS_BREW) \033[0m"
	@echo "\033[34m   linux: debian/ubuntu: \033[0m"
	@echo "\033[34m\t      HAS_APT: $(HAS_APT) \033[0m"
	@echo "\033[34m\t     HAS_DPKG: $(HAS_DPKG) \033[0m"
	@echo "\033[34m   linux: redhat/centos/fedora: \033[0m"
	@echo "\033[34m\t      HAS_YUM: $(HAS_YUM) \033[0m"
	@echo "\033[34m   linux: redhat/centos/fedora/suse: \033[0m"
	@echo "\033[34m\t      HAS_RPM: $(HAS_RPM) \033[0m"
	@echo "\033[34m   windows: \033[0m"
	@echo "\033[34m\t   HAS_WINGET: $(HAS_WINGET) \033[0m"
	@echo "\033[34m\t    HAS_CHOCO: $(HAS_CHOCO) \033[0m"
	@echo "\033[32m --- \033[0m"
