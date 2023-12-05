/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetFloatGetUint.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "projects/graph/ArbitraryValue/src/ArbitraryValue.h"
#include "projects/common/types/ValueTypes.h"

#ifdef PI
#undef PI
#endif

#define PI 3.141529

class TestArbitraryValueSetFloatGetUint : public TestBase {
public:
    /**
    * @name Class constructor
    * @brief initialize test class
    *
    * @param string n : test name
    */
    TestArbitraryValueSetFloatGetUint(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetFloatGetUint() {}

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
            t.getUint(true);
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
     * @param z uint
     * @return bool (result of test)
     */
    bool test_func(float v, uint z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getFloat() == 0.0, "Expect default (loose-typed) value") &&
               expect(t.setFloat(v), "Expect setter ok") &&
               expect(t.getType() == Float, "Expect type Float") &&
               expect(t.getUint() == z, "Expect value returned (loose-typing)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(0.0, 0), "Given float(0) expect matching uint.") &&
               expect(test_func(1.0, 1), "Given float(1) expect matching uint.") &&
               expect(test_strong_type(1.0), "Expect strong type enforcement (false) ok.");
    }
};