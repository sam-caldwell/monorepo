/**
 * @name ArbitraryKvBtree/src/duplicates.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name duplicates
 * @brief Return the duplicate count.
 *
 * @return int
 */
inline int ArbitraryKvBtree::duplicates() {
    calcStats();
    _lockDuplicateCount->wait()->on();
    int v = _duplicateCount;
    _lockDuplicateCount->off();
    return _duplicateCount;
}
