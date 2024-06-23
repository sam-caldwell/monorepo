/**
 * @name ArbitraryKvBtree/src/getUint.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "findNode.cpp"

/**
 * @name getUint
 * @brief Return a uint value for a matching criteria.
 *
 * @throws runtime_error exception if node does not exist.
 *
 * @param criteria string
 * @return uint
 */
inline uint ArbitraryKvBtree::getUint(string criteria) {
    lock->wait()->on();
    uint result;
    ArbitraryKvBtNode *here = findNode(criteria, false);
    if (here)
        result = here->getUint(_strongTyping);
    else {
        lock->off();
        throw out_of_range("node (" + criteria + ") not found");
    }
    lock->off();
    return result;
}