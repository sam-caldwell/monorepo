/**
 * @name networking/TcpCommon/src/client/receiveTask.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/networking/TcpCommon/src/FileDescriptor.h"

namespace networking {
    /**
     * @name receiveTask
     * @brief receive a message.
     */
    void Client::receiveTask() {

        while (isConnected()) {

            const FdWait::Result waitResult = networking::FdWait::waitFor(SocketFileDescriptor, 0);

            if (waitResult == networking::FdWait::Result::FAILURE) {

                throw std::runtime_error(strerror(errno));

            } else if (waitResult == networking::FdWait::Result::TIMEOUT) {

                continue;

            }

            char receivedMessage[MAX_PACKET_SIZE];
            const size_t numOfBytesReceived = recv(SocketFileDescriptor.get(), receivedMessage, MAX_PACKET_SIZE, 0);

            if (numOfBytesReceived < 1) {
                const bool clientClosedConnection = (numOfBytesReceived == 0);
                std::string disconnectionMessage;

                if (clientClosedConnection) {

                    disconnectionMessage = "Client closed connection";

                } else {

                    disconnectionMessage = strerror(errno);

                }

                setConnected(false);
                publishEvent(ClientEvent::DISCONNECTED, disconnectionMessage);
                return;

            } else {

                publishEvent(ClientEvent::INCOMING_MSG, receivedMessage);

            }

        }

    }

}