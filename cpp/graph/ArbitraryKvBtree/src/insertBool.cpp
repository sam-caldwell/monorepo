/**
 * @name ArbitraryKvBtree/src/insertBool.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "insertNode.cpp"

/**
 * @name insert [bool]
 * @brief Insert boolean node.
 *
 * @param key string
 * @param value bool
 * @return bool
 */
inline bool ArbitraryKvBtree::insert(string key, bool v) {
    return insertNode(new ArbitraryKvBtNode(key, v));
}
