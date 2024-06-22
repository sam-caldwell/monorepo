/**
 * @name Configuration/src/insertUint.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */




/**
 * @name insert[uint]
 * @brief insert new node with an unsigned integer (uint)
 *
 * @param key string
 * @param n uint
 * @return bool
 */
inline bool Configuration::insert(string key, uint n) {
    return tree->insert(key, n);
}