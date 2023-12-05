/**
 * @name ArbitraryKvBtree/src/Exists.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "findNode.cpp"
/**
 * @name exists
 * @brief Return whether a given node exists.  This will NOT
 *        affect the node where root is pointing in the end.
 *
 * @param criteria string
 * @return bool
 */
inline bool ArbitraryKvBtree::exists(string criteria) {
    return findNode(criteria);
}
