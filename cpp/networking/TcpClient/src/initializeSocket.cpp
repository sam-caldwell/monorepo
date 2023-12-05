/**
 * @name networking/TcpClient/src/initializeSocket.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
namespace networking {
    /**
     * @name initializeSocket
     * @brief initialize the network socket
     */
    void TcpClient::initializeSocket() {
        PipeResponse ret;

        SocketFileDescriptor.set(socket(AF_INET, SOCK_STREAM, 0));
        const bool socketFailed = (SocketFileDescriptor.get() == -1);
        if (socketFailed) {
            throw std::runtime_error(strerror(errno));
        }
    }
}