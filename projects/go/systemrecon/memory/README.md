systemrecon/memory
==================

## Description

Memory-related systems reconniassance functions.

## Functions

### `MemInfo()`
#### Description
Provides the same approximate functionality as `cat /proc/meminfo` does in Linux
returning a key-value map of memory-related information for the operating system

#### Supported On
* Windows
* MacOs
* Linux

### `RamSize()`
#### Description
Returns the integer-value amount of RAM in the system (in KB)

#### Supported On
* Windows
* MacOs
* Linux
