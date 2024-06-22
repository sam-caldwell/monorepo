/**
 * @name networking/TcpCommon/src/clientEvent.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef networking_TcpCommon_ClientEvent_H
#define networking_TcpCommon_ClientEvent_H

namespace networking {
    /**
     * @enum ClientEvent
     * @brief Describes the nature of a client event.
     */
    enum ClientEvent {
        DISCONNECTED,
        INCOMING_MSG
    };
}

#endif /* networking_TcpCommon_ClientEvent_H */