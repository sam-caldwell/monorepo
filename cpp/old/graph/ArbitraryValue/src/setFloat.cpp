/**
 * @name ArbitraryValue/src/setFloat.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */
#include "memoryCleanup.cpp"



/**
 * @name setFloat
 * @brief Stores the given value (v) as byte data
 *
 * @param float v
 * @return bool
 */
bool ArbitraryValue::setFloat(float v) {
    lock->wait()->on();
    try {
        memoryCleanup();
        type = Float;
        sz = sizeof(float);
        value = (float *) malloc(sz);
        if (!value)
            throw ValueAssignmentError("Memory allocation failed.");
        *((float *) value) = v;
        lock->off();
        return true;
    } catch (exception e) {
        lock->off();
        throw ValueAssignmentError(e.what());
    }
}