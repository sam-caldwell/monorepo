/**
 * @name networking/TcpCommon/src/FdWait.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef networking_TcpCommon_WaitFor_H
#define networking_TcpCommon_WaitFor_H

#include <cstdint>
#include <cstdio>
#include "FileDescriptor.h"
#include <sys/select.h>

#define SELECT_FAILED -1
#define SELECT_TIMEOUT 0

namespace networking {

    namespace FdWait {
        /**
         * @name waitFor
         * @brief wait for I/O operation on a given file descriptor.
         * @return Result
         */
        Result waitFor(const FileDescriptor &fileDescriptor, uint32_t timeoutSeconds) {
            //
            //See https://pubs.opengroup.org/onlinepubs/007904875/basedefs/sys/select.h.html
            //
            //From <sys/select.h>...
            //   Create/configure timeval struct...
            struct timeval tv;
            tv.tv_sec = (timeoutSeconds == 0) ? MAX_CLIENT_TIMEOUT : timeoutSeconds;
            tv.tv_usec = 0;

            fd_set fds;

            //From <sys/select.h>...
            //    Set the bit for the file descriptor fd in the file descriptor set fdset.
            FD_ZERO(&fds);

            //From <sys/select.h>...
            //    Returns a non-zero value if the bit for the file descriptor fd is set in the
            //    file descriptor set by fdset, and 0 otherwise.
            FD_SET(fileDescriptor.get(), &fds);

            //From <sys/select.h>...
            //    int select(int, fd_set *restrict, fd_set *restrict, fd_set *restrict, struct timeval *restrict);
            const int selectRet = select(fileDescriptor.get() + 1, &fds, nullptr, nullptr, &tv);

            switch (selectRet) {
                case SELECT_FAILED:
                    return Result::FAILURE;
                    break;
                case SELECT_TIMEOUT:
                    return Result::TIMEOUT;
                    break;
                default:
                    return Result::SUCCESS;
            }
        }

    }
}

#endif /* networking_TcpCommon_WaitFor_H */