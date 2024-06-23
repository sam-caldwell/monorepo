/**
 * @name Configuration/tests/TestParseFloat.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#define PUBLIC_FOR_TESTING_ONLY() public:
#define UNDO_PUBLIC_FOR_TESTING_ONLY() protected:

#include "../../../../types/src/ConfigStateMap/ConfigStateMap.h"
#include "../src/Configuration.h"

class TestParseFloat : public TestBase {
private:
    Configuration cfg;
    ConfigStateMap map;
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParseFloat(string n) {
        name = n;
        map.add("myKey", "myRemoteKey", String);
    }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParseFloat() {}

    /**
     * @name test_float
     * @brief Given a string value (1.0,1.0, '')
     *
     * @param raw string
     * @param defaultValue float
     * @return float
     */
    float test_float(string raw, float defaultValue) {
        return cfg.parseFloat(raw, defaultValue);
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        debug(name + "::main() start");
        return expect(test_float("", 42.0) == 42.0, "expect test_float('',42.0) ok.") &&
               expect(test_float("3.14", 1.0) == (float)(3.140000), "expect test_float('3.14',1.0) ok.") &&
               expect(test_float("-1", 2.0) == -1.0, "expect test_float('-1', 2.0) ok");
    }
};
