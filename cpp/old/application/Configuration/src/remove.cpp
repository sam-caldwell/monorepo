/**
 * @name Configuration/src/remove.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name remove
 * @brief remove node using key
 *
 * @param s string
 * @return bool
*/
inline bool Configuration::remove(string key) {
    return tree->remove(key);
}
