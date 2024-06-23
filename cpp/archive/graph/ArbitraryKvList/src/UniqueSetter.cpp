/**
 * @name ArbitraryKvList/src/UniqueSetter.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name unique [setter]
 * @brief set the current unique-ness
 *
 * @param unique bool guarantee unique keys
 * @return bool (unique flag)
 */
inline bool ArbitraryKvList::unique(bool unique) {
    //ToDo: implement StackedLocks
    _unique = unique;
    return unique;
}