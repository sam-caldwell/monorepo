/**
 * @name Configuration/tests/TestParseDouble.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#define PUBLIC_FOR_TESTING_ONLY() public:
#define UNDO_PUBLIC_FOR_TESTING_ONLY() protected:

#include "projects/common/types/ConfigStateMap.h"
#include "projects/application/Configuration/src/Configuration.h"

class TestParseDouble : public TestBase {
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
    TestParseDouble(string n) {
        name = n;
        map.add("myKey", "myRemoteKey", String);
    }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParseDouble() {}

    /**
     * @name test_double
     * @brief Given a string value (1.0,1.0, '')
     *
     * @param raw string
     * @param defaultValue double
     * @return double
     */
    double test_double(string raw, double defaultValue) {
        return cfg.parseDouble(raw, defaultValue);
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        debug(name + "::main() start");
        return expect(test_double("", 42.0) == 42.0, "expect test_double('',42.0) ok.") &&
               expect(test_double("3.14", 1.0) == 3.14, "expect test_double('3.14',1.0) ok") &&
               expect(test_double("-1", 2.0) == -1.0, "expect test_double('-1', 2.0) ok");
    }
};
