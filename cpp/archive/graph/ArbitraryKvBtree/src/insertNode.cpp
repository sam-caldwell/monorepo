/**
 * @name projects/src/ArbitraryKvBtree/src/insertNode.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef insertNode_H
#define insertNode_H

#include "../../../../exceptions/exceptions.h"
#include "findNode.cpp"
#include "seekRoot.cpp"



/**
 * @name insertNode
 * @brief Insert a given node (n) into the binary tree (root).
 *
 * @param *n ArbitraryKvBtNode pointer
 * @return bool (true: success, false: operation failed)
 */
bool ArbitraryKvBtree::insertNode(ArbitraryKvBtNode *n) {
    bool result = false;
    lock->wait()->on();
    /**
     * Perform our duplicate check first.  Note that we SimpleLock
     * the tree for this part to prevent changes between uniqueness
     * check and insertion.
     *
     * @note we use findNode() instead of exists() to turn off local locking
     *       since we already have a SimpleLock.
     */
    if (_unique && findNode(n->key(), false)) {
        /**
         * _unique flag is set (true) so uniqueness is enforced.
         * findNode() returns a non-null pointer, so a duplicate exists.
         * return false because the operation failed.
         */
        result = false;
    } else {
        /**
         * Check if root node is null.
         *  - If so, we are adding the initial node.
         *  - If not, then we will insert into the existing tree.
         */
        if (root) {
            ArbitraryKvBtNode *here = seekRoot(root);
            do {
                if (n->key() < here->key())
                    if (here->left) {
                        here = here->left;
                    } else {
                        result = (bool) (here->left = n);
                        here->left->root = here;
                        {/**start of statistics block*/
                            //Lock stats
                            _lockNodeCount->on();
                            _lockBalanceCount->on();

                            _nodeCount++;
                            _balanceCount--; //Decrease (-1) for left

                            //Unlock stats
                            _lockNodeCount->off();
                            _lockBalanceCount->off();
                        }/**end of statistics block*/
                        break;
                    }
                else if (n->key() > here->key())
                    if (here->right) {
                        here = here->right;
                    } else {
                        result = (bool) (here->right = n);
                        here->right->root = here;
                        {/**start of statistics block*/
                            //Lock stats
                            _lockNodeCount->on();
                            _lockBalanceCount->on();

                            _nodeCount++;
                            _balanceCount++; //Increase (+1) for right.

                            //Unlock stats
                            _lockNodeCount->off();
                            _lockBalanceCount->off();
                        }/**end of statistics block*/
                        break;
                    }
                else { //Duplicate node.
                    while (here && here->right) here = here->right;
                    if (_unique) {
                        result = false;
                    } else {
                        /**
                         * Performance impact:
                         *    - no impact for less-than values which
                         *      will branch left off the first node
                         *      in the duplicate chain.
                         *    - for nodes greater than the duplicate
                         *      node chain, there is a performance
                         *      impact which is measured in direct
                         *      proportion to the number of
                         *      duplicate nodes.
                         */
                        result = (bool) (here->right = n);
                        here->right->root = here;
                        {/**start of statistics block*/
                            //Lock stats
                            _lockNodeCount->on();
                            _lockBalanceCount->on();
                            _lockDuplicateCount->on();

                            _nodeCount++;
                            _balanceCount++; //increase balance because duplicates store to the right.
                            _duplicateCount++;

                            //Unlock stats
                            _lockNodeCount->off();
                            _lockBalanceCount->off();
                            _lockDuplicateCount->off();
                        }/**end of statistics block*/
                        break;
                    }
                }
            } while (true);
        } else {
            result = (bool) (root = n);
            {/**start of statistics block*/
                //Lock stats
                _lockNodeCount->on();
                _lockBalanceCount->on();

                _nodeCount++;
                _balanceCount = 0;   //We have 0 balance because this is the 1st node.
                _duplicateCount = 0; //We have 0 duplicates because this is the only node.

                //Unlock stats
                _lockNodeCount->off();
                _lockBalanceCount->off();
            }/**end of statistics block*/
        }
    }
    lock->off();
    return result;
}

#endif /*insertNode_H*/