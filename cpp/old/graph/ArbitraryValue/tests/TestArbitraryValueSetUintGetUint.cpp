/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetUintGetUint.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryValue.h"
#include "../../../../types/ValueTypes.h"

class TestArbitraryValueSetUintGetUint : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetUintGetUint(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetUintGetUint() {}

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
            t.getUint(true);
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
     * @param v uint
     * @param z uint
     * @return bool (result of test)
     */
    bool test_func(uint v, uint z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getUint() == 0, "Expect default (loose-typed) value 0. got:"+to_string(t.getUint())) &&
               expect(t.setUint(v), "Expect setter ok") &&
               expect(t.getType() == Uint, "Expect type uint") &&
               expect(t.getUint() == z, "Expect uint value returned (loose-typing)") &&
               expect(t.getUint(true) == z, "Expect uint value returned (strong-typing)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(0, 0), "Given a uint(0) expect uint 0") &&
               expect(test_func(1, 1), "Given a uint(1) expect uint 1") &&
               expect(test_func(1337, 1337), "Given a uint(1337) expect uint 1337") &&
               expect(test_strong_type(1337), "Expect no exception or type violation occurs");
    }
};