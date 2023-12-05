/**
 * @name SimpleLock/tests/TestOnOff.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/SimpleLock/src/SimpleLock.h"

class TestOnOff : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestOnOff(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestOnOff() {}

    /**
     * @name test_on_check_off()
     * @brief test on(), check(), off()
     *
     * @return bool
     */
    bool test_on_check_off() {
        debug(name + "::test_on_check_off() start");
        SimpleLock lock;
        if (lock.check()) return false; //expect no lock
        lock.on();
        if (!lock.check()) return false; //expect lock
        lock.off();
        if (lock.check()) return false; //expect no lock
        return true;
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_on_check_off(), "test_on_check_off() ok");
    }
};