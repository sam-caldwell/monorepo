/**
 * ArbitraryKvList/src/resetNodeRight.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef resetNodeRight_H
#define resetNodeRight_H

/**
 * @name resetNodeRight
 * @brief reset the current node to the right-most position and return the result.
 *
 * @param *p ArbitraryKvNode
 * @return ArbitraryKvNode pointer
 */
inline ArbitraryKvNode *resetNodeRight(ArbitraryKvNode *p) {
    if (p)
        do {
            p = (p->right) ? p->right : p;
        } while (p->right);
    return p;
}

#endif /* resetNodeRight_H */