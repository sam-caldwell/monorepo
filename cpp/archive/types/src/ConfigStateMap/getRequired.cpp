/**
 * @name types/src/ConfigStateMap/getRequired.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getRequired
 * @brief return required field for a given LhsString (key).
 *
 * @throws out_of_range if key not found.
 *
 * @param key LhsString
 * @return bool (required field value)
 */
inline bool ConfigStateMap::getRequired(LhsString key) {
    try {
        return RequiredMap.at(key);
    } catch (out_of_range) {
        throw out_of_range("ConfigStateMap::getRequired() encountered 'key not found' (key:" + key + ")");
    }
}
