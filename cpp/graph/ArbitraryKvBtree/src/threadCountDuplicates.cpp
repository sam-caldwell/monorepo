/**
 * @name ArbitraryKvBtree/src/threadCountDuplicates.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "duplicateRecursiveCount.cpp"

/**
 * @name threadCountDuplicates
 * @brief Wrapper for recursive duplicate-node counter
 *        used by calcStats to multi-thread the operation.
 *
 * @param *n ArbitraryKvBtNode pointer
 */
void ArbitraryKvBtree::threadCountDuplicates(ArbitraryKvBtNode *n) {
    //DEBUG("ArbitraryKvBtree::threadCountDuplicates: running")
    this->_lockDuplicateCount->on();
    this->_duplicateCount = duplicateRecursiveCount(this->root);
    this->_lockDuplicateCount->off();
    //DEBUG("threadCountDuplicates: completed. _duplicates:" + to_string(this->_duplicateCount))
}

