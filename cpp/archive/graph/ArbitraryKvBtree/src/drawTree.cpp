/**
 * @name ArbitraryKvBtree/src/drawTree.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "drawTreeRecursively.cpp"

/**
 * @name drawTree()
 * @brief Draw Binary Tree.
 * @return string
 */
inline string ArbitraryKvBtree::drawTree() {
    return drawTreeRecursively(root, "root");
}
