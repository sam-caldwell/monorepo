/**
 * @name ArbitraryValue/src/destructors.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */

#include "memoryCleanup.cpp"



/**
 * Class destructor.
 *   - Safely erase our memory
 */
ArbitraryValue::~ArbitraryValue() {
    lock->wait()->on();
    memoryCleanup();
    lock->off();
    delete lock;
}