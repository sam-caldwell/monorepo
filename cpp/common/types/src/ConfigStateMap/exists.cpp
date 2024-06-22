/**
 * @name types/src/exists.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name exists
 * @brief return true/false whether lhs string exists.
 *
 * @param key LhsString
 * @return bool (true: exists, false: not exists)
 */
inline bool ConfigStateMap::exists(LhsString key) {
    try {
        return (StringMap.at(key).length() > 0);
    } catch (out_of_range) {
        return false;
    }
}
