/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetFloatGetString.cpp
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

class TestArbitraryValueSetFloatGetString : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetFloatGetString(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetFloatGetString() {}

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
            t.getString(true);
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
     * @param z string
     * @return bool (result of test)
     */
    bool test_func(float v, string z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getFloat() == 0.0, "Expect default (loose-typed) value") &&
               expect(t.setFloat(v), "Expect setter ok") &&
               expect(t.getType() == Float, "Expect type Float") &&
               expect(t.getString() == z, "Expect value returned (loose-typing)") &&
               expect(test_strong_type(0), "Expect strong type enforcement (0) ok.");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        const float pi = 3.141529;
        return expect(test_func(0, "0.000000"), "Expect float(0) to return string '0.000000'") &&
               expect(test_func(1, "1.000000"), "Expect float(1) to return string '1.000000'") &&
               expect(test_func(-1, "-1.000000"), "Expect float(-1) to return string '-1.000000'") &&
               expect(test_func(PI, to_string(PI)), "Expect float (pi) to return string '3.141529'") &&
               expect(test_func(-PI, to_string(-PI)), "Expect float (-pi) to return string '-3.141529'") &&
               expect(test_strong_type(1.0), "Expect strong type enforcement (false) ok.");
    }
};