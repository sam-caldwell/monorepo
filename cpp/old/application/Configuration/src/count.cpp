/**
 * @name Configuration/src/count.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name count
 * @brief return the number of configuration data objects
 *
 * @return int
 */
inline int Configuration::count() {
    return tree->count();
}