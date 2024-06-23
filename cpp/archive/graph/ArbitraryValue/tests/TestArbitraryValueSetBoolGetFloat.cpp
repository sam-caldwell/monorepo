/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetBoolGetFloat.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryValue.h"
#include "../../../../types/ValueTypes.h"

class TestArbitraryValueSetBoolGetFloat : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetBoolGetFloat(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetBoolGetFloat() {}

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
            t.getFloat(true);
            return false;
        } catch (runtime_error &e) {
            return true;
        }
    }

    /**
     * @name test_func
     * @brief Given an ArbitraryValue class instance (t) with a null (default state),
     *        we expect we can use the appropriate getter/setter methods to set and
     *        retrieve a given value of a given type.
     *
     * @param v bool
     * @param z float
     * @return bool (result of test)
     */
    bool test_func(bool v, float z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(!t.getBool(), "Expect default (loose-typed) value") &&
               expect(t.setBool(v), "Expect setter ok") &&
               expect(t.getType() == Bool, "Expect type Bool") &&
               expect(t.getFloat() == z, "Expect value returned (loose-typing)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(false, 0.0), "expect bool (false) to return float (0.0)") &&
               expect(test_func(true, 1.0), "expect bool (true) to return float (1.0)") &&
               expect(test_strong_type(true), "Expect strong type enforcement (true) ok.");
    }
};