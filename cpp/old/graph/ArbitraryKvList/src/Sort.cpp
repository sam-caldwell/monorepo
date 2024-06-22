/**
 * @name ArbitraryKvList/src/Sort.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "SelectSortFunction.cpp"

/**
 * @name sort
 * @brief Sort the list ascending or descending.
 *
 * @param order SortOrder
 * @return bool
 */
bool ArbitraryKvList::sort(SortOrder order) {
    bool result = false;
    FuncEvaluate evaluate = SelectSortFunction(order);
    if (!evaluate) {
        return false;
    }
    if (root) {
        lock->wait()->on();
        for (int p = 0; p < count(); p++) {
            ArbitraryKvNode *LHS = resetNodeLeft(root);
            for (int q = 0; q < (count() - p); q++) {
                if ((LHS->right) && (evaluate(LHS, LHS->right)))
                    swapNodes(LHS, LHS->right);
            }
        }
        result = true;
    }
    lock->off();
    return result;
}