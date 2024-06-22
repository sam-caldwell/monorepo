/**
 * @name ArbitraryKvBtree/src/getDouble.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "findNode.cpp"


/**
 * @name getDouble
 * @brief Return a double value for a matching criteria.
 *
 * @throws runtime_error exception if node does not exist.
 *
 * @param criteria string
 * @return double
 */
inline double ArbitraryKvBtree::getDouble(string criteria) {
    lock->wait()->on();
    double result;
    ArbitraryKvBtNode *here = findNode(criteria, false);
    if (here)
        result = here->getDouble(_strongTyping);
    else {
        lock->off();
        throw out_of_range("node (" + criteria + ") not found");
    }
    lock->off();
    return result;
}