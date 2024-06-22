/**
 * @name ArbitraryKvList/src/RemoveKeys.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name remove [keys]
 * @brief Delete a node with the given key.
 *        This will delete all matching keys.
 *
 * @param key string
 * @return bool
 */
bool ArbitraryKvList::remove(string key) {
    lock->wait()->on();
    bool result = true;
    while (search(key))
        result = result && remove();
    lock->off();
    return result;
}
