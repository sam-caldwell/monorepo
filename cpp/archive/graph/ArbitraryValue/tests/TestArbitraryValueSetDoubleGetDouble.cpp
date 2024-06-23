/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetDoubleGetDouble.cpp
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

class TestArbitraryValueSetDoubleGetDouble : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetDoubleGetDouble(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetDoubleGetDouble() {}

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
            t.getDouble(true);
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
     * @param v double
     * @return bool (result of test)
     */
    bool test_func(double v) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getDouble() == 0, "Expect default (loose-typed) value") &&
               expect(t.setDouble(v), "Expect setter ok") &&
               expect(t.getType() == Double, "Expect type Double") &&
               expect(t.getDouble() == v, "Expect float value returned (loose-typing)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(0.0), "Expect double 0.0 to get/set proper value") &&
               expect(test_func(-1.0), "Expect double -1.0 to get/set proper value") &&
               expect(test_func(1.0), "Expect double 1.0 to get/set proper value") &&
               expect(test_func(-PI), "Expect double -3.141529 to get/set proper value") &&
               expect(test_func(PI), "Expect double 3.141529 to get/set proper value") &&
               expect(test_strong_type(1.0), "Expect strong type enforcement (false) ok.");

    }
};