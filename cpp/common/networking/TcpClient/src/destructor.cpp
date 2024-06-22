/**
 * @name networking/TcpClient/src/destructor.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
namespace networking {
    /**
     * @name class destructor
     * @brief close any open connection.
     */
    TcpClient::~TcpClient() {
        close();
    }
}