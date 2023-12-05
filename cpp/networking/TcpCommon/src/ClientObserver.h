/**
 * @name networking/TcpCommon/src/ClientObserver.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef networking_TcpCommon_ClientObserver_H
#define networking_TcpCommon_ClientObserver_H

#include <string>
#include <functional>
#include "projects/networking/TcpCommon/src/PipeResponse/main.h"

namespace networking {
    /**
     * @struct ClientObserver
     * @brief
     */
    struct ClientObserver {
        std::string wantedIP = "";
        std::function<void(const char *msg, size_t size)> IncomingPacketHandler = nullptr;
        std::function<void(const PipeResponse &ret)> DisconnectionHandler = nullptr;
    };
}

#endif /* networking_TcpCommon_ClientObserver_H */
