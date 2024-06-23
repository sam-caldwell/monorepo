/**
 * @name ArbitraryKvBtree/src/insertFloat.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "insertNode.cpp"

/**
 * @name insert [float]
 * @brief Insert float node.
 *
 * @param key string
 * @param value float
 * @return bool
 */
inline bool ArbitraryKvBtree::insert(string key, float v) {
    return insertNode(new ArbitraryKvBtNode(key, v));
}
