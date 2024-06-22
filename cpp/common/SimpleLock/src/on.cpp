/**
 * @name SimpleLock/src/on.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

using namespace std;

/**
 * @name on
 * @brief enable the SimpleLock.
 *
 * @return SimpleLock*
 */
inline SimpleLock *SimpleLock::on() {
    lock = true;
    return this;
}
