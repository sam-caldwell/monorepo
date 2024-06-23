/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetDoubleGetString.cpp
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

class TestArbitraryValueSetDoubleGetString : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetDoubleGetString(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetDoubleGetString() {}

    /**
     * @name test_strong_type
     * @brief Test strongTyping enforcement where we expect exception.
     *
     * @param v float
     * @return bool
     */
    bool test_strong_type(double v) {
        debug(name + "::test_strong_type(" + to_string(v) + ") start");
        ArbitraryValue t;
        t.setDouble(v);
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
     * @param v double
     * @param z string
     * @return bool (result of test)
     */
    bool test_func(double v, string z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getDouble() == 0.0, "Expect default (loose-typed) value") &&
               expect(t.setDouble(v), "Expect setter ok") &&
               expect(t.getType() == Double, "Expect type Double") &&
               expect(t.getString() == z, "Expect float value returned (loose-typing)") &&
               expect(test_strong_type(0), "Expect strong type enforcement (0) ok.");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(0, "0.000000"), "Expect double(0) to return string '0.000000'") &&
               expect(test_func(1, "1.000000"), "Expect double(1) to return string '1.000000'") &&
               expect(test_func(-1, "-1.000000"), "Expect double(-1) to return string '-1.000000'") &&
               expect(test_func(PI, "3.141529"), "Expect double(pi) to return string '3.141529'") &&
               expect(test_func(-PI, "-3.141529"), "Expect double(-pi) to return string '-3.141529'") &&
               expect(test_func(-PI, to_string(-PI)), "Expect float (-pi) to return string '-3.141529'");
    }
};