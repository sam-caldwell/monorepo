/**
 * @name ArbitraryKvList/src/GetNodePointer.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getNodePointer
 * @brief Get Node pointer for given key
 *
 * @param key string
 * @return ArbitraryKvNode* pointer (null if key does not exist).
 */
ArbitraryKvNode *ArbitraryKvList::getNodePointer(string key) {
    bool locked_locally = false;
    if (!lock->check()) { //Is there a SimpleLock already set?
        lock->on();  // No: Set a SimpleLock only if not set already.
        locked_locally = true;
    }
    ArbitraryKvNode *node = (search(key)) ? root : NULL;
    if (locked_locally) lock->off();
    return node;
}
