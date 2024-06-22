/**
 * @name ArbitraryKvList/src/Key.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name key
 * @brief Get the key value of the current (root) node.
 *
 * @return string
 */
string ArbitraryKvList::key() {
    //ToDo: implement StackedLocks
    string key = (root) ? root->key() : "";
    //SimpleLock->off();
    return key;
}
