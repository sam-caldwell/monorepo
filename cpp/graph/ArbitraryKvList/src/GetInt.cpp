/**
 * @name ArbitraryKvList/src/GetInt.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getInt
 * @brief Return the current node (int).
 *        Root node will be left on this node.
 * @warning will throw runtime_error exception if type mismatch and strongTyping enabled
 *
 * @param string key
 * @return int
 */
int ArbitraryKvList::getInt(string key) {
    lock->wait()->on();
    if (search(key)) {
        int result = root->getInt(_strongTyping);
        lock->off();
        return result;
    } else {
        lock->off();
        throw NotFound();
    }
}
