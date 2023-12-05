/**
* @name networking/TcpClient/src/connectTo.cpp
* @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
* @author Sam Caldwell <mail@samcaldwell.net>
*/

#include "projects/networking/TcpCommon/src/constants.h"

namespace networking {
    /**
     * @name receiveTask
     * @brief Receive server packets, and notify user
     */
    void TcpClient::receiveTask() {
        while (_isConnected) {
            switch (FdWait::waitFor(SocketFileDescriptor, SOCKET_WAIT_TIMEOUT)) {
                case FdWait::Result::FAILURE:
                    throw std::runtime_error(strerror(errno));
                    break;
                case FdWait::Result::TIMEOUT:
                    continue;
                default:
                    break;
            }

            char msg[MAX_PACKET_SIZE];
            const size_t numOfBytesReceived = recv(SocketFileDescriptor.get(), msg, MAX_PACKET_SIZE, 0);

            if (numOfBytesReceived < 1) {
                _isConnected = false;
                publishServerDisconnected(PipeResponse::failure(
                        (numOfBytesReceived == 0)
                        ? serverClosedConnection
                        : strerror(errno)));
            } else {
                publishServerMsg(msg, numOfBytesReceived);
            }
        }
    }
}