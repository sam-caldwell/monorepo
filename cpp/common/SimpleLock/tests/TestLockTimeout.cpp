/**
 * @name SimpleLock/tests/TestLockTimeout.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../src/SimpleLock.h"

class TestLockTimeout : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestLockTimeout(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestLockTimeout() {}

    /**
     * @name test_timeout_default()
     * @brief test timeout (default settings)
     *
     * @return bool
     */
    bool test_timeout_default() {
        debug(name + "::test_timeout_default() start");
        SimpleLock lock;
        try {
            debug(name + "::test_timeout_default() start test");
            lock.on()->wait()->off();
            debug(name + "::test_timeout_default() test complete");
            return true;
        } catch (LockTimeoutExpired e) {
            debug(name + "::test_timeout_default() timeout expired:" + ((string) e.what()));
            return false;
        }
    }

    /**
     * @name test_timeout()
     * @brief test timeout
     *
     * @return bool
     */
    bool test_timeout(uint n) {
        SimpleLock lock(n);
        debug(name + "::test_timeout() start");
        try {
            debug(name + "::test_timeout() set lock");
            lock.on();
            debug(name + "::test_timeout() wait on lock");
            lock.wait();
            debug(name + "::test_timeout() turn off lock");
            lock.off();
            debug(name + "::test_timeout() test complete");
            return true;
        } catch (LockTimeoutExpired e) {
            debug(name + "::test_timeout() timeout expired:" + ((string) e.what()));
            return false;
        }
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_timeout(15), "Test timeout(15) ok") &&
               expect(test_timeout(1500), "test timeout(1500) ok") &&
               expect(test_timeout(3000), "test timeout(3000) ok") &&
               expect(test_timeout_default(), "expect test timeout (defaults) OK");
    }
};