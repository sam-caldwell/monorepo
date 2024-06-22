/**
 * @name networking/TcpCommon/src/client/send.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../messages.h"

namespace networking {
    /**
     * @name send (c-string)
     * @brief send a message of a given size to the destination server.
     *
     * @param msg - const char (null-terminated c string)
     * @param msgSize message size (length)
     */
    void Client::send(const char *msg, size_t msgSize) const {

        const size_t numBytesSent = ::send(SocketFileDescriptor.get(), (char *) msg, msgSize, 0);

        const bool sendFailed = (numBytesSent < 0);
        if (sendFailed) {
            throw std::runtime_error(strerror(errno));
        }

        const bool notAllBytesWereSent = (numBytesSent < msgSize);
        if (notAllBytesWereSent)
            throw std::runtime_error(msgNotAllBytesSent);
    }

    /**
     * @name send (string)
     * @brief send a message of a given size to the destination server.
     *
     * @param msg - std::string by reference
     */
    inline void Client::send(std::string *msg) const {
        send(msg->c_str(), msg->length());
    }
}