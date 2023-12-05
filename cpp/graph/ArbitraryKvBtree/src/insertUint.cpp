/**
 * @name ArbitraryKvBtree/src/insertUint.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "insertNode.cpp"

/**
 * @name insert [uint]
 * @brief Insert uint node.
 *
 * @param key string
 * @param value uint
 * @return bool
 */
inline bool ArbitraryKvBtree::insert(string key, uint v) {
    return insertNode(new ArbitraryKvBtNode(key, v));
}
