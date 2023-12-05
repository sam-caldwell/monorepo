/**
 * @name flags/src/arch.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef FLAGS_ARCH_H
#define FLAGS_ARCH_H

/**
 * CPU architecture
 */
#define ARCH_X32 0
#define ARCH_X64 1
#define ARCH_ARM32 2
#define ARCH_ARM64 3

#if (defined(__amd64__) || defined(__x86_64__))

/*  */#define CPU_ARCH ARCH_X64

#elif defined(__aarch64__)

/*  */#define CPU_ARCH ARCH_ARM64

#elif defined(__arm__)

/*  */#define CPU_ARCH ARCH_ARM32

#elif defined(__i386__)

/*  */#define CPU_ARCH ARCH_X32

#else /*CPU_ARCH */

/*  */#ifdef CPU_ARCH
/*  */#undef CPU_ARCH
/*  */#endif
/*  */#error Unsupported Architecture

#endif/*CPU_ARCH */

#endif /* FLAGS_ARCH_H */