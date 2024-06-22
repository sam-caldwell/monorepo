/**
 * @name ArbitraryValue/src/getType.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */

/**
 * @name getType
 * @brief return the ValueType of the given object
 *
 * @return ValueTypes
 */
ValueTypes inline ArbitraryValue::getType() {
    return type;
}