/**
 * @name Configuration/tests/TestParseBool.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#define PUBLIC_FOR_TESTING_ONLY() public:
#define UNDO_PUBLIC_FOR_TESTING_ONLY() protected:

#include "../../../../types/src/ConfigStateMap/ConfigStateMap.h"
#include "../src/Configuration.h"

class TestParseBool : public TestBase {
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
    TestParseBool(string n) {
        name = n;
        map.add("myKey", "myRemoteKey", String);
    }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParseBool() {}

    /**
     * @name test_bool
     * @brief Given a string value (true,false, '')
     *
     * @param raw string
     * @param defaultValue bool
     * @return bool
     */
    bool test_bool(string raw, bool defaultValue) {
        return cfg.parseBool(raw, defaultValue);
    }

    /**
     * @name test_error
     * @brief Given an invalid input, expect an exception
     *
     *
     * @param raw string
     * @param defaultValue bool
     * @return bool
     */
    bool test_error(string raw, bool defaultValue) {
        try {
            cfg.parseBool(raw, defaultValue);
            return false;
        } catch (runtime_error &e) {
            debug("expected error: " + string(e.what()));
            return true;
        }
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_bool("true", true), "expect test_bool('true',true) ok") &&
               expect(test_bool("true", false), "expect test_bool('true',false) ok.") &&
               expect(test_bool("false", true) == false, "expect test_bool('false',true) ok") &&
               expect(test_bool("false", false) == false, "expect test_bool('false',false) ok") &&
               expect(test_bool("", true), "expect test_bool('',true) ok") &&
               expect(test_bool("", false) == false, "expect test_bool('',false) ok") &&
               expect(!test_error("true", true), "expect test_error() ok (baseline)") &&
               expect(test_error("-1", true), "expect test_error() ok (-1)") &&
               expect(test_error("0.0", true), "expect test_error() ok (0.0)") &&
               expect(test_error("---", true), "expect test_error() ok (---)") &&
               expect(test_error("Invalid Boolean", true), "expect test_error() ok ('Invalid Boolean')");
    }
};
