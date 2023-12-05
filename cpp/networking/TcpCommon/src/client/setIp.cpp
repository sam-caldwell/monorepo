/**
 * @name networking/TcpCommon/src/client/setIp.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

namespace networking {
/**
* @name setIp
* @brief set the ip address.
* @param ip string
*/
    inline void Client::setIp(const std::string &ip) {
        _ip = ip;
    }
}