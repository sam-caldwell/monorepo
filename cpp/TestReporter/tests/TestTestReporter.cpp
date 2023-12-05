/**
 * @name TestReporter/tests/TestTestReporter.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/SimpleLock/src/SimpleLock.h"

class TestTestReporter : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestTestReporter(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestTestReporter() {}

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return true;
    }
};