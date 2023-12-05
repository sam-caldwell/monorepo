/**
 * @name Configuration/src/getInt.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getInt
 * @brief return the integer value of a given node identified by key.
 *
 * @throws out_of_range exception on invalid key
 *
 * @param key string
 * @return int
 */
inline int Configuration::getInt(string key) {
    try {
        return tree->getInt(key);
    } catch (out_of_range &e) {
        throw out_of_range("Configuration::getInt() error:" + string(e.what()));
    }
}