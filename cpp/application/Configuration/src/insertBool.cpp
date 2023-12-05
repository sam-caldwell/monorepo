/**
 * @name Configuration/src/insertBool.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */




/**
 * @name insert[bool]
 * @brief insert new node with a boolean value
 *
 * @param key string
 * @param n bool
 * @return bool (true: success, false: operation failed)
 */
inline bool Configuration::insert(string key, bool n) {
    return tree->insert(key, n);
}