/**
* @name networking/TcpClient/src/subscribe.cpp
* @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
* @author Sam Caldwell <mail@samcaldwell.net>
*/
namespace networking {
    /**
     * @name subscribe
     * @brief
     *
     * @param observer
     */
    void TcpClient::subscribe(const ClientObserver &observer) {
        std::lock_guard <std::mutex> lock(_subscribersMtx);
        _subscibers.push_back(observer);
    }
}