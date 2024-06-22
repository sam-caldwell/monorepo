/**
 * @name ArbitraryKvBtree/src/countRecursiveCount.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name countRecursiveCount
 * @brief Recursively Count Nodes.
 *
 * @param *n ArbitraryKvBtNode pointer
 * @return int
 */
int ArbitraryKvBtree::countRecursiveCount(ArbitraryKvBtNode *n) {
    return (n) ? (
            1 +
            countRecursiveCount(n->left) +
            countRecursiveCount(n->right)
    ) : 0;
}
