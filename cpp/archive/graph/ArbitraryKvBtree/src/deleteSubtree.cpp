/**
 * @name ArbitraryKvBtree/src/deleteSubtree.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 * @brief This file contains the node removal (delete) functionality.
 */
#include "seekRoot.cpp"
#include "findNode.cpp"

/**
 * @name deleteSubtree
 * @brief Remove a node and its subtree where the relative root has a given key name
 *
 * @param *n ArbitraryKvBtNode pointer
 * @return bool
 */

bool ArbitraryKvBtree::deleteSubtree(ArbitraryKvBtNode *n) {
    /**
     * Perform a null pointer check, return false on null.
     */
    if (!n) return false;
    /**
     * Get the (detached) parent node.
     * At this point, node n is orphaned.
     */
    ArbitraryKvBtNode *parent = detachRoot(n);
    /**
     * Recursively delete the subtrees of n.
     */
    while (deleteSubtree(n->left) || deleteSubtree(n->right)) {}
    /**
     * Having deleted all n->left and n->right nodes, we now
     * delete node (n).
     */
    if (n) delete n;
    return true;
}