/**
 * @name networking/TcpClient/src/sendMsg.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
namespace networking {
    /**
     * @name sendMsg
     * @brief
     *
     * @param msg
     * @param size
     * @return
     */
    PipeResponse TcpClient::sendMsg(const char *msg, size_t size) {
        const size_t numBytesSent = send(SocketFileDescriptor.get(), msg, size, 0);

        if (numBytesSent < 0) { // send failed
            return PipeResponse::failure(strerror(errno));
        }
        if (numBytesSent < size)
            return PipeResponse::failure(msgNotAllBytesSent);
        return PipeResponse::success();
    }
}