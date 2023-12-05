/**
 * @name networking/TcpClient/src/publishServerMsg.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

namespace networking {
    /**
    * @name publishServerMsg
    * @brief publish an IncomingPacketHandler client message to the observer.
    * The observer will get only messages that originated from clients with an IP address
    * identical to the observer requested IP address.
    *
    * @param msg cstring
    * @param msgSize cstring length
    */
    void TcpClient::publishServerMsg(const char *msg, size_t msgSize) {
        std::lock_guard <std::mutex> lock(_subscribersMtx);
        for (const auto &subscriber: _subscibers) {
            if (subscriber.IncomingPacketHandler) {
                subscriber.IncomingPacketHandler(msg, msgSize);
            }
        }
    }
}