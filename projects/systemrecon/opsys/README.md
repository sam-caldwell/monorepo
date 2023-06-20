SystemRecon/Opsys
=================

## Description

Cross-platform system recon tools for enumerating the
operating system characteristics and configuration of
a target host

## Features

### `Opsys()`

#### Description

Returns the current host's operating system using `GOOS`

#### Supported On

* Windows
* MacOs
* Linux
* (any system supported by GOOS)

### `OpsysFamily()`

#### Description

Parses the system version of a given host and maps it
to a branded family of the operating system. For example,
Windows 10 vs Windows XP, MacOS Ventura versus Lion or
a Linux distribution (Debian, Ubuntu, Redhat, etc.)

#### Supported On

* Windows
> Note that while many Windows versions are supported
> an unsupported version will return major-minor version
> strings.  Adding more is trivial.

* MacOs
> Note that while many versions are supported
> an unsupported version will return major-minor version
> strings.  Adding more is trivial.

* Linux
> Only Ubuntu, Debian, Redhat, Centos, Fedora and a few
> other linux distributions are discoverable here.
> 
> Sorry Arch-bros, but when your hobbyist system becomes mainstream
> on systems I get paid to work on, I'll add them here.  Otherwise
> send a Pull Request, and I'll make sure your egos don't get
> bruised.  Not everyone uses Arch.  It's painful...but...well...
> this is reality in 2023.

#### To Add Support for New Operating Systems:
* Add a `GetFamily()` function for your operating system
  and make sure your build constraints and function signature
  match.

* In [OpsysFamily.go](OpSysFamily.go), add a call to
  your system-specific `GetFamily()` function.

#### To Add More Version Mappings For Existing Systms
* Edit the system-specific `GetFamily()` function.
* Please be sure to add the new system to the `words`
  constants files to keep consistent spelling and
  formatting across all our tooling.