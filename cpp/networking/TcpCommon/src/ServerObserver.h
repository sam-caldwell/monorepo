/**
 * @name networking/TcpCommon/src/ServerObserver.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef networking_TcpCommon_ServerObserver_H
#define networking_TcpCommon_ServerObserver_H

#include <string>
#include <functional>
#include "client/main.h"

namespace networking {
    /**
     * @struct ServerObserver
     * @brief
     */
    struct ServerObserver {
        std::string wantedIP = "";
        std::function<void(const std::string &clientIP, const char *msg, size_t size)> IncomingPacketHandler;
        std::function<void(const std::string &ip, const std::string &msg)> DisconnectionHandler;
    };
}

#endif /* networking_TcpCommon_ServerObserver_H */