/**
 * @name ArbitraryKvList/src/GetFloat.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getFloat
 * @brief Return the current node (float).
 *        Root node will be left on this node.
 * @warning will throw runtime_error exception if type mismatch and strongTyping enabled
 *
 * @param string key
 * @return float
 */
float ArbitraryKvList::getFloat(string key) {
    lock->wait()->on();
    if (search(key)) {
        float result = root->getFloat(_strongTyping);
        lock->off();
        return result;
    } else {
        lock->off();
        throw NotFound();
    }
}
