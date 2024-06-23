/**
 * @name ArbitraryKvList/src/SelectSortFunction.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef SelectSortFunction_H
#define SelectSortFunction_H

#include "../../../../types/SortOrder.h"

/**
 * @name SelectSortFunction
 * @brief Select the sort order function based on the input
 *        which will point the caller to either GreaterThan()
 *        or LessThan().
 *
 * @param SortOrder order
 * @return FuncEvaluate function pointer.
 */
FuncEvaluate SelectSortFunction(SortOrder order) {
    switch (order) {
        case Ascending: {
            return GreaterThan;
            break;
        }
        case Descending: {
            return LessThan;
            break;
        }
        default: {
            break;
        }
    }
    return NULL;
}

#endif /* SelectSortFunction_H */
