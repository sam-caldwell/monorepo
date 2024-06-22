/**
 * @name SimpleLock/src/check.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

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
