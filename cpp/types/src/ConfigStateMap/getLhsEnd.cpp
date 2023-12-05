/**
 * @name types/src/ConfigStateMap/getLhsEnd.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getLhsEnd
 * @brief return an iterator to the LhsString pointing at the end of the map
 *
 * @return std::map<LhsString, RhsString>::iterator
 */
inline auto ConfigStateMap::getLhsEnd() {
    return StringMap.end();
}
