/**
 * @name networking/TcpCommon/src/client/isConnected.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
namespace networking {
    /**
     * @name isConnected
     * @brief determines if the connection is connected
     * @return bool
     */
    inline bool Client::isConnected() const {
        return _isConnected;
    }
}