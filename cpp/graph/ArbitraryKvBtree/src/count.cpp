/**
 * @name ArbitraryKvBtree/src/count.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name count
 * @brief Count the nodes.
 *
 * @return int
 */
inline int ArbitraryKvBtree::count() {
    calcStats();
    _lockNodeCount->wait()->on();
    int v = _nodeCount;
    _lockNodeCount->off();
    return _nodeCount;
}