/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetIntGetInt.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryValue.h"
#include "../../../../types/ValueTypes.h"

class TestArbitraryValueSetIntGetInt : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetIntGetInt(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetIntGetInt() {}

    /**
     * @name test_strong_type
     * @brief Test strongTyping enforcement where we expect exception.
     *
     * @param v int
     * @return bool
     */
    bool test_strong_type(int v) {
        debug(name + "::test_strong_type(" + to_string(v) + ") start");
        ArbitraryValue t;
        t.setInt(v);
        try {
            t.getInt(true);
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
     * @param v int
     * @return bool (result of test)
     */
    bool test_func(int v) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getString() == "", "Expect default (loose-typed) value ''.") &&
               expect(t.setInt(v), "Expect setter ok") &&
               expect(t.getType() == Int, "Expect type Int") &&
               expect(t.getInt() == v, "Expect float value returned (loose-typing)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(0),"Given int(0), expect int(0)") &&
               expect(test_func(1),"Given int(1), expect int(1)") &&
               expect(test_func(-1),"Given int(-1), expect int(-1)") &&
               expect(test_func(INT_MIN),"Given int(INT_MIN), expect int(INT_MIN)") &&
               expect(test_func(-INT_MIN),"Given int(-INT_MIN), expect int(-INT_MIN)") &&
               expect(test_strong_type(0), "Expect strong type enforcement (0) ok.");
    }
};