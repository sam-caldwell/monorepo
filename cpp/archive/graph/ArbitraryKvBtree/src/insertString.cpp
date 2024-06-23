/**
 * @name ArbitraryKvBtree/src/insertString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "insertNode.cpp"

/**
 * @name insert [string]
 * @brief Insert string node.
 *
 * @param key string
 * @param value string
 * @return bool
 */
inline bool ArbitraryKvBtree::insert(string key, string v) {
    return insertNode(new ArbitraryKvBtNode(key, v));
}
