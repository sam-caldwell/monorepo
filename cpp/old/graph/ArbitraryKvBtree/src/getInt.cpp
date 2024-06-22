/**
 * @name ArbitraryKvBtree/src/getInt.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "findNode.cpp"


/**
 * @name getInt
 * @brief Return a int value for a matching criteria.
 *
 * @throws runtime_error exception if node does not exist.
 *
 * @param criteria string
 * @return int
 */
inline int ArbitraryKvBtree::getInt(string criteria) {
    lock->wait()->on();
    int result;
    ArbitraryKvBtNode *here = findNode(criteria, false);
    if (here)
        result = here->getInt(_strongTyping);
    else {
        lock->off();
        throw out_of_range("node (" + criteria + ") not found");
    }
    lock->off();
    return result;
}