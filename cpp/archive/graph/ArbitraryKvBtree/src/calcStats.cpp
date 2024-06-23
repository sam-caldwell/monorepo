/**
 * @name ArbitraryKvBtree/src/CalcStats.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This file calculates statistics for the Binary Tree
 * class.  In the Binary Tree and higher-order graphs
 * statistics are useful in maintaining the health of
 * the tree (e.g. determining when to perform a balance
 * operation).
 */
#include "threadCountBalance.cpp"
#include "threadCountNodes.cpp"
#include "threadCountDuplicates.cpp"

/**
 * @name calcStats
 * @brief Calculate stats about the tree.
 *
 *    @note This method will spawn a set of threads which will concurrently
 *    execute to calculate different statistics about the tree, such
 *    as node count, tree balance and duplicate nodes.
 *
 *    @note Each statistic-specific thread will SimpleLock the statistic for the
 *    duration of the thread's execution for safety.
 */
void ArbitraryKvBtree::calcStats() {
    if (!_statsOk) {
        lock->wait()->on();
        thread t1(&ArbitraryKvBtree::threadCountNodes, this, root);
        thread t2(&ArbitraryKvBtree::threadCountBalance, this, root);
        thread t3(&ArbitraryKvBtree::threadCountDuplicates, this, root);
        t1.join();
        t2.join();
        t3.join();
        _statsOk = true;
        lock->off();
    }
}
