/**
 * @name ArbitraryKvBtree/src/removeNode.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "detachRoot.cpp"

/**
 * @name removeNode
 * @brief Remove a given ArbitraryKvBtNode object.
 * @warning SimpleLock by caller is assumed.
 * @warning this is going to create an unbalanced tree in most cases.
 *
 * @param *n ArbitraryKvBtNode pointer
 * @return bool
 */
//bool ArbitraryKvBtree::removeNode(ArbitraryKvBtNode *n, bool removeSubtree = true) {
//    DEBUG("ArbitraryKvBtree::removeNode() starting")
//    bool result = false;
//    /**
//     * parent_tree - deleted_node (n) parent (dangling end of the root tree).
//     * @warning This pointer may be null
//     */
//    ArbitraryKvBtNode *parent_tree = detachRoot(n);
//
//    if (removeSubtree) {
//        /**
//         * Remove node and its child trees...
//         *      If removeSubtree is set, we will simply delete the node
//         *      and cause a cascading operation facilitated by node
//         *      destructors.
//         */
//        DEBUG("removing subtree")
//        deleteSubtree(n);
//        DEBUG("subtree removed")
//        result = true;
//    } else {
//        DEBUG("removing only a single node.")
//        /**
//         * Remove only the specified node (n).
//         *      We are deleting only the single node and NOT
//         *      the left and right subtrees. This will result
//         *      in an unbalanced tree.
//         */
//        /**
//         * subtree_lhs - deleted node (n) left branch subtree.
//         *             - this points to the top of the subtree.
//         * @warning This pointer may be null
//         */
//        ArbitraryKvBtNode *subtree_lhs = n->left;
//        /**
//         * subtree_rhs - deleted node (n) right branch subtree.
//         *             - this points to the top of the subtree
//         * @warning This pointer may be null
//         */
//        ArbitraryKvBtNode *subtree_rhs = n->right;
//        /**
//         * Detach the left and right subtrees from the node (n)
//         * then delete the node (n).
//         */
//        n->left = n->right = NULL;
//        if(n) delete n; // ONLY the targeted node is deleted.
//        /**
//         * At this point we have either a parent tree and
//         * two orphaned subtrees or some combination of that
//         * state.
//         */
//        if (!parent_tree) {
//            /**
//             * In this branch, we have determined there is no parent tree.
//             * Now one of the subtrees must become the parent tree.
//             *
//             * @note To do this, we determine the new parent_tree by
//             *       looking at the balance of our current before the
//             *       deletion.  If the tree
//             */
//            if (_balanceCount < 0) {
//                /**
//                 * Our tree skews to the left, so we merge to the right.
//                 * Then we set subtree_rhs to NULL to avoid duplicates.
//                 */
//                parent_tree = subtree_rhs; //points to top of subtree
//                subtree_rhs = NULL;
//                /**
//                 * Move parent_tree pointer to the left-most endpoint.
//                 * We will join the subtree_lhs to this point.
//                 */
//                while (parent_tree && parent_tree->left) parent_tree = parent_tree->left;
//            } else {
//                /**
//                 * Our tree skews to the right, and we then
//                 * merge to the left. Our subtree_lhs is then set to NULL
//                 * to avoid duplicates.
//                 */
//                parent_tree = subtree_lhs; //points to top of subtree
//                subtree_lhs = NULL;
//                /**
//                 * Move parent_tree pointer to the right-most endpoint.
//                 * We will join the subtree_rhs to this point.
//                 */
//                while (parent_tree && parent_tree->left) parent_tree = parent_tree->right;
//            }
//        }
//        /**
//         * Join the subtree pointers to the parent tree.
//         */
//        parent_tree->left = subtree_lhs;
//        parent_tree->right = subtree_rhs;
//        /**
//         * Note: at this point we likely have a very unbalanced
//         * tree.  We should call reBalanceTree() to fix this.
//         */
//        result = true;
//    }
//    DEBUG("ArbitraryKvBtree::removeNode() terminating")
//    return true;
//}
