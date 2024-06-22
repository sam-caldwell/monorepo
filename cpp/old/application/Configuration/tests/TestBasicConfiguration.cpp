/**
 * @name Configuration/tests/TestBasicConfiguration.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../src/Configuration.h"

class TestBasicConfiguration : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestBasicConfiguration(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestBasicConfiguration() {}

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        Configuration cfg;
        return true;
    }
};
