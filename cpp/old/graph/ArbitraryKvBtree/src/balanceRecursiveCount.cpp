/**
 * @name ArbitraryKvBtree/src/balanceRecursiveCount.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name balanceRecursiveCount
 * @brief Recursively calculate balance score for the tree.  left: -1, right: +1
 *
 * @param *n ArbitraryKvBtNode pointer
 * @return int
 */
int ArbitraryKvBtree::balanceRecursiveCount(ArbitraryKvBtNode *n, int d = 0) {
    if (n)
        return d + balanceRecursiveCount(n->left, -1) + balanceRecursiveCount(n->right, +1);
    else
        return 0;
}
