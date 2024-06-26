/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetFloatGetFloat.cpp
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

class TestArbitraryValueSetFloatGetFloat : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetFloatGetFloat(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetFloatGetFloat() {}

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
            t.getFloat(true);
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
     * @param v float
     * @param z float
     * @return bool (result of test)
     */
    bool test_func(float v, float z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getFloat() == 0.0, "Expect default (loose-typed) value") &&
               expect(t.setFloat(v), "Expect setter ok") &&
               expect(t.getType() == Float, "Expect type Float") &&
               expect(t.getFloat() == z, "Expect value returned (loose-typing)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        const float pi = 3.141529;
        return expect(test_func(0,0), "Given float(0), expect return to match") &&
               expect(test_func(1,1), "Given float(1), expect return to match") &&
               expect(test_func(-1,-1), "Given float(-1), expect return to match") &&
               expect(test_func(PI,PI), "Given float(PI), expect return to match") &&
               expect(test_func(-PI,-PI), "Given float(-PI), expect return to match") &&
               expect(test_strong_type(1.0), "Expect strong type enforcement (false) ok.");
    }
};