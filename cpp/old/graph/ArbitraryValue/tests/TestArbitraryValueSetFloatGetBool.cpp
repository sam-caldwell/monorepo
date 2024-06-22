/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetFloatGetBool.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryValue.h"
#include "../../../../types/ValueTypes.h"

#ifdef PI
#undef PI
#endif

#define PI 3.141529

class TestArbitraryValueSetFloatGetBool : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetFloatGetBool(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetFloatGetBool() {}

    /**
     * @name test_strong_type
     * @brief Test strongTyping enforcement where we expect exception.
     *
     * @param v float
     * @return bool
     */
    bool test_strong_type(float v) {
        debug(name + "::test_strong_type(" + to_string(v) + ") start");
        ArbitraryValue t;
        t.setFloat(v);
        try {
            t.getBool(true);
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
     * @param v float
     * @param z int
     * @return bool (result of test)
     */
    bool test_func(float v, bool z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(!t.getBool(), "Expect default (loose-typed) value") &&
               expect(t.setFloat(v), "Expect setter ok") &&
               expect(t.getType() == Float, "Expect type Float") &&
               expect(t.getBool() == z, "Expect value returned (loose-typing)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(0.0, false), "given float(0) expect false") &&
               expect(test_func(1.0, true), "given float(1) expect true") &&
               expect(test_func(3.141529, true), "given float(pi) expect true") &&
               expect(test_func(-3.141529, true), "given float(-pi) expect true") &&
               expect(test_strong_type(1.0), "Expect strong type enforcement (false) ok.");
    }
};