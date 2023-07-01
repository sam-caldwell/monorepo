Cpu Caches (L1, L2, L3)
=======================

## Description

This document details how we collect CPU Cache sizes from our supported systems.

## General Information

There are three (3) caches in a CPU:

| Cache    | Description                                                                                                                                |
|----------|--------------------------------------------------------------------------------------------------------------------------------------------|
| L1 cache | fastest, but smallest, data and instructions This is actually two caches (data, instructions) represented as `l1d` and `l1i` respectively. |
| L2 cache | slower, but bigger, data-only                                                                                                              |
| L3 cache | slowest, but biggest, data-only                                                                                                            |

Each operating system has a slightly different way of obtaining this information.

## Getting the data

### In Linux

The size (in KB) of each cache can be found in the `/sys/` file system using the following pattern:

```
/sys/devices/system/cpu/cpu0/cache/index{0,1,2,3}/size
```

Here we always assume `cpu0` (the CPUs are typically identical on a system for architectural simplicity).
Our `index{0,1,2,3}` refers to `L1d`, `L1i`, `L2` and `L3` cache, respectively. For example...

```text
# cat /sys/devices/system/cpu/cpu0/cache/index*/size
32K
32K
256K
16384K
```

Thus, we can read these files and obtain our cache sizes. But there is other information about the caches
that we can learn from this area:

```text
# ls -la /sys/devices/system/cpu/cpu0/cache/index0/     
total 0
drwxr-xr-x 2 root root    0 Jun 30 22:51 .
drwxr-xr-x 6 root root    0 Jun 30 22:51 ..
-r--r--r-- 1 root root 4096 Jun 30 22:51 coherency_line_size
-r--r--r-- 1 root root 4096 Jun 30 22:51 id
-r--r--r-- 1 root root 4096 Jun 30 22:51 level
-r--r--r-- 1 root root 4096 Jun 30 22:51 number_of_sets
-r--r--r-- 1 root root 4096 Jun 30 22:51 physical_line_partition
-r--r--r-- 1 root root 4096 Jun 30 22:54 shared_cpu_list
-r--r--r-- 1 root root 4096 Jun 30 22:51 shared_cpu_map
-r--r--r-- 1 root root 4096 Jun 30 22:51 size
-r--r--r-- 1 root root 4096 Jun 30 22:51 type
-rw-r--r-- 1 root root 4096 Jun 30 22:54 uevent
-r--r--r-- 1 root root 4096 Jun 30 22:51 ways_of_associativity
```

For example, we can see which CPUs share a given CPU's cache:

```text
# cat /sys/devices/system/cpu/cpu0/cache/index{0,1,2,3}/shared_cpu_list
0
0
0
0-7
```

We can identify which is our L1i and L1d cache:

```text
# cat /sys/devices/system/cpu/cpu0/cache/index{0,1,2,3}/type           
Data
Instruction
Unified
Unified
# cat /sys/devices/system/cpu/cpu0/cache/index{0,1,2,3}/level
1
1
2
3
```

### In Darwin (MacOS)

The size (in KB) of each cache can be found using system -n <cache name>, but it can also be found using the `sysctl`
command. Consistent with its principles of being a closed culture of paranoia (ironically one started by a hippie
like Steve Jobs), Apple's operating system is not well documented. There aren't a lot of ways to interact with the
operating system directly. But rather than reverse-engineer `sysctl` we just make a command-shell call out to this
always-present tool to get our information from one of three system properties:

* hw.l1icachesize
* hw.l2cachesize
* hw.l3cachesize

### In Windows (Windows 10+)
# Todo: see WHAT-19