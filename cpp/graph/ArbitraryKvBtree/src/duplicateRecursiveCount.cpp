/**
 * @name ArbitraryKvBtree/src/duplicateRecursiveCount.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name duplicateRecursiveCount
 * @brief Recursively calculate duplicate count.
 *
 * @param *n ArbitraryKvBtNode pointer
 * @return int
 */
int ArbitraryKvBtree::duplicateRecursiveCount(ArbitraryKvBtNode *n) {
    if (n && n->root && (n->key() == n->root->key()))
        return (1 +
                duplicateRecursiveCount(n->left) +
                duplicateRecursiveCount(n->right));
    else
        return 0;
}
