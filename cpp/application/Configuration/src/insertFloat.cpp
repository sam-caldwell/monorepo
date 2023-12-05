/**
 * @name Configuration/src/insertFloat.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */




/**
 * @name insert[float]
 * @brief insert new node with a single-precision floating-point number
 *
 * @param key string
 * @param n float
 * @return bool
 */
inline bool Configuration::insert(string key, float n) {
    return tree->insert(key, n);
}