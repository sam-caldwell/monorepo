/**
 * @name ArbitraryValue/src/memoryCleanup.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */
#ifndef memoryCleanup_H
#define memoryCleanup_H

#include "safeEraseMemory.cpp"

/**
 * @name memoryCleanup()
 * @brief free any allocated memory object
 * @note use before assigning a new value to avoid leaks.
 */
inline void ArbitraryValue::memoryCleanup() {
    bool releaseLock = false;
    if (!lock->check()) {
        releaseLock = true;
        lock->on();
    } // should guarantee a SimpleLock if not already present
    if (value) {
        safeEraseMemory(value, sz);
        free(value);
    }
    sz = 0;
    type = Null;
    if (releaseLock)
        lock->off(); //if we set a SimpleLock here, release the SimpleLock
}

#endif /*memoryCleanup_H*/