/**
 * @name ArbitraryKvBtree/src/threadCountNodes.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "countRecursiveCount.cpp"

/**
 * @name threadCountNodes
 * @brief Wrapper for recursive node counter used by calcStats
 * to multi-thread the operation.
 * @param *n ArbitraryKvBtNode pointer
 */
void ArbitraryKvBtree::threadCountNodes(ArbitraryKvBtNode *n) {
    //DEBUG("ArbitraryKvBtree::threadCountNodes: running")
    this->_lockNodeCount->on();
    this->_nodeCount = countRecursiveCount(this->root);
    this->_lockNodeCount->off();
    //DEBUG("threadCountNodes: completed. _count:" + to_string(this->_nodeCount))
}

