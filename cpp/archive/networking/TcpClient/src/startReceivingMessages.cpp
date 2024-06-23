/**
 * @name networking/TcpClient/src/startReceivingMessages.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
namespace  networking {
/**
 * @name startReceivingMessages
 * @brief spawn the thread for receiving messages.
 */
    inline void TcpClient::startReceivingMessages() {
        _receiveTask = new std::thread(&TcpClient::receiveTask, this);
    }
}