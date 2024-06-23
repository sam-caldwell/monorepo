/**
 * @name SimpleLock/src/SimpleLock.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 * @todo: Document SimpleLock
 */

#ifndef SIMPLE_LOCK_H
#define SIMPLE_LOCK_H

#ifdef _WIN32
/*
 * Windows
 */
#include <Windows.h>  //Provides windows time

#else
/*
 * Non-windows
 */
#include <unistd.h>   //Provides posix time.

#endif /* ifdef _WIN32 */

#ifndef SIMPLE_LOCK_DEFAULT_TIMEOUT
/**
 * default lock timeout.
 *     The timeout starts at lock->on and
 *     expires in the specified number of
 *     milliseconds.
 */
#define SIMPLE_LOCK_DEFAULT_TIMEOUT 1800
#endif /* SIMPLE_LOCK_DEFAULT_TIMEOUT */

#ifndef SIMPLE_LOCK_DEFAULT_INTERVAL
/**
 * default timeout interval.
 *     This is the interval between timeout
 *     checks.
 */
#define SIMPLE_LOCK_DEFAULT_INTERVAL 1000
#endif /* SIMPLE_LOCK_DEFAULT_INTERVAL */

#ifdef _WIN32
#include <Windows.h>  //Provides windows time
#else

#include <unistd.h>   //Provides posix time.

#endif

#include <iostream>
#include <chrono>
#include <thread>
#include "../../exceptions/exceptions.h"

class SimpleLock {
public:
    /**
     * @name Default Class Constructor.
     * @brief default state SimpleLock is disabled.
     */
    SimpleLock();

    /**
     * @name Class Constructor.
     * @brief default state SimpleLock is disabled.
     *
     * @param uint timeout_ms - timeout for wait()
     * @param uint interval_ms - timeout check interval (default: SIMPLE_LOCK_DEFAULT_INTERVAL)
     */
    SimpleLock(uint timeout_ms, uint interval_ms);

    /**
     * @name Class Destructor.
     * @brief noop destructor
     */
    ~SimpleLock();

    /**
     * @name SetTimeout
     * @brief sets the timeout (in seconds)
     *
     * @param int timeout_ms - timeout for wait()
     * @param int interval_ms - timeout check interval (default: SIMPLE_LOCK_DEFAULT_INTERVAL)
     * @return SimpleLock*
     */
    SimpleLock *SetTimeout(uint timeout_ms, uint interval_ms);

    /**
     * @name on
     * @brief enable the SimpleLock.
     *
     * @return SimpleLock*
     */
    inline SimpleLock *on();

    /**
     * @name off
     * @brief disable the SimpleLock.
     *
     * @return SimpleLock* pointer
     */
    inline SimpleLock *off();

    /**
     * @name wait
     * @brief wait for SimpleLock to release for t seconds.
     *
     * @return SimpleLock* pointer
     */
    SimpleLock *wait();

    /**
     * @name check
     * @brief return boolean (true: locked, false: unlocked) status
     *
     * @return bool
     */
    inline bool check();

private:
    /**
     * lock - the flag that indicates the current lock state.
     */
    bool lock;
    /**
     * timeout - The amount of time to be awaited until the lock expires.
     */
    int timeout;
    /**
     * interval - The amount of time to wait between timeout checks.
     */
    int interval;
};

#include "check.cpp"
#include "constructors.cpp"
#include "destructors.cpp"
#include "off.cpp"
#include "on.cpp"
#include "SetTimeout.cpp"
#include "wait.cpp"

#endif //SIMPLE_LOCK_H