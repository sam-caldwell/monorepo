/**
 * @name ArbitraryKvList/src/CountMatches.cpp
 * @copyright (c) 2022 Sam Caldwell. All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name count [matches]
 * @brief Count the number of nodes with the given key.
 *
 * @param key string
 * @return int (count of matches)
 */
int ArbitraryKvList::count(string criteria) {
    //ToDo: Implement stacked SimpleLock
    int match_count = 0;
    if (root) {
        ArbitraryKvNode *curr;
        curr = root;
        do {
            curr = (curr->left) ? curr->left : curr;
        } while (curr->left);
        for (int i = 0; i < count(); i++) {
            if (curr->key() == criteria) match_count++;
            curr = curr->right;
        }
    }
    return match_count;
}
