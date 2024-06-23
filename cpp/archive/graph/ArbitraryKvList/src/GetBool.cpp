/**
 * @name ArbitraryKvList/src/GetBool.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getBool
 * @brief Search for and return the given node as boolean.
 * @warning will throw runtime_error exception if type mismatch and strongTyping enabled
 *
 * @param key string
 * @return bool
 */
bool ArbitraryKvList::getBool(string key) {
    lock->wait()->on();
    if (search(key)) {
        bool result = root->getBool(_strongTyping);
        lock->off();
        return result;
    } else {
        lock->off();
        throw NotFound();
    }
}

/**
 * @name getBool
 * @brief return boolean value of current node.
 *
 * @warning will throw out_of_range exception if no node exists.
 * @warning will throw runtime_error exception if type mismatch and strongTyping enabled
 *
 * @return bool
 */
bool ArbitraryKvList::getBool() {
    lock->wait()->on();
    if (root) {
        bool result = root->getBool(_strongTyping);
        lock->off();
        return result;
    } else {
        throw out_of_range("null list");
    }
}