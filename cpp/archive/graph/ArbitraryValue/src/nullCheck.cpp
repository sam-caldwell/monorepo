/**
 * @name ArbitraryValue/src/nullCheck.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */

/**
 * @name nullCheck
 * @brief throw exception if value is null.
 *
 * @param ValueTypes t
 */
void ArbitraryValue::nullCheck() {
    if (!value || (type == Null)) throw NullValue();
}
