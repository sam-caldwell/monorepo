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

#include "../../exceptions/exceptions.h"
#include "../../../common/SimpleLock/src/SimpleLock.h"
#include <queue>

using namespace std;

template<class ElementType>
class Channel {
public:
    using type = ElementType;

    /**
     * @name Class constructor.
     * @brief Setup our locks, limits and initial state
     * @note if capacity<=0, capacity will be unlimited.
     *
     * @param capacity
     */
    Channel(int capacity = 0) {
        writeLock = new SimpleLock(SIMPLE_LOCK_DEFAULT_TIMEOUT);
        readLock = new SimpleLock(SIMPLE_LOCK_DEFAULT_TIMEOUT);
        queueCapacity = (capacity <= 0) ? 0 : capacity; // zero (0) is unlimited.
        _close = false;
        r_block_backoff = 1;
        w_block_backoff = 1;
    };

    /**
     * @name Class destructor
     * @brief Clear locks, set close then terminate.
     */
    ~Channel() {
        writeLock->wait()->on(); //Ensure we are locked until we can close.
        _close = true; // ensure we are closed so we can safely kill things.
        writeLock->off(); //Disable our SimpleLock.
    }

    /**
     * @name Disabled class constructor
     * @brief Disabling to prevent move/copy of channel.
     */
    Channel(const Channel &) = delete;

    /**
     * @name Disabled class constructor
     * @brief Disabling to prevent move/copy of channel
     */
    Channel(Channel &&) = delete;

    /**
     * @name Channel assignment operator
     * @brief Disabling to prevent move/copy of channel/
     */
    Channel &operator=(const Channel &) = delete;

    /**
     * @name Channel assignment operator
     * @brief Disabling to prevent move/copy of channel/
     */
    Channel &operator=(Channel &&) = delete;

    /**
     * @name push()
     * @brief pushes a new element onto the channel.
     *
     * @param input ElementType
     */
    void push(ElementType input) {
        if (closed()) throw closedChannelException();
        writeLock->wait()->on();
        if (queueCapacity > 0) {
            do {
                this_thread::sleep_for(std::chrono::milliseconds(w_block_backoff++));
                if (w_block_backoff >= w_backoff_max) w_block_backoff /= 2;
            } while (data.size() >= queueCapacity);
            w_block_backoff = w_backoff_min;
        }
        data.push(input);
        writeLock->off();
    }

    /**
     * @name pop()
     * @brief pop the next element off the channel.
     *
     * @return ElementType
     */
    ElementType pop() {
        readLock->wait()->on();
        while (data.empty()) {
            if (r_block_backoff > r_backoff_limit)
                r_block_backoff = 1;
            this_thread::sleep_for(std::chrono::milliseconds(r_block_backoff));
            r_block_backoff *= 2;
        }
        r_block_backoff /= 2;
        ElementType output = data.front();
        this_thread::sleep_for(1ms);
        data.pop();
        readLock->off();
        return output;
    }

    /**
     * @name Channel size getter
     * @brief Return number of elements in the channel
     * @return int
     */
    inline int size() { return data.size(); };

    /**
     * @name Channel empty
     * @brief return whether the channel is empty.
     * @return bool
     */
    inline bool empty() { return data.empty(); };

    /**
     * @name close channel
     * @brief set the channel to close to prevent further writes.
     * @note any further writes will result in a closeChannelException.
     * @return bool
     */
    inline bool close() { return _close = true; };

    /**
     * @name is closed getter
     * @brief return whether or not the channel is closed.
     * @return bool
     */
    inline bool closed() { return _close; };

private:

    SimpleLock *writeLock;

    SimpleLock *readLock;

    queue <ElementType> data;

    bool _close;
    //bool _r_block;
    //bool _w_block;

    int queueCapacity;
    const int r_backoff_limit = 1024;
    int r_block_backoff;
    const int w_backoff_min = 1;
    const int w_backoff_max = 1024;
    int w_block_backoff;

};

#endif  /** Channel_H */