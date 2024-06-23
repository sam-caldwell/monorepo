/**
 * @name ArbitraryKvBtree/src/deleteNode.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 * @brief This file contains the node removal (delete) functionality.
 */
#include "detachRoot.cpp"
#include "seekRoot.cpp"
#include "findNode.cpp"

/**
 * @name deleteNode
 * @brief Remove a single node with a matching key name.
 * @warning DO NOT PASS NULL NODE POINTERS (*n)
 *
 * @param *n ArbitraryKvBtNode pointer
 * @return bool
 */

bool ArbitraryKvBtree::deleteNode(ArbitraryKvBtNode *n) {
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
     * Detach left and right subtrees.
     */
    ArbitraryKvBtNode *lhs;
    ArbitraryKvBtNode *rhs;
    if (n) {
        lhs = n->left;
        if (lhs) lhs->root = NULL;
    }
    if (n) {
        rhs = n->right;
        if (rhs) lhs->root = NULL;
    }
    if (lhs) detachRoot(lhs);
    if (rhs) detachRoot(rhs);
    /**
     * delete the node (n)
     */
    if (n) {
        n->root = NULL;
        n->left = NULL;
        n->right = NULL;
        delete n;
    }
    /**
     * Attach the parent node to the lhs subtree.
     * If lhs does not exist, parent->left will
     * be NULL as well.
     */
    if (lhs) {
        if (parent) {
            if (parent->left == NULL) {
                if (lhs->key() < parent->key())
                    parent->left = lhs;
                else
                    parent->right = lhs;
            } else if (parent->right == NULL) {
                if (lhs->key() < parent->key())
                    parent->left = lhs;
                else
                    parent->right = lhs;
            } else {
                throw std::runtime_error("LOGIC ERROR: Parent should have either open left or right pointer");
            }
            /**
             * Backlink to parent.  This could be a null
             */
            lhs->root = parent;
        } else {
            /**
             * given no parent, lhs becomes the parent.
             */
            parent = lhs;
            if (lhs) lhs->root = NULL;
        }
    }
    /**
     * move parent to the right-most node of
     * our new attached subtree if there is
     * a right-hand subtree (RHS).
     */
    if (rhs) {
        /**
         * Shift lhs pointer as far right as possible.
         * The right most node of LHS will still be
         * less than RHS.
         */
        while (parent && (parent->left || parent->right)) {
            if (rhs->key() < parent->key())
                parent = parent->left;
            else if(parent->key() > rhs->key())
                parent = parent->right;
            else
                while(parent->right) parent=parent->right;
        }
        parent->right=rhs;
        /**
         * Attach rhs subtree
         */
        parent->right = rhs;
        rhs->root = parent;
    }
    /**
     * Return true
     */
    root = (root) ? seekRoot(root) : seekRoot(parent);
    return true;
}