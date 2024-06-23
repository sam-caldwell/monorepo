/**
 * @name SimpleLock/src/constructors.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "SimpleLock.h"

using namespace std;

/**
 * @name Default Class Constructor.
 * @brief default state SimpleLock is disabled.
 */
SimpleLock::SimpleLock() {
    lock = false;
    SetTimeout(SIMPLE_LOCK_DEFAULT_TIMEOUT, SIMPLE_LOCK_DEFAULT_INTERVAL);
//    cout << "timeout: " << timeout << endl;
}

/**
 * @name Class Constructor.
 * @brief default state SimpleLock is disabled.
 *
 * @param uint timeout_ms - timeout for wait()
 * @param uint interval_ms - timeout check interval (default: SIMPLE_LOCK_DEFAULT_INTERVAL)
 */
SimpleLock::SimpleLock(uint timeout_ms, uint interval_ms = SIMPLE_LOCK_DEFAULT_INTERVAL) {
    SetTimeout(timeout_ms, interval_ms);
    lock = false;
}


