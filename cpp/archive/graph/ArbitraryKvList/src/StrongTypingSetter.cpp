/**
 * @name ArbitraryKvList/src/StrongTypingSetter.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name unique [setter]
 * @brief enable/disable strong typing enforcement
 *
 * @param flag bool guarantee strong typing
 * @return bool (_strongTyping flag)
 */
inline bool ArbitraryKvList::strongTyping(bool flag) {
    //ToDo: implement StackedLocks
    _strongTyping = flag;
    return _strongTyping;
}