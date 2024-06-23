/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetBoolGetInt.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryValue.h"
#include "../../../../types/ValueTypes.h"

class TestArbitraryValueSetBoolGetInt : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetBoolGetInt(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetBoolGetInt() {}

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
            t.getInt(true);
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
     * @param z int
     * @return bool (result of test)
     */
    bool test_func(bool v, int z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(!t.getBool(), "Expect default (loose-typed) value") &&
               expect(t.setBool(v), "Expect setter ok") &&
               expect(t.getType() == Bool, "Expect type Bool") &&
               expect(t.getInt() == z, "Expect value returned (loose-typing)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect (test_func(false,0),"Expect bool false to return int 0") &&
               expect (test_func(true, 1), "Expect bool true to return int 1") &&
               expect(test_strong_type(true), "Expect strong type enforcement (true) ok.");
    }
};