/**
 * @name Configuration/tests/TestParseInt.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#define PUBLIC_FOR_TESTING_ONLY() public:
#define UNDO_PUBLIC_FOR_TESTING_ONLY() protected:

#include "projects/common/types/ConfigStateMap.h"
#include "projects/application/Configuration/src/Configuration.h"

class TestParseInt : public TestBase {
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
    TestParseInt(string n) {
        name = n;
        map.add("myKey", "myRemoteKey", String);
    }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParseInt() {}

    /**
     * @name test_int
     * @brief Given a string value ("<int>", "")
     *
     * @param raw string
     * @param defaultValue int
     * @return int
     */
    int test_int(string raw, float defaultValue) {
        return cfg.parseInt(raw, defaultValue);
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        debug(name + "::main() start");
        return expect(test_int("", 42) == 42, "expect test_int('',42) ok") &&
               expect(test_int("3.141529", 1) == 3, "expect test_int('3.141529',1) ok") &&
               expect(test_int("1337", 1) == 1337, "expect test_int('1337',1) ok") &&
               expect(test_int("-1", -5) == -1, "expect test_int('-1', -5) ok");
    }
};
