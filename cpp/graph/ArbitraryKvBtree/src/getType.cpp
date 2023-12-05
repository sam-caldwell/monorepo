/**
 * @name ArbitraryKvBtree/src/getType.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "findNode.cpp"


/**
 * @name getType
 * @brief Return the ValueTypes value of a given node.
 *
 * @throws out_of_range exception if node does not exist.
 *
 * @param criteria string
 * @return ValueTypes
 */
inline ValueTypes ArbitraryKvBtree::getType(string criteria) {
    lock->wait()->on();
    ValueTypes result = Null;
    ArbitraryKvBtNode *here = findNode(criteria, false);
    if (here) {
        result = here->getType();
        lock->off();
    } else {
        lock->off();
        throw out_of_range("node ('" + criteria + "') not found");
    }
    return result;
}
