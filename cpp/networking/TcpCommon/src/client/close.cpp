/**
 * @name networking/TcpCommon/src/client/close.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

namespace networking {
    /**
     * @name close
     * @brief close the connection
     */
    void Client::close() {
        terminateReceiveThread();

        const bool closeFailed = (::close(SocketFileDescriptor.get()) == -1);
        if (closeFailed) {
            throw std::runtime_error(
                    string("TcoCommon/client failed:") +
                    string(strerror(errno)));
        }
    }
}