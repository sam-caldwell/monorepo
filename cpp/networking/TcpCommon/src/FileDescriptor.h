/**
 * @name networking/TcpCommon/src/clientEvent.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef networking_TcpCommon_FileDescriptor_H
#define networking_TcpCommon_FileDescriptor_H

namespace networking {
    /**
     * @class FileDescriptor
     * @brief provides a network socket file descriptor state.
     */
    class FileDescriptor {
    public:
        /**
         * @name set
         * @brief set the file descriptor
         *
         * @param fd - int
         */
        void set(int fd) { socketFileDescriptor = fd; }

        /**
         * @name get
         * @brief return the file descriptor
         *
         * @return int
         */
        int get() const { return socketFileDescriptor; }

    private:
        /**
         * @name socketFileDescriptor
         * @brief a file descriptor for the current network socket...'cause all things *nix are files...
         */
        int socketFileDescriptor = 0;

    };
}

#endif /* networking_TcpCommon_FileDescriptor_H */