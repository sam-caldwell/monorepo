/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetDoubleGetFloat.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryValue.h"
#include "../../../../types/ValueTypes.h"
#include <cmath>

#ifdef PI
#undef PI
#endif

#define PI 3.141529

class TestArbitraryValueSetDoubleGetFloat : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetDoubleGetFloat(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetDoubleGetFloat() {}

    /**
     * @name test_strong_type
     * @brief Test strongTyping enforcement where we expect exception.
     *
     * @param v double
     * @return bool
     */
    bool test_strong_type(double v) {
        debug(name + "::test_strong_type(" + to_string(v) + ") start");
        ArbitraryValue t;
        t.setDouble(v);
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
     * @param v double
     * @param z float
     * @return bool (result of test)
     */
    bool test_func(double v, float z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getDouble() == 0.0, "Expect default (loose-typed) value") &&
               expect(t.setDouble(v), "Expect setter ok") &&
               expect(t.getType() == Double, "Expect type Double") &&
               expect(t.getFloat() == z, "Expect float value returned (loose-typing)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(0.0, 0), "Expect float 0.0 to get/set proper value") &&
               expect(test_func(-1.0, -1), "Expect float -1.0 to get/set proper value") &&
               expect(test_func(1.0, 1), "Expect float 1.0 to get/set proper value") &&
               expect(test_func(-PI, -PI), "Expect float -3.141529 to get/set proper value") &&
               expect(test_func(PI, PI), "Expect float 3.141529 to get/set proper value") &&
               expect(test_strong_type(1.0), "Expect strong type enforcement (false) ok.");
    }
};