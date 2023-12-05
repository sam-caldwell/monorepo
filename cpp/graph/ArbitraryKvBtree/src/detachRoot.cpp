/**
 * @name ArbitraryKvBtree/src/detachRoot.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 */
#ifndef detachRoot_H
#define detachRoot_H

/**
 * @name detachRoot
 * @brief This will detach a given node from it's parent node
 *        (if any) and return the root node pointer.
 *
 * @param *n ArbitraryKvBtNode pointer.
 * @return ArbitraryKvBtNode pointer.
 */
inline ArbitraryKvBtNode *detachRoot(ArbitraryKvBtNode *n) {
    ArbitraryKvBtNode *result = NULL;
    if (n && n->root) {
        result = n->root;
        if (n->root->left == n) n->root->left = NULL;
        if (n->root->right == n) n->root->right = NULL;
        n->root = NULL;
    }
    return result;
}

#endif /*detachRoot_H*/
