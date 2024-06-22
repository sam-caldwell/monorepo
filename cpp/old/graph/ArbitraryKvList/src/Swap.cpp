/**
 * @name ArbitraryKvList/src/Swap.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "swapNodes.cpp"

/**
 * @name swap
 * @brief Swap two nodes (by key)
 *
 * @param key1 string
 * @param key2 string
 * @return bool
 */
bool ArbitraryKvList::swap(string key1, string key2) {
    lock->wait()->on();
    bool result = swapNodes(getNodePointer(key1), getNodePointer(key2));
    lock->off();
    return result;
}