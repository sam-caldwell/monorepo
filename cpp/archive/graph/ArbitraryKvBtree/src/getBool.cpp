/**
 * @name ArbitraryKvBtree/src/getBool.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "findNode.cpp"

/**
 * @name getBool
 * @brief Return a boolean value for a matching criteria.
 *
 * @throws runtime_error exception if node does not exist.
 *
 * @param criteria string
 * @return bool
 */
inline bool ArbitraryKvBtree::getBool(string criteria) {
    lock->wait()->on();
    bool result;
    ArbitraryKvBtNode *here = findNode(criteria, false);
    if (here)
        result = here->getBool(_strongTyping);
    else {
        lock->off();
        throw out_of_range("node (" + criteria + ") not found");
    }
    lock->off();
    return result;
}