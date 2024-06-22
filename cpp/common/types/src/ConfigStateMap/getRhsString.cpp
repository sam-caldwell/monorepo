/**
 * @name types/src/ConfigStateMap/getRhsString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getRhsString
 * @brief Given a mapping of an left-hand (lhs) string to a right-hand (rhs) string
 *        this method returns the right-hand (rhs) string associated with its lhs.
 *
 * @throws out_of_range if key not found.
 *
 * @param lhs LhsString
 * @return string (rhs)
 */
inline string ConfigStateMap::getRhsString(LhsString lhs) {
    try {
        return StringMap.at(lhs);
    } catch (out_of_range) {
        throw out_of_range("ConfigStateMap::getRhsString() encountered 'key not found' (key:" + lhs + ")");
    }
}