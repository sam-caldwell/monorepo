/**
 * @name networking/TcpCommon/src/client/publishEvent.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <string.h>
#include "projects/networking/TcpCommon/src/clientEvent.h"

namespace networking {
    /**
     * @name publishEvent
     * @brief publish an event to the event handler
     *
     * @param clientEvent - enumerated type.
     * @param msg
     */
    void Client::publishEvent(ClientEvent clientEvent, const std::string &msg) {
        _eventHandlerCallback(*this, clientEvent, msg);
    }
}