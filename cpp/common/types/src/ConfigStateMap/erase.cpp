/**
 * @name types/src/ConfigStateMap/erase.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name erase
 * @brief Remove all mappings for the given string.
 *
 * @param lhs LhsString
 * @return bool (operation result)
 */
bool ConfigStateMap::erase(LhsString lhs) {
    try {
        RhsString rhs = StringMap.at(lhs);
        ValueTypeMap.erase(rhs);
        StringMap.erase(lhs);
        RequiredMap.erase(lhs);
        ValidatorMap.erase(lhs);
    } catch (std::out_of_range) {
        return false;
    }
    return true;
}
