/**
 * @name SimpleLock/src/SetTimeout.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

using namespace std;

/**
 * @name SetTimeout
 * @brief sets the timeout (in seconds)
 *
 * @param int timeout_ms - timeout for wait()
 * @param int interval_ms - timeout check interval (default: SIMPLE_LOCK_DEFAULT_INTERVAL)
 * @return SimpleLock*
 */
SimpleLock *SimpleLock::SetTimeout(uint timeout_ms, uint interval_ms = SIMPLE_LOCK_DEFAULT_INTERVAL) {
    if (timeout_ms <= 0) {
        throw out_of_range("Timeout must be > 0");
    }
    if (interval_ms <= 0) {
        throw out_of_range("Timeout must be > 0");
    }
    timeout = timeout_ms;
    interval = interval_ms;
    return this;
}
