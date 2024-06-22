/**
 * @name Configuration/src/getString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getString
 * @brief return the string value of a given node identified by key.
 *
 * @throws out_of_range exception on invalid key
 *
 * @param key string
 * @return string
 */
inline string Configuration::getString(string key) {
    try {
        return tree->getString(key);
    } catch (out_of_range &e) {
        throw out_of_range("Configuration::getString() error:" + string(e.what()));
    }
}