/**
 * @name networking/TcpCommon/src/client/setConnected.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

namespace networking {
    /**
     * @name setConnected
     * @brief set the connected status.
     *
     * @param flag
     */
    void networking::Client::setConnected(bool flag) {
        _isConnected = flag;
    }
}