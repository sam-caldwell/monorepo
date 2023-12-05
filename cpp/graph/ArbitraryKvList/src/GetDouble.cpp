/**
 * @name ArbitraryKvList/src/GetDouble.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getDouble
 * @brief Return the current node (double).
 *        Root node will be left on this node.
 * @warning will throw runtime_error exception if type mismatch and strongTyping enabled
 *
 * @param string key
 * @return double
 */
double ArbitraryKvList::getDouble(string key) {
    lock->wait()->on();
    if (search(key)) {
        double result = root->getDouble(_strongTyping);
        lock->off();
        return result;
    } else {
        lock->off();
        throw NotFound();
    }
}