/**
 * @name ArbitraryKvBtree/src/threadCountBalance.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "balanceRecursiveCount.cpp"

/**
 * @name threadCountBalance
 * @brief Wrapper for recursive tree-balance counter used by calcStats
 * to multi-thread the operation.
 * @param *n ArbitraryKvBtNode pointer
 */
void ArbitraryKvBtree::threadCountBalance(ArbitraryKvBtNode *n) {
    this->_lockBalanceCount->on();
    this->_balanceCount = balanceRecursiveCount(this->root, 0);
    this->_lockBalanceCount->off();
}
