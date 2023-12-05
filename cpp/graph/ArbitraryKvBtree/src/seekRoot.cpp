/**
 * @name projects/src/ArbitraryKvBtree/src/seekRoot.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef seekRoot_H
#define seekRoot_H

/**
 * @name seekRoot
 * @brief This method will seek the root of a given tree.
 * @param *n ArbitraryKvBtNode *
 * @return *ArbitraryKvBtNode pointer
 */
ArbitraryKvBtNode *seekRoot(ArbitraryKvBtNode *n) {
    ArbitraryKvBtNode *p = n;
    while (p && p->root) p = p->root;
    return p;
}

#endif /*seekRoot_H*/