/**
 * @name ArbitraryKvBtree/src/getString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "findNode.cpp"


/**
 * @name getString
 * @brief Return a string value for a matching criteria.
 *
 * @throws runtime_error exception if node does not exist.
 *
 * @param criteria string
 * @return string
 */
inline string ArbitraryKvBtree::getString(string criteria) {
    lock->wait()->on();
    string result;
    ArbitraryKvBtNode *here = findNode(criteria, false);
    if (here)
        result = here->getString(_strongTyping);
    else {
        lock->off();
        throw out_of_range("node (" + criteria + ") not found");
    }
    lock->off();
    return result;
}