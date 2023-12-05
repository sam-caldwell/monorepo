/**
 * ArbitraryKvList/src/resetNodeLeft.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef resetNodeLeft_H
#define resetNodeLeft_H

/**
 * @name resetNodeLeft
 * @brief reset the current node to the left-most position and return the result.
 *
 * @param *p ArbitraryKvNode
 * @return ArbitraryKvNode pointer
 */
inline ArbitraryKvNode *resetNodeLeft(ArbitraryKvNode *p) {
    if (p)
        do {
            p = (p->left) ? p->left : p;
        } while (p->left);
    return p;
}

#endif /* resetNodeLeft_H */