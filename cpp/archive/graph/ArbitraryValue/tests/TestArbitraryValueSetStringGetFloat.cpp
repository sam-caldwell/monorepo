/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetStringGetFloat.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryValue.h"
#include "../../../../types/ValueTypes.h"
#include <stdlib.h>

class TestArbitraryValueSetStringGetFloat : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetStringGetFloat(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetStringGetFloat() {}

    /**
     * @name test_strong_type
     * @brief Test strongTyping enforcement where we expect exception.
     *
     * @param v string
     * @return bool
     */
    bool test_strong_type(string v) {
        debug(name + "::test_strong_type(" + v + ") start");
        ArbitraryValue t;
        t.setString(v);
        try {
            t.getFloat(true);
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
     * @param v string
     * @param z float
     * @return bool (result of test)
     */
    bool test_func(string v, float z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getString() == "", "Expect default (loose-typed) value ''.") &&
               expect(t.setString(v), "Expect setter ok") &&
               expect(t.getType() == String, "Expect type String") &&
               expect(t.getFloat() == z, "Expect float value returned (loose-typing)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func("0", 0), "Given string '0' expect float(0)") &&
               expect(test_func("1", 1), "Given string '1' expect float(1)") &&
               expect(test_func("3.141529", 3.141529), "Given string pi expect float pi") &&
               expect(test_func("-3.141529", -3.141529), "Given string -pi expect float -pi") &&
               expect(test_func("test", 0), "Given string 'test', expect float(0)") &&
               expect(test_func("", 0), "Given string '', expect float(0)") &&
               expect(test_strong_type("0"), "Expect strong type enforcement (0) ok.");
    }
};