Command: What
=============

## Description

The tool formerly known as `wtf` in the Asymmetric-Toolkit
is now named `what` for this repo because...well...some of
y'all have clearly never been oncall asking "WTF?" about an
unfamiliar system.

## Usage

```bash
what <option> [--printError]

options:
  --help          : display usage information
  --printError    : indicates whether error messages are printed to stdout.

  --cpus          : Return number of CPU cores
  --cpu-arch      : Return cpu architecture (e.g. arm64, amd64)
  --cpu-info      : Return CPU specifications
  --os            : Return the operating system (e.g. darwin, linux, windows)
  --os-family     : Return an operating system family (e.g. Windows 10)
  --os-version    : Return the operating system version
  --ram           : Return the amount of memory (in KB)qq
  --software-list : Return a list of software installed on the system

```

> NOTE: the above may not be updated. Always use `what --help`
> for the latest features for your specific version

> TODO:
> Add versioning (--version)

