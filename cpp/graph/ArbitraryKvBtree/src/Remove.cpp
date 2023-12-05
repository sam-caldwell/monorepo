/**
 * @name ArbitraryKvBtree/src/Remove.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 * @brief This file contains the node removal (delete) functionality.
 */
#include "deleteNode.cpp"
#include "deleteSubtree.cpp"

/**
 * @name remove
 * @brief Remove all nodes with a given key (criteria)
 *
 * @param criteria string
 * @param removeSubtree bool
 * @return bool
 */
bool ArbitraryKvBtree::remove(string criteria, bool removeSubtree = false) {
    /**
     * Lock the tree.
     */
    lock->wait()->on();
    /**
     * Find our node (no node means delete not possible).
     */
    bool result = true;
    ArbitraryKvBtNode *targetNode = findNode(criteria, false);
    /**
     * If targetNode is NULL, there is no node to delete.
     * Otherwise we delete this targetNode.
     */
    if (targetNode) {
        /**
         * Before deleting anything, we need to ensure the root node
         * pointer is at the top of the tree.  Otherwise we may find
         * it points to a node that is subject to deletion.
         */
        root = seekRoot(root);
        /**
         * if the root node is our target node, we need to point to NULL
         * to allow the root node to be deleted.
         */
        root = (root == targetNode) ? NULL : root;
        /**
         * determine if we will delete the subtree or a specific node.
         */
        result = (removeSubtree) ? deleteSubtree(targetNode) : deleteNode(targetNode);
    }/*end of if(targetNode)*/
    /**
     * Invalidate the stats cache.  This will cause
     * the stats to be regenerated on the next call
     * to a stats method.
     */
    _statsOk = false;
    /**
     * Free the tree SimpleLock.
     */
    lock->off();
    /**
     * return boolean result.
     */
    return result;
}/*ArbitraryKvBtree::remove()*/

