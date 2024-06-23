/**
 * @name networking/TcpCommon/src/client/terminateReceiveThread.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

namespace  networking{
    /**
     * @name terminateReceiveThread
     * @brief terminate the receiver thread and set the connection state to false.
     */
    void Client::terminateReceiveThread() {
        setConnected(false);
        if (_receiveThread) {
            _receiveThread->join();
            delete _receiveThread;
            _receiveThread = nullptr;
        }
    }
}