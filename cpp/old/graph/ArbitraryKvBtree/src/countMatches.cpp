/**
 * @name ArbitraryKvBtree/src/countMatches.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "seekRoot.cpp"

/**
 * @name count(criteria)
 * @brief Count the number of nodes with the given key.
 *
 * @param key string
 * @return int (count of matches)
 */
int ArbitraryKvBtree::count(string criteria) {
    bool local_lock;
    if (!lock->check()) {
        lock->on();
        local_lock = true;
    }
    root = seekRoot(root);
    ArbitraryKvBtNode *here = findNode(criteria, false);
    int matchCount = 0;
    if (here == NULL) {
        return false;
    } else {
        matchCount++;
        while(here->right){ //Count duplicate keys.
            here=here->right;
            if(here->key() == criteria) matchCount++;
        }
    }
    if (local_lock) lock->off();
    return matchCount;
}
