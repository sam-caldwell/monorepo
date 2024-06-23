/**
 * @name ArbitraryKvList/src/Search.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name search
 * @brief Search for a given key and leave root pointing
 *        to the given node.  If we do not find the criteria,
 *        then root will stay where it was.
 *
 * @param criteria string
 * @return bool (true: found, false: not found)
 */
bool ArbitraryKvList::search(string criteria) {
    //ToDo: Implement stacked SimpleLock
    if (root) {
        ArbitraryKvNode *curr;
        curr = root;
        do { //reset curr left.
            curr = (curr->left) ? curr->left : curr;
        } while (curr->left);
        for (int i = 0; i < count(); i++) {
            if (curr->key() == criteria) return (root = curr);
            curr = curr->right;
        }
    }
    return false; //Note: we leave root where it was when we started.
}