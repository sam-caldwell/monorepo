/**
 * @name SimpleLock/src/wait.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "SimpleLock.h"

using namespace std;

/**
 * @name wait
 * @brief wait for SimpleLock to release for t seconds.
 *
 * @return SimpleLock* pointer
 */
SimpleLock *SimpleLock::wait() {
    for (int pass = (timeout / interval); lock && (pass > 0); --pass) {
        std::this_thread::sleep_for(std::chrono::duration<double, std::milli>(interval));
    }
    lock = false;
    //ToDo: the locking entity should be able to register a callback on timeout we should call here.
    return this;
}
