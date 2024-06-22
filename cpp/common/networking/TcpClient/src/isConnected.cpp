/**
 * @name networking/TcpClient/src/isConnected.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
namespace networking {
    /**
     * @name isConnected
     * @brief return the connected status of the socket.
     * @return bool
     */
    inline bool TcpClient::isConnected() const {
        return _isConnected;
    }
}