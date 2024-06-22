/**
 * @name ArbitraryKvBtree/src/balance.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name balance
 * @brief Return the tree's balance score.
 *
 * @return int
 */
inline int ArbitraryKvBtree::balance() {
    calcStats();
    _lockBalanceCount->wait()->on();
    int v = _balanceCount;
    _lockBalanceCount->off();
    return v;
}
