/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetUintGetDouble.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryValue.h"
#include "../../../../types/ValueTypes.h"

class TestArbitraryValueSetUintGetDouble : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetUintGetDouble(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetUintGetDouble() {}

    /**
     * @name test_strong_type
     * @brief Test strongTyping enforcement where we expect exception.
     *
     * @param v uint
     * @return bool
     */
    bool test_strong_type(uint v) {
        debug(name + "::test_strong_type(" + to_string(v) + ") start");
        ArbitraryValue t;
        t.setUint(v);
        try {
            t.getDouble(true);
            return false;
        } catch (runtime_error &e) {
            return ((string)(e.what())=="type mismatch");
        }
    }

    /**
     * @name test_func
     * @brief Given an ArbitraryValue class instance (t) with a null (default state),
     *        we expect we can use the appropriate getter/setter methods to set and
     *        retrieve a given value of a given type.
     *
     * @param v uint
     * @param z double
     * @return bool (result of test)
     */
    bool test_func(uint v, double z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getUint() == 0, "Expect default value 0") &&
               expect(t.setUint(v), "Expect setter ok") &&
               expect(t.getType() == Uint, "Expect type Uint") &&
               expect(t.getDouble(false) == z, "Expect string value '" + to_string(z) + "");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(0, 0.0), "Given uint(0), expect double 0") &&
               expect(test_func(1, 1.0), "Given uint(1), expect double 1") &&
               expect(test_func(UINT_MAX, (double) (UINT_MAX)), "Given uint(max) expect double UINT_MAX") &&
               expect(test_strong_type(0.0), "Expect exception on type violation");
    }
};