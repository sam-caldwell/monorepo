/**
 * @name Configuration/src/getBool.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getBool
 * @brief return the string value of a given node identified by (rhs) key.
 *
 * @throws out_of_range exception on invalid key
 *
 * @param key string
 * @return bool
 */
inline bool Configuration::getBool(string key) {
    try {
        return tree->getBool(key);
    } catch (out_of_range &e) {
        throw out_of_range("Configuration::getBool() error:" + string(e.what()));
    }
}