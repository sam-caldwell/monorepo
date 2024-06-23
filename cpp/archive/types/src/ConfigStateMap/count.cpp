/**
 * @name types/src/count.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name count
 * @brief Return the number of records in the map.
 *
 * @return int
 */
int ConfigStateMap::count() {
    return StringMap.size();
}
