/**
 * @name networking/TcpCommon/src/client/setEventsHandler.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

namespace networking {
    /**
     * @name setEventsHandler
     * @brief set the event handler callback function.
     *
     * @param eventHandler event handler function pointer.
     */
    inline void Client::setEventsHandler(const ClientEventHandler &eventHandler) {
        _eventHandlerCallback = eventHandler;
    }
}