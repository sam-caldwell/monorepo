/**
 * @name networking/TcpClient/src/connectTo.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
namespace networking {
    /**
     * @name connectTo
     * @brief
     *
     * @param address string
     * @param port int
     * @return PipeResponse
     */
    PipeResponse TcpClient::connectTo(const std::string &address, int port) {
        try {
            initializeSocket();
            setAddress(address, port);
        } catch (const std::runtime_error &error) {
            return PipeResponse::failure(error.what());
        }

        const int connectResult = connect(SocketFileDescriptor.get(), (struct sockaddr *) &_server, sizeof(_server));
        const bool connectionFailed = (connectResult == -1);
        if (connectionFailed) {
            return PipeResponse::failure(strerror(errno));
        }

        startReceivingMessages();
        _isConnected = true;
        _isClosed = false;

        return PipeResponse::success();
    }
}