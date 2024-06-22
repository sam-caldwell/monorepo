/**
 * @name types/src/ValueTypesToString.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name ValueTypeToString
 * @brief return a string representation of the ValueTypes value given
 *
 * @param v
 * @return string
 */
inline string ValueTypeToString(ValueTypes v) {
    switch (v) {
        case Null:
            return "Null";
        case Bool:
            return "Bool";
        case Double:
            return "Double";
        case Float:
            return "Float";
        case Int:
            return "Int";
        case String:
            return "String";
        case Uint:
            return "Uint";
        default:
            throw ("Unknown ValueTypes");
    }
}