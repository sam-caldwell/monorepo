/**
 * @name ArbitraryValue/src/setUint.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */
#include "memoryCleanup.cpp"



/**
 * @name setUint
 * @brief Stores the given value (v) as byte data
 *
 * @param uint *v
 * @return bool
 */
bool ArbitraryValue::setUint(uint v) {
    lock->wait()->on();
    try {
        memoryCleanup();
        type = Uint;
        sz = sizeof(uint);
        value = (uint *) malloc(sz);
        if (!value)
            throw ValueAssignmentError("Memory allocation failed.");
        *((uint *) value) = v;
        lock->off();
        return true;
    } catch (exception e) {
        lock->off();
        throw ValueAssignmentError(e.what());
    }
}