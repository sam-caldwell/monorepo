/**
 * @name ArbitraryKvList/src/Exists.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name exists
 * @brief Determine if a node with the given key exists.
 *        To return the contents of the node, use 'search()'
 *
 * @param criteria string
 * @return operation outcome (success/fail)
 */
bool ArbitraryKvList::exists(string criteria) {
    //ToDo: Implement stacked SimpleLock
    if (root) {
        ArbitraryKvNode *curr;
        curr = root;
        do {
            curr = (curr->left) ? curr->left : curr;
        } while (curr->left);
        for (int i = 0; i < count(); i++) {
            if (curr->key() == criteria) return true;
            curr = curr->right;
        }
    }
    return false;
}
