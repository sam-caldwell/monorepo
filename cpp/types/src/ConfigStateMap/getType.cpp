/**
 * @name types/src/ConfigStateMap/getType.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getType
 * @brief Fetch a ValueType for a given mapping.
 *
 * @throws out_of_range if key not found.
 *
 * @param lhs LhsString
 * @return ValueTypes
 */
inline ValueTypes ConfigStateMap::getType(LhsString lhs) {
    try {
        return ValueTypeMap.at(StringMap.at(lhs));
    } catch (out_of_range) {
        throw out_of_range("ConfigStateMap::getType() encountered 'key not found' (key:" + lhs + ")");
    }
}
