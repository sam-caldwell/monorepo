/**
 * @name SimpleLock/tests/TestDoubleLock.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/SimpleLock/src/SimpleLock.h"

class TestDoubleLock : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestDoubleLock(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestDoubleLock() {}

    /**
     * @name test_double_locks()
     * @brief test will run two on() calls then a single off().
     *
     * @return bool
     */
    bool test_double_locks() {
        debug(name + "::test_double_locks() start");
        SimpleLock lock;
        lock.SetTimeout(1);
        return expect(!lock.check(), "Expect lock is not enabled: state " + to_string(lock.check())) &&
               expect(lock.on(), "Expect lock enables(0)") &&
               expect(lock.check(), "Expect lock is enabled(1): state " + to_string(lock.check())) &&
               expect(lock.on(), "Expect lock enables(1)") &&
               expect(lock.check(), "Expect lock is enabled(2): state " + to_string(lock.check())) &&
               expect(lock.off(), "Expect lock enables(2)") &&
               expect(!lock.check(), "Expect lock is not enabled: state " + to_string(lock.check()));
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_double_locks(), "test_double_locks() ok");
    }
};