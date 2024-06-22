/**
 * @name ArbitraryKvBtree/src/findNode.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef ArbitraryKvBtree_find_H
#define ArbitraryKvBtree_find_H

#include "seekRoot.cpp"

/**
 * @name findNode()
 * @brief Return a pointer to any node which exists and has the
 *        given criteria key.  Where no node exists, return NULL.
 *
 * @param criteria string
 * @param lockOperation bool
 * @return ArbitraryKvBtNode pointer
 */
ArbitraryKvBtNode *ArbitraryKvBtree::findNode(string criteria, bool lockOperation = false) {
    if (lockOperation) lock->wait()->on();
    ArbitraryKvBtNode *here = seekRoot(root);
    while (here) {
        if (criteria < here->key()) {
            here = here->left;
        } else if (criteria > here->key()) {
            here = here->right;
        } else {
            break; //Node found.  Exit loop.
        }
    }
    if (lockOperation) lock->off();
    return here;
}

#endif /**ArbitraryKvBtree_find_H*/
