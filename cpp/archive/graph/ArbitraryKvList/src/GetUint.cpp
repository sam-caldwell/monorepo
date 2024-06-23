/**
 * @name ArbitraryKvList/src/GetUint.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getUint
 * @brief Return the current node (uint).
 *        Root node will be left on this node.
 * @warning will throw runtime_error exception if type mismatch and strongTyping enabled
 *
 * @param string key
 * @return uint
 */
uint ArbitraryKvList::getUint(string key) {
    lock->wait()->on();
    if (search(key)) {
        uint result = root->getUint(_strongTyping);
        lock->off();
        return result;
    } else {
        lock->off();
        throw NotFound();
    }
}
