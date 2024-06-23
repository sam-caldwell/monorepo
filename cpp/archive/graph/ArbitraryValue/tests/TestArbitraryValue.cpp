/**
 * @name ArbitraryValue/tests/TestArbitraryValue.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryValue.h"

#ifdef PI
#undef PI
#endif /*PI*/
#define PI 3.141529

class TestArbitraryValue : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValue(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValue() {}

    /**
     * @name test_default_state
     * @brief verify we can instantiate an ArbitraryValue object
     *        and that its initial state meets expectations.
     *
     * @return bool
     */
    bool test_default_state() {
        ArbitraryValue v;
        return expect(v.getType() == Null, "Expect default type is Null.  Got: " + ValueTypeToString(v.getType()));
    }

    /**
     * @name test_bool
     * @brief initialize ArbitraryValue with boolean
     *
     * @param n bool
     * @return bool
     */
    bool test_bool(bool n) {
        debug(name + "::test_bool("+ to_string(n)+") start");
        ArbitraryValue v(n);
        debug(name + "::test_bool("+ to_string(n)+") object initialized");
        return expect(v.getType() == Bool, "Expect default type is Bool");
    }

    /**
     * @name test_double
     * @brief initialize ArbitraryValue with double
     *
     * @return bool
     */
    bool test_double() {
        ArbitraryValue v((double) (PI));
        return expect(v.getType() == Double, "Expect default type is Double");
    }

    /**
     * @name test_float
     * @brief initialize ArbitraryValue with float
     *
     * @return bool
     */
    bool test_float() {
        ArbitraryValue v((float) (PI));
        return expect(v.getType() == Float, "Expect default type is Float");
    }

    /**
     * @name test_int
     * @brief initialize ArbitraryValue with int
     *
     * @return bool
     */
    bool test_int() {
        ArbitraryValue v((int) (1));
        return expect(v.getType() == Int, "Expect default type is Int");
    }

    /**
     * @name test_string
     * @brief initialize an arbitrary value object with a string.
     *
     * @return bool
     */
    bool test_string() {
        ArbitraryValue v((string)"test_arbitrary_value_string");
        return expect(v.getType() == String, "Expect default type is String");
    }

    /**
     * @name test_uint
     * @brief initialize an arbitrary value object with a uint (unsigned int).
     *
     * @return bool
     */
    bool test_uint() {
        ArbitraryValue v((uint) (1));
        return expect(v.getType() == Uint, "Expect default type is Uint");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_default_state(), "Expect ArbitraryValue will instantiate without error") &&
               expect(test_bool(true), "Expect ArbitraryValue will instantiated with bool(true)") &&
               expect(test_bool(false), "Expect ArbitraryValue will instantiated with bool(false)") &&
               expect(test_double(), "Expect ArbitraryValue will instantiated with double") &&
               expect(test_float(), "Expect ArbitraryValue will instantiated with float") &&
               expect(test_int(), "Expect ArbitraryValue will instantiated with int") &&
               expect(test_string(), "Expect ArbitraryValue will instantiated with String") &&
               expect(test_uint(), "Expect ArbitraryValue will instantiated with uint");
    }
};
