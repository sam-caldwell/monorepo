/**
 * @name ArbitraryKvBtree/src/insertInt.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "insertNode.cpp"

/**
 * @name insert [int]
 * @brief Insert int node.
 *
 * @param key string
 * @param value int
 * @return bool
 */
inline bool ArbitraryKvBtree::insert(string key, int v) {
    return insertNode(new ArbitraryKvBtNode(key, v));
}
