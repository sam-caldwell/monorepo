/**
 * @name ArbitraryKvList/src/GetString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getString
 * @brief Return the current node (string).
 *        Root node will be left on this node.
 * @warning will throw runtime_error exception if type mismatch and strongTyping enabled
 *
 * @param string key
 * @return string
 */
string ArbitraryKvList::getString(string key) {
    lock->wait()->on();
    if (search(key)) {
        string result = root->getString(_strongTyping);
        lock->off();
        return result;
    } else {
        lock->off();
        throw NotFound();
    }
}