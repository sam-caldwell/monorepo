/**
 * @name ArbitraryKvBtree/src/getFloat.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "findNode.cpp"

/**
 * @name getFloat
 * @brief Return a float value for a matching criteria.
 *
 * @throws runtime_error exception if node does not exist.
 *
 * @param criteria string
 * @return float
 */
inline float ArbitraryKvBtree::getFloat(string criteria) {
    lock->wait()->on();
    float result;
    ArbitraryKvBtNode *here = findNode(criteria, false);
    if (here)
        result = here->getFloat(_strongTyping);
    else {
        lock->off();
        throw out_of_range("node (" + criteria + ") not found");
    }
    lock->off();
    return result;
}