#
# PACKAGE MANAGERS
#
# linux: debian/ubuntu
HAS_APT=$(shell if command -v apt-get; then \
  echo 'yes';\
else \
  echo 'no'; \
fi;\
)

has_apt:
	@echo "$(HAS_APT)"