/**
 * @name Configuration/tests/TestParseUint.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#define PUBLIC_FOR_TESTING_ONLY() public:
#define UNDO_PUBLIC_FOR_TESTING_ONLY() protected:

#include "../../../../types/src/ConfigStateMap/ConfigStateMap.h"
#include "../src/Configuration.h"

class TestParseUint : public TestBase {
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
    TestParseUint(string n) {
        name = n;
        map.add("myKey", "myRemoteKey", String);
    }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParseUint() {}

    /**
     * @name test_uint
     * @brief Given a string value ("<uint>", "")
     *
     * @param raw string
     * @param defaultValue uint
     * @return uint
     */
    uint test_uint(string raw, uint defaultValue) {
        return cfg.parseUint(raw, defaultValue);
    }

    /**
     * @name test_uint_error
     * @brief Given an out of range (<0) value, ensure we throw out_of_range exception
     *
     * @param raw string
     * @param defaultValue uint
     * @return bool
     */
    bool test_uint_error(string raw, uint defaultValue) {
        try {
            cfg.parseUint(raw, defaultValue);
            return false;
        } catch (out_of_range &e) {
            debug("expected exception caught: " + string(e.what()));
            return string(e.what()) == string("signed integer (i<0) is out of bounds for uint.");
        }
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        debug(name + "::main() start");
        return expect(test_uint("", 42) == 42, "expect test_uint('',42) ok") &&
               expect(test_uint("3.141529", 1) == 3, "expect test_uint('3.141529',1) ok") &&
               expect(test_uint("1337", 1) == 1337, "expect test_uint('1337',1) ok") &&
               expect(test_uint_error("-1", -5), "expect test_uint('-1', -5) ok (exception expected)");
    }
};
