/**
 * @name ArbitraryValue/src/setBool.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */
#include "memoryCleanup.cpp"




/**
 *  @name setBool()
 *  @brief Stores the given value (v) as byte data
 *
 * @param bool *v
 * @return bool
 */
bool ArbitraryValue::setBool(bool v) {
    lock->wait()->on();
    if (v) memoryCleanup();
    type = Bool;
    sz = sizeof(char);
    value = (char *) malloc(sz);
    if (!value)
        throw ValueAssignmentError("Memory allocation failed.");
    *((char *) value) = (v) ? v_true : v_false;
    lock->off();
    return true;
}
