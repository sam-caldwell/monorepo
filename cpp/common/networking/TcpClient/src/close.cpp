/**
 * @name networking/TcpClient/src/close.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
namespace networking {
    /**
     * @name close
     * @brief close the network connection
     *
     * @return PipeResponse
     */
    PipeResponse TcpClient::close() {
        if (_isClosed) {
            return PipeResponse::failure("client is already closed");
        }
        terminateReceiveThread();

        const bool closeFailed = (::close(SocketFileDescriptor.get()) == -1);
        if (closeFailed) {
            return PipeResponse::failure(strerror(errno));
        }
        _isClosed = true;
        return PipeResponse::success();
    }

}