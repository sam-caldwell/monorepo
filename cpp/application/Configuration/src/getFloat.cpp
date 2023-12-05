/**
 * @name Configuration/src/getFloat.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getFloat
 * @brief return the single-precision floating-point
 *        value of a given node identified by key.
 *
 * @throws out_of_range exception on invalid key
 *
 * @param key string
 * @return float
 */
inline float Configuration::getFloat(string key) {
    try {
        return tree->getFloat(key);
    } catch (out_of_range &e) {
        throw out_of_range("Configuration::getFloat() error:" + string(e.what()));
    }
}