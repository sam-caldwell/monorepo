/**
 * ArbitraryKvList/src/ResetRight.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "resetNodeRight.cpp"

/**
 * @name resetRight
 * @brief Move root pointer to the right-most node.
 */
void ArbitraryKvList::resetRight() {
    bool local_lock = false;
    if (!lock->check()) {
        lock->on();
        local_lock = true;
    }
    root = resetNodeRight(root);
    lock->off();
}
