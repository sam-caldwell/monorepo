/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetIntGetString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryValue.h"
#include "../../../../types/ValueTypes.h"

class TestArbitraryValueSetIntGetString : public TestBase {
public:
    /**
     * Class constructor
     * @param string n : test name
     */
    TestArbitraryValueSetIntGetString(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetIntGetString() {}

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
            t.getString(true);
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
     * @param v int
     * @param z string
     * @return bool (result of test)
     */
    bool test_func(int v, string z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getString() == "", "Expect default (loose-typed) value ''.") &&
               expect(t.setInt(v), "Expect setter ok") &&
               expect(t.getType() == Int, "Expect type Int") &&
               expect(t.getString() == z, "Expect float value returned (loose-typing)") &&
               expect(test_strong_type(0), "Expect strong type enforcement (0) ok.");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(0, "0"), "Given int(0), expect '0'") &&
               expect(test_func(1, "1"), "Given int(1), expect '1'") &&
               expect(test_func(-1, "-1"), "Given int(-1), expect '-1'") &&
               expect(test_func(INT_MAX, to_string(INT_MAX)),
                      "Given int(INT_MAX), expect '" + to_string(INT_MAX) + "'") &&
               expect(test_func(INT_MIN, to_string(INT_MIN)),
                      "Given int(INT_MIN), expect 'INT_MIN'") &&
               expect(test_strong_type(1.0), "Expect strong type enforcement (false) ok.");
        return true;
    }
};