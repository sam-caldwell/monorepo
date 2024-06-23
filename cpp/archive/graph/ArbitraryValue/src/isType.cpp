/**
 * @name ArbitraryValue/src/isType.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */

/**
 * @name isType
 * @brief return true/false whether or not the class is of type t.
 *
 * @param ValueTypes t
 * @return bool
 */
bool ArbitraryValue::isType(ValueTypes t) {
    return getType() == t;
}