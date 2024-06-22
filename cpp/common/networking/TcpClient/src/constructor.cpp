/**
 * @name networking/TcpClient/src/constructor.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
namespace networking {
    /**
     * @name class constructor
     * @brief set initial state.
     */
    TcpClient::TcpClient() {
        _isConnected = false;
        _isClosed = true;
    }
}