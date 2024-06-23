/**
 * @name types/src/ConfigStateMap/validate.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name validate
 * @brief execute the validator for a given lhs string against a given input
 *        and return the boolean result.
 *
 * @param lhs LhsString (the map record referencing the validator.
 * @param input string (the value to validate)
 * @return bool (operation result)
 */
inline bool ConfigStateMap::validate(LhsString lhs, string input) {
    ConfigStateMapValidator func;
    try {
        return ValidatorMap.at(lhs);
    } catch (out_of_range &e) {
        return true; //default validation is true (if no validator was specified.
    }

}