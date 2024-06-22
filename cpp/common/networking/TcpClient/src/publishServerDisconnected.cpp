/**
 * @name networking/TcpClient/src/publishServerDisconnected.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

namespace networking {
    /**
     * @name publishServerDisconnected
     * @brief Publish a client disconnection to observer, and the observer will only get a notification
     *        where the client Ip address is identical to the specific observer requested Ip.
     *
     * @param result PipeResponse (by reference)
     */
        void TcpClient::publishServerDisconnected(const PipeResponse &result) {
        std::lock_guard <std::mutex> lock(_subscribersMtx);
        for (const auto &subscriber: _subscibers) {
            if (subscriber.DisconnectionHandler) {
                subscriber.DisconnectionHandler(result);
            }
        }
    }
}