/**
 * @name flags/src/opsys.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef FLAGS_OPSYS_H
#define FLAGS_OPSYS_H

/**
 * operating system flags...
 */
#define OPSYS_WINDOWS 0
#define OPSYS_LINUX 1
#define OPSYS_MACOS 2
/**
 * Determine OPSYS
 */
#if (defined(_WIN32) || defined(_WIN64) || defined(__WIN32__) || defined(__TOS_WIN__) || defined(__WINDOWS__))

/*  */#define OPSYS OPSYS_WINDOWS

#elif (defined(linux) || defined(__linux) || defined(__linux__) || defined(__gnu_linux__))

/*  */#define OPSYS OPSYS_LINUX

#elif (defined(macintosh) || defined(Macintosh) || (defined(__APPLE__) && defined(__MACH__)))

/*  */#define OPSYS OPSYS_MACOS

#else

/*  */#ifdef OPSYS
/*  */#undef OPSYS
/*  */#endif
/*  */#error Unsupported OS

#endif/* OPSYS */

#endif /* FLAGS_OPSYS_H */