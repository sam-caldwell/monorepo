/**
 * @name ArbitraryKvList/src/ResetLeft.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "resetNodeLeft.cpp"

/**
 * @name resetLeft
 * @brief Move root pointer to the left-most node.
 */
void ArbitraryKvList::resetLeft() {
    bool local_lock = false;
    if (!lock->check()) {
        lock->on();
        local_lock = true;
    }
    root = resetNodeLeft(root);
    lock->off();
}
