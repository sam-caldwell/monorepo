/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetBoolGetDouble.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "projects/graph/ArbitraryValue/src/ArbitraryValue.h"
#include "projects/common/types/ValueTypes.h"

class TestArbitraryValueSetBoolGetDouble : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetBoolGetDouble(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetBoolGetDouble() {}

    /**
     * @name test_strong_type
     * @brief Test strongTyping enforcement where we expect exception.
     *
     * @param v float
     * @return bool
     */
    bool test_strong_type(float v) {
        debug(name + "::test_strong_type(" + to_string(v) + ") start");
        ArbitraryValue t;
        t.setBool(v);
        try {
            t.getDouble(true);
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
     * @param v float
     * @param z double
     * @return bool (result of test)
     */
    bool test_func(float v, double z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(!t.getBool(), "Expect default (loose-typed) value") &&
               expect(t.setBool(v), "Expect setter ok") &&
               expect(t.getType() == Bool, "Expect type Bool") &&
               expect(t.getDouble() == z, "Expect value returned (loose-typing)");
    }

    /**
     * @name test_func
     * @brief Given an ArbitraryValue class instance (t) with a null (default state),
     *        we expect we can use the appropriate getter/setter methods to set and
     *        retrieve a given value of a given type.
     *
     * @return bool (result of test)
     */
    bool test_func(bool v, double z) {
        ArbitraryValue t;
        if (t.getBool()) return false; //expected false.  got true.
        t.setBool(v);
        if (t.getType() != Bool) return false; //expected String.
        if (t.getBool() != v) {
            cout << "setBool() failed to set boolean.  getBool() failed." << endl;
            return false;
        }
        bool result = (t.getDouble() == z);
        return result;
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return test_func(false, 0.0) &&
               test_func(true, 1.0) &&
               expect(test_strong_type(true), "Expect strong type enforcement (true) ok.");
    }
};