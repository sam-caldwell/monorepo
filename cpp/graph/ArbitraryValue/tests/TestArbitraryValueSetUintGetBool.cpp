/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetUintGetBool.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "projects/graph/ArbitraryValue/src/ArbitraryValue.h"
#include "projects/common/types/ValueTypes.h"

class TestArbitraryValueSetUintGetBool : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetUintGetBool(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetUintGetBool() {}

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
            t.getBool(true);
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
     * @param z bool
     * @return bool (result of test)
     */
    bool test_func(uint v, bool z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getUint() == 0, "Expect default value 0") &&
               expect(t.setUint(v), "Expect setter ok") &&
               expect(t.getType() == Uint, "Expect type Uint") &&
               expect(t.getBool(false) == z, "Expect bool value '" + to_string(z) + "'");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(0, false), "Given uint(0), expect bool false") &&
               expect(test_func(1, true), "Given uint(1), expect bool true") &&
               expect(test_func(UINT_MAX, true), "Given uint(max) expect bool true") &&
               expect(test_strong_type(true), "Expect exception on type violation(true)") &&
               expect(test_strong_type(false), "Expect exception on type violation(false)");
    }
};