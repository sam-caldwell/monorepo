/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetDoubleGetUint.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "projects/graph/ArbitraryValue/src/ArbitraryValue.h"
#include "projects/common/types/ValueTypes.h"

class TestArbitraryValueSetDoubleGetUint : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetDoubleGetUint(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetDoubleGetUint() {}

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
     * @param v double
     * @param z uint
     * @return bool (result of test)
     */
    bool test_func(double v, uint z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getDouble() == 0.0, "Expect default (loose-typed) value") &&
               expect(t.setDouble(v), "Expect setter ok") &&
               expect(t.getType() == Double, "Expect type Double") &&
               expect(t.getUint() == z, "Expect float value returned (loose-typing)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(0.0, 0), "Expect double (0) input to result in matching uint") &&
               expect(test_func(1, 1), "Expect double(1) input to result in matching uint") &&
               expect(test_func((double)(UINT_MAX), UINT_MAX), "Expect (double)(UINT_MAX), UINT_MAX)") &&
               expect(test_strong_type(1.0), "Expect strong type enforcement (false) ok.");
    }
};