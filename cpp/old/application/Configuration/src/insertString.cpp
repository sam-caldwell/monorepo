/**
 * @name Configuration/src/insertString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */




/**
 * @name insert[string]
 * @brief insert new node with a string
 *
 * @param key string
 * @param n string
 * @return bool
 */
inline bool Configuration::insert(string key, string n) {
    return tree->insert(key, n);
}