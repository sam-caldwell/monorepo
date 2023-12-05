/**
 * @name networking/TcpClient/src/connectTo.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
namespace networking {
    /**
     * @name terminateReceiveThread
     * @brief terminate the receiver threads
     */
    void TcpClient::terminateReceiveThread() {
        _isConnected = false;

        if (_receiveTask) {
            _receiveTask->join();
            delete _receiveTask;
            _receiveTask = nullptr;
        }
    }
}
