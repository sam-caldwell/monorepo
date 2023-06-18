check/package-managers/list:
	$(PRINT_START)
	@echo "$(ANSI_BLUE) detected package managers: $(ANSI_RESET)"
	@echo "$(ANSI_BLUE)   macos/linux: $(ANSI_RESET)"
	@echo "$(ANSI_BLUE)\t     HAS_BREW: $(HAS_BREW) $(ANSI_RESET)"
	@echo "$(ANSI_BLUE)   linux: debian/ubuntu: $(ANSI_RESET)"
	@echo "$(ANSI_BLUE)\t      HAS_APT: $(HAS_APT) $(ANSI_RESET)"
	@echo "$(ANSI_BLUE)\t     HAS_DPKG: $(HAS_DPKG) $(ANSI_RESET)"
	@echo "$(ANSI_BLUE)   linux: redhat/centos/fedora: $(ANSI_RESET)"
	@echo "$(ANSI_BLUE)\t      HAS_YUM: $(HAS_YUM) $(ANSI_RESET)"
	@echo "$(ANSI_BLUE)   linux: redhat/centos/fedora/suse: $(ANSI_RESET)"
	@echo "$(ANSI_BLUE)\t      HAS_RPM: $(HAS_RPM) $(ANSI_RESET)"
	@echo "$(ANSI_BLUE)   windows: $(ANSI_RESET)"
	@echo "$(ANSI_BLUE)\t   HAS_WINGET: $(HAS_WINGET) $(ANSI_RESET)"
	@echo "$(ANSI_BLUE)\t    HAS_CHOCO: $(HAS_CHOCO) $(ANSI_RESET)"
	$(PRINT_DONE)
