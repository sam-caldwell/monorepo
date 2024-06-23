/**
 * @name SimpleLock/src/check.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "SimpleLock.h"

using namespace std;

/**
 * @name check
 * @brief return boolean (true: locked, false: unlocked) status
 *
 * @return bool
 */
inline bool SimpleLock::check() {
    return lock;
}
