/**
 * @name types/src/ConfigStateMap/rhsExists.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name rhsExists
 * @brief Does the rhs string exist?
 *
 * @param rhs RhsString
 * @return bool (operational result)
 */
bool ConfigStateMap::rhsExists(RhsString rhs) {
    try {
        ValueTypeMap.at(rhs);
        return true;
    } catch (out_of_range) {
        return false;
    }
}
