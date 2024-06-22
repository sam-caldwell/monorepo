/**
 * @name ArbitraryKvList/src/InsertUint.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../../../../exceptions/exceptions.h"

/**
 * @name insert [uint]
 * @brief Insert new node to immediate left of root.
 *        Check uniqueness flag and guarantee unique
 *        key if enabled.
 *
 * @param key uint
 * @param v bool
 * @return operational outcome (success/fail)
 */
bool ArbitraryKvList::insert(string key, uint v) {
    lock->wait()->on();
    if ((_unique) && (exists(key))) return false;
    if (root) {
        ArbitraryKvNode *temp = NULL;
        ArbitraryKvNode *node = new ArbitraryKvNode(key, v);
        if (root->left) {
            node->right = root;
            node->left = root->left;
            root->left = node;
            node->left->right = node;
            //root->right is unchanged.
        } else {
            root->left = new ArbitraryKvNode(key, v);
            root->left->right = root;
        }
        if (_moveLeftOnInsert)
            root = root->left;
    } else {
        root = new ArbitraryKvNode(key, v);
    }
    _count++;
    lock->off();
    return true;
}