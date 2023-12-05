/**
 * @name ArbitraryKvBtree/src/Destructors.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "seekRoot.cpp"

/**
 * @name class destructor
 * @brief Delete all data objects.
 */
ArbitraryKvBtree::~ArbitraryKvBtree() {
    lock->wait()->on();
    _statsOk = true; //prevent any re-calc.
    if (root) {
        root = seekRoot(root);
        /*
         * Delete the tree node-by-node.
         */
        while (root && (root->left || root->right)) {
            if (root->left) root = root->left;
            else if (root->right) root = root->right;
            else {
                ArbitraryKvBtNode *target = root;
                root = root->root; //Move root back one to continue.
                target->root = NULL;
                delete target;
            }
        }
    }
    lock->off();
}
