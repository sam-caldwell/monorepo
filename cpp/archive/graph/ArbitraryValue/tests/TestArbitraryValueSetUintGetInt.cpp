/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetUintGetInt.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryValue.h"
#include "../../../../types/ValueTypes.h"

class TestArbitraryValueSetUintGetInt : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetUintGetInt(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetUintGetInt() {}

    /**
     * @name test_strong_type
     * @brief Test strongTyping enforcement where we expect exception.
     *
     * @param v uint
     * @param z int
     * @return bool
     */
    bool test_strong_type(uint v, int z) {
        debug(name + "::test_strong_type(" + to_string(v) + "," + to_string(z) + ") start");
        ArbitraryValue t;
        t.setUint(v);
        try {
            t.getInt(true);
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
     * @return bool (result of test)
     */
    bool test_func(uint v, int z) {
        debug(name + "::test_func(" + to_string(v) + "," + to_string(z) + ") start");
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getUint() == 0, "Expect default (loose-typed) value 0. got:" + to_string(t.getUint())) &&
               expect(t.setUint(v), "Expect setter ok") &&
               expect(t.getType() == Uint, "Expect type uint") &&
               expect(t.getInt(false) == z, "Expect uint value returned (loose-typing)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(0, 0.0), "Given uint(0), expect int 0") &&
               expect(test_func(1, 1.0), "Given uint(1), expect int 1") &&
               expect(test_func(UINT_MAX, (int) (UINT_MAX)), "Given uint(max) expect int UINT_MAX") &&
               expect(test_func(-UINT_MAX, (int) (-UINT_MAX)), "Given uint(-max) expect int -UINT_MAX") &&
               expect(test_strong_type(0,0), "Expect exception on type violation");
    }
};