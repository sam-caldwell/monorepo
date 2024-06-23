/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetBoolGetBool.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include <string>
#include "../src/ArbitraryValue.h"

class TestArbitraryValueSetBoolGetBool : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetBoolGetBool(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetBoolGetBool() {}

    /**
     * @name test_strong_type
     * @brief Test strongTyping enforcement where we expect exception.
     *
     * @param v bool
     * @return bool
     */
    bool test_strong_type(bool v) {
        debug(name + "::test_strong_type(" + to_string(v) + ") start");
        ArbitraryValue t;
        t.setBool(v);
        try {
            t.getBool(true);
            return true;
        } catch (runtime_error &e) {
            return false;
        }
    }

    /**
     * @name test_func
     * @brief Given an ArbitraryValue class instance (t) with a null (default state),
     *        we expect we can use the appropriate getter/setter methods to set and
     *        retrieve a given value of a given type.
     *
     * @param v bool
     * @param z bool
     * @return bool (result of test)
     */
    bool test_func(bool v, bool z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(!t.getBool(), "Expect default (loose-typed) value") &&
               expect(t.setBool(v), "Expect setter ok") &&
               expect(t.getType() == Bool, "Expect type Bool") &&
               expect(t.getBool() == z, "Expect value returned (loose-typing)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(true, true), "Given a true value, expect the same to be properly set.") &&
               expect(test_func(false, false), "Given a false value, expect the same to be properly set.") &&
               expect(test_strong_type(true), "Expect strong type enforcement (true) ok.");
    }
};

