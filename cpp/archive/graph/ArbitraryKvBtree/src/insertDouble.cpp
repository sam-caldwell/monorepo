/**
 * @name ArbitraryKvBtree/src/insertDouble.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "insertNode.cpp"

/**
 * @name insert [double]
 * @brief Insert double node.
 *
 * @param key string
 * @param value double
 * @return bool
 */
inline bool ArbitraryKvBtree::insert(string key, double v) {
    return insertNode(new ArbitraryKvBtNode(key, v));
}
