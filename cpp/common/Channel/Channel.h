/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 * @brief Implement a go-like message queue for passing data
 * between threads safely.  This implements both a
 * boundless and bounded (buffered) Channel implementation
 * with an iterator for use in looping.
 */

#ifndef Channel_H
#define Channel_H

#include "../exceptions/exceptions.h"
#include "../SimpleLock/src/SimpleLock.h"
#include <queue>

#include <queue>
#include <mutex>
#include <condition_variable>
#include <iostream>
#include <stdexcept>

template<typename T>
class Channel {
public:
    Channel(size_t sizeLimit = 0) : sizeLimit(sizeLimit), closed(false) {}

    ~Channel() {
        close(); // Ensure channel is closed and resources are cleaned up
    }

    void operator<<(const T &value) {
        std::unique_lock<std::mutex> lock(mtx);
        if (closed) {
            throw std::runtime_error("Channel closed");
        }

        if (sizeLimit > 0) {
            cv_push.wait(lock, [this]() { return queue.size() < sizeLimit; });
        }

        queue.push(value);
        cv_pop.notify_one();
    }

    bool operator>>(T &value) {
        std::unique_lock<std::mutex> lock(mtx);
        cv_pop.wait(lock, [this]() { return !queue.empty() || closed; });

        if (!queue.empty()) {
            value = queue.front();
            queue.pop();
            cv_push.notify_one();
            return true;
        }
        return false;
    }

    void close() {
        std::unique_lock<std::mutex> lock(mtx);
        closed = true;
        cv_pop.notify_all();
    }

    bool is_closed() const {
        std::unique_lock<std::mutex> lock(mtx);
        return closed;
    }

    void hardLimit(size_t sz) {
        sizeLimit = sz;
    }

private:
    std::queue<T> queue;
    mutable std::mutex mtx;  // mutable to allow locking in const member function
    std::condition_variable cv;
    bool closed;
    std::condition_variable cv_push; // For producer waiting on size limit
    std::condition_variable cv_pop;  // For consumer waiting on non-empty queue
    size_t sizeLimit;
};

#endif  /** Channel_H */