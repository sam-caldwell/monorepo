/**
 * @name networking/TcpCommon/src/client/getIp.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
namespace networking {
    /**
     * @name getIp
     * @brief return the current Ip address string
     * @return string
     */
    inline std::string Client::getIp() const {
        return _ip;
    }
}