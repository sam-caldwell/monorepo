/**
 * @name ArbitraryValue/src/setInt.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */
#include "memoryCleanup.cpp"



/**
 * @name setInt
 * @brief Stores the given value (v) as byte data
 *
 * @param int v
 * @return bool
 */
bool ArbitraryValue::setInt(int v) {
    lock->wait()->on();
    memoryCleanup();
    type = Int;
    sz = sizeof(int);
    value = (int *) malloc(sz);
    if (!value)
        throw ValueAssignmentError("Memory allocation failed.");
    *((int *) value) = v;
    lock->off();
    return true;
}
