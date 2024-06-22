/**
 * @name networking/TcpCommon/src/client/startListen.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

namespace networking {
    /**
     * @name startListen
     * @brief span a listener thread.
     */
    void Client::startListen() {
        setConnected(true);
        _receiveThread = new std::thread(&Client::receiveTask, this);
    }
}