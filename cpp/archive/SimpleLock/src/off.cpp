/**
 * @name SimpleLock/src/off.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "SimpleLock.h"

using namespace std;

/**
 * @name off
 * @brief disable the SimpleLock.
 *
 * @return SimpleLock* pointer
 */
inline SimpleLock *SimpleLock::off() {
    lock = false;
    return this;
}
