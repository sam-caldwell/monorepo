/**
 * @name ArbitraryValue/tests/TestArbitraryValueTypeCheck.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryValue.h"
#include "../../../../types/ValueTypes.h"

class TestArbitraryValueTypeCheck : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueTypeCheck(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueTypeCheck() {}

    /**
     * @name TestNullType
     * @brief Given a newly created ArbitraryValue class instance (t),
     *        We expect  it's datatype will be ValueTypes Null.
     *
     * @return  bool (result of test)
     */
    bool testNullType() {
        ArbitraryValue t; // By default we expect a Null Value
        return t.getType() == Null;
    }

    /**
     * @name testTypeMismatchBool
     * @brief Given a newly created ArbitraryValue class
     *        instance (t) instantiated with a boolean 'true'
     *        value, we expect the class instance will be
     *        configured with ValueTypes Bool.
     *
     * @return  bool (result of test)
     */
    bool testTypeMismatchBool() {
        ArbitraryValue t(true);
        return t.getType() == Bool;
    }

    /**
     * @name testTypeMismatchDouble
     * @brief Given a newly created ArbitraryValue class instance (t)
     * instantiated with a double-precision float, we expect the class
     * instance will be configured with ValueTypes Double.
     *
     * @return  bool (result of test).
     */
    bool testTypeMismatchDouble() {
        ArbitraryValue t((double) (1.0));
        return t.getType() == Double;
    }

    /**
     * @name testTypeMismatchFloat
     * @brief Given a newly created ArbitraryValue class instance (t)
     * instantiated with a single-precision float, we expect the class
     * instance will be configured with ValueTypes Float.
     *
     * @return  bool (result of test).
     */
    bool testTypeMismatchFloat() {
        ArbitraryValue t((float) (1.0));
        return t.getType() == Float;
    }

    /**
     * @name testTypeMismatchInt
     * @brief Given a newly created ArbitraryValue class instance (t)
     * instantiated with an integer, we expect that the class instance
     * will be configured with ValueTypes Int.
     *
     * @return bool (result of test)
     */
    bool testTypeMismatchInt() {
        ArbitraryValue t((int) (-1));
        return t.getType() == Int;
    }

    /**
     * @name testTypeMismatchString
     * @brief Given a newly created ArbitraryValue class instance (t)
     * instantiated with a string value, we expect the class instance
     * will be configured with ValueTypes String.
     *
     * @return bool (result of test)
     */
    bool testTypeMismatchString() {
        ArbitraryValue t((string)("testString"));
        return t.getType() == String;
    }

    /**
     * @name testTypeMismatchUint
     * @brief Given a newly created ArbitraryValue class instance (t)
     * instantiated with an unsigned integer, we expect that the class
     * instance will be configured with ValueTypes Uint.
     *
     * @return bool (result of test)
     */
    bool testTypeMismatchUint() {
        ArbitraryValue t((uint) (1));
        return t.getType() == Uint;
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(testNullType(), "Expect default type of an ArbitraryValue to be Null") &&
               expect(testTypeMismatchBool(), "Given boolean, Expect ArbitraryValue to be Bool") &&
               expect(testTypeMismatchDouble(), "Given double, Expect ArbitraryValue to be Double") &&
               expect(testTypeMismatchFloat(), "Given float, Expect ArbitraryValue to be Float") &&
               expect(testTypeMismatchInt(), "Given int, Expect ArbitraryValue to be Int") &&
               expect(testTypeMismatchString(), "Given string, Expect ArbitraryValue to be String") &&
               expect(testTypeMismatchUint(), "Given uint, Expect ArbitraryValue to be Uint");
    }
};
