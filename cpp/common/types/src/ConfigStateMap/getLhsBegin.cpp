/**
 * @name types/src/ConfigStateMap/getLhsBegin.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getLhsBegin
 * @brief return an iterator to the LhsString pointing at the beginning of the map
 *
 * @return std::map<LhsString, RhsString>::iterator
 */
inline auto ConfigStateMap::getLhsBegin() {
    return StringMap.begin();
}
