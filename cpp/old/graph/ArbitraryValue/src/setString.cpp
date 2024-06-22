/**
 * @name ArbitraryValue/src/setString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */
#include "memoryCleanup.cpp"



/**
 * @name setString(*v))
 * @brief Stores the given value (v) (NULL))
 *
 * @param string v
 * @return bool
 */
bool ArbitraryValue::setString(string v) {
    lock->wait()->on();
    try {
        memoryCleanup();
        type = String;

        sz = v.size();

        value = (char *) malloc((sz + 1) * sizeof(char)); //+1 for null-termination.
        if (!value)
            throw ValueAssignmentError("Memory allocation failed.");
        for (int i = 0; i < sz; i++) {
            *(((char *) value) + i) = v.c_str()[i];
        }
        lock->off();
        return true;
    } catch (exception e) {
        lock->off();
        throw ValueAssignmentError(e.what());
    }
}