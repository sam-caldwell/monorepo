/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetIntGetBool.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "projects/graph/ArbitraryValue/src/ArbitraryValue.h"
#include "projects/common/types/ValueTypes.h"

class TestArbitraryValueSetIntGetBool : public TestBase {
public:
    /**
     * Class constructor
     * @param string n : test name
     */
    TestArbitraryValueSetIntGetBool(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetIntGetBool() {}

    /**
     * @name test_strong_type
     * @brief Test strongTyping enforcement where we expect exception.
     *
     * @param v int
     * @return bool
     */
    bool test_strong_type(bool v) {
        debug(name + "::test_strong_type(" + to_string(v) + ") start");
        ArbitraryValue t;
        t.setInt(v);
        try {
            t.getBool(true);
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
     * @param z bool
     * @return bool (result of test)
     */
    bool test_func(int v, bool z) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getBool() == false, "Expect default (loose-typed) value") &&
               expect(t.setInt(v), "Expect setter ok") &&
               expect(t.getType() == Int, "Expect type Int") &&
               expect(t.getBool() == z, "Expect float value returned (loose-typing)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_func(0, false), "Given int(0), expect boolean false") &&
               expect(test_func(1, true), "Given int(1), expect boolean true") &&
               expect(test_func(-1, true), "Given int(-1) expect boolean true") &&
               expect(test_func(INT_MAX, true), "Given int(min), expect boolean true") &&
               expect(test_func(INT_MIN, true), "Given int(max), expect boolean true") &&
               expect(test_strong_type(true), "Expect strong type enforcement (true) ok.") &&
               expect(test_strong_type(false), "Expect strong type enforcement (false) ok.");
    }
};