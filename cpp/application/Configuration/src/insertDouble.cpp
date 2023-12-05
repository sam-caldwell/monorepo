/**
 * @name Configuration/src/insertDouble.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */




/**
 * @name insert[double]
 * @brief insert new node with a double-precision floating-point value.
 *
 * @param key string
 * @param n double
 * @return bool
 */
inline bool Configuration::insert(string key, double n) {
    return tree->insert(key, n);
}