/**
 * @name ArbitraryKvBtree/src/isLocked.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name isLocked
 * @brief return tree SimpleLock state.
 * @return bool.
 */
inline bool ArbitraryKvBtree::isLocked() {
    return lock->check();
}