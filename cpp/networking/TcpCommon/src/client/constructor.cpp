/**
 * @name networking/TcpCommon/src/client/constructor.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

namespace networking {
    /**
     * @name class constructor
     * @brief set up state.
     *
     * @param fileDescriptor
     */
    Client::Client(int fileDescriptor) {
        SocketFileDescriptor.set(fileDescriptor);
        setConnected(false);
    }
}