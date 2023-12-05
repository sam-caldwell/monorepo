/**
 * @name ArbitraryKvList/src/Destructor.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name Class destructor.
 * @brief clean up at shutdown.
 */
ArbitraryKvList::~ArbitraryKvList() {
    lock->wait();
    delete root;
    _count--;
}