/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetBoolGetUint.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "projects/graph/ArbitraryValue/src/ArbitraryValue.h"
#include "projects/common/types/ValueTypes.h"

class TestArbitraryValueSetBoolGetUint : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetBoolGetUint(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetBoolGetUint() {}

    /**
     * @name test_strong_type
     * @brief Test strongTyping enforcement where we expect exception.
     *
     * @param v bool
     * @return bool
     */
    bool test_strong_type(bool v) {
        debug(name + "::test_strong_type(" + to_string(v) + ") start");
        ArbitraryValue t;
        t.setBool(v);
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
     * @param v bool
     * @param z uint
     * @return bool (result of test)
     */
    bool test_func(bool v, uint z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getBool() == false, "Expect default (loose-typed) value") &&
               expect(t.setBool(v), "Expect setter ok") &&
               expect(t.getType() == Bool, "Expect type Bool") &&
               expect(t.getUint() == z, "Expect float value returned (loose-typing)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(false,0),"Expect bool false to return uint 0") &&
               expect(test_func(true,1),"Expect bool true to return uint 1") &&
               expect(test_strong_type(true), "Expect strong type enforcement (false) ok.");
    }
};