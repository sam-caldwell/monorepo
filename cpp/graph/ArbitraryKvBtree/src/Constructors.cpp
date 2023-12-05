/**
 * @name Configuration/src/Constructors.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name default class constructor
 * @brief initialize state (enforce unique)
 * @note Initialize SimpleLock mechanism
 * @note Ensure root is null
 * @note Ensure count is zero (0)
 * @note Unique is false.
 */
ArbitraryKvBtree::ArbitraryKvBtree() {
    lock = new SimpleLock(SIMPLE_LOCK_DEFAULT_TIMEOUT);
    _unique = true;
    _strongTyping = false;
    _statsOk = false;

    _lockNodeCount = new SimpleLock(SIMPLE_LOCK_DEFAULT_TIMEOUT);
    _lockBalanceCount = new SimpleLock(SIMPLE_LOCK_DEFAULT_TIMEOUT);
    _lockDuplicateCount = new SimpleLock(SIMPLE_LOCK_DEFAULT_TIMEOUT);

    _nodeCount = _balanceCount = _duplicateCount = 0;

    root = NULL;
}

/**
 * @name Class constructor
 * @brief Initialize state (unique configurable)
 * @note Ensure root is null
 * @note Ensure count is zero (0)
 * @note Set unique flag
 * @note Set balanceThreshold value
 *
 * @param unique bool.
 * @param strongTyping bool (default false)
 */
ArbitraryKvBtree::ArbitraryKvBtree(bool unique, bool strongTyping = false) {
    lock = new SimpleLock(SIMPLE_LOCK_DEFAULT_TIMEOUT);
    _unique = unique;
    _strongTyping = strongTyping;
    _statsOk = false;

    _lockNodeCount = new SimpleLock(SIMPLE_LOCK_DEFAULT_TIMEOUT);
    _lockBalanceCount = new SimpleLock(SIMPLE_LOCK_DEFAULT_TIMEOUT);
    _lockDuplicateCount = new SimpleLock(SIMPLE_LOCK_DEFAULT_TIMEOUT);

    _nodeCount = _balanceCount = _duplicateCount = 0;

    root = NULL;
}
