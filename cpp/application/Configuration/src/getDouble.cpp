/**
 * @name Configuration/src/getDouble.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getDouble
 * @brief return the double-precision floating-point
 *        value of a given node identified by a key.
 *
 * @throws out_of_range exception on invalid key
 *
 * @param key string
 * @return double
 */
inline double Configuration::getDouble(string key) {
    try {
        return tree->getDouble(key);
    } catch (out_of_range &e) {
        throw out_of_range("Configuration::getDouble() error:" + string(e.what()));
    }
}