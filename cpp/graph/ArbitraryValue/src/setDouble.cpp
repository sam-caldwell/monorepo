/**
 * @name ArbitraryValue/src/setDouble.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */
#include "memoryCleanup.cpp"



/**
*  @name setDouble
*  @brief Stores the given value (v) as byte data
*
* @param double v
* @return bool
*/
bool ArbitraryValue::setDouble(double v) {
    lock->wait()->on();
    try {
        memoryCleanup();
        type = Double;
        sz = sizeof(double);
        value = (double *) malloc(sz);
        if (!value)
            throw ValueAssignmentError("Memory allocation failed.");
        *((double *) value) = v;
        lock->off();
        return true;
    } catch (exception e) {
        lock->off();
        throw ValueAssignmentError(e.what());
    }
}
