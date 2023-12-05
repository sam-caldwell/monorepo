/**
 * @name ArbitraryValue/src/safeEraseMemory.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */

#ifndef safeEraseMemory_H
#define safeEraseMemory_H

/**
 * @name safeEraseMemory()
 * @brief overwrite the value memory buffer of the given
 *        size, then ensure the buffer is null-terminated.
 *
 * @param *p void pointer
 * @param sz int
 */
void ArbitraryValue::safeEraseMemory(void *p, int sz) {
    char *dst = (char *) value;
    for (; sz; sz--)
        *(dst++) = rand() % 26 + 'a';
    *(dst) = '\0';
}

#endif /*safeEraseMemory_H*/