/**
 * @name ArbitraryKvList/src/Constructors.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name Default Class constructor.
 * @brief Initialize class (SimpleLock mechanism, root pointer, count, and set unique false).
 */
ArbitraryKvList::ArbitraryKvList() {
    lock = new SimpleLock(SIMPLE_LOCK_DEFAULT_TIMEOUT);
    _unique = false;
    _moveLeftOnInsert = false;
    _count = 0;
    root = NULL;
}

/**
 * @name Class Constructor
 * @brief Initialize class (SimpleLock mechanism, root pointer,
 *        count, set unique and moveLeftOnInsert flag).
 *
 * @param unique bool.
 * @param moveLeftOnInsert bool flag to move root pointer left on insert.
 * @param strongTyping bool flag to enforce strong typing. (default false)
 */
ArbitraryKvList::ArbitraryKvList(bool unique, bool moveLeftOnInsert = false, bool strongTyping=false) {
    lock = new SimpleLock(SIMPLE_LOCK_DEFAULT_TIMEOUT);
    _moveLeftOnInsert = moveLeftOnInsert;
    _strongTyping = strongTyping;
    _unique = unique;
    _count = 0;
    root = NULL;
}
