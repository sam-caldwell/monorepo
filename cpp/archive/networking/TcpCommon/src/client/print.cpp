/**
 * @name networking/TcpCommon/src/client/print.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../FileDescriptor.h"

namespace networking {
    /**
     * @name print
     * @brief print a status of our current connection.
     */
    void Client::print() const {
        std::cout << "------------------------STATE------------------------" << endl
                  << "  IP address: " << getIp() << std::endl
                  << "  Socket FD:  " << SocketFileDescriptor.get() << std::endl
                  << "  Connected?: " << string(isConnected() ? "True" : "False") << std::endl
                  << "-----------------------------------------------------" << endl;
    }
}