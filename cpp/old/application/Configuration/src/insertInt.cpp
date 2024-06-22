/**
 * @name Configuration/src/insertInt.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name insert[int]
 * @brief insert new node with an integer
 *
 * @param key string
 * @param n int
 * @return bool
 */
inline bool Configuration::insert(string key, int n) {
    return tree->insert(key, n);
}