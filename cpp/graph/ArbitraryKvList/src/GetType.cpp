/**
 * @name ArbitraryKvList/src/GetType.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getType
 * @brief returns the ValueTypes type of the named node.
 *
 * @warning will throw NotFound() exception if key not found
 *
 * @param string key
 * @return ValueTypes
 */
ValueTypes ArbitraryKvList::getType(string key) {
    lock->wait()->on();
    if (search(key)) {
        string result = root->getType();
        lock->off();
        return result;
    } else {
        lock->off();
        throw out_of_range("not found");
    }
}

/**
 * @name getType
 * @brief returns the ValueTypes type of the current node.
 *
 * @return ValueTypes
 */
ValueTypes ArbitraryKvList::getType() {
    lock->wait()->on();
    if(root) {
        ValueTypes result = root->getType();
        lock->off();
        return result;
    } else {
        throw out_of_range("null list");
    }
}