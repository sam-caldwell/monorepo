/**
 * @name Configuration/src/getUint.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getUint
 * @brief return the unsigned-integer (uint) value of a given node identified by key.
 *
 * @throws out_of_range exception on invalid key
 *
 * @param key string
 * @return uint
 */

inline uint Configuration::getUint(string key) {
    try {
        return tree->getUint(key);
    } catch (out_of_range &e) {
        throw out_of_range("Configuration::getUint() error:" + string(e.what()));
    }
}