/**
 * @name flags/tests/TestBoolean.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/common/flags/flags.h"

class TestBoolean : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestBoolean(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestBoolean() {/*noop*/}

    /**
     * @name test_bool
     * @brief test the boolean flags
     *
     * @params a - int
     * @params b - int
     * @return bool
     */
    bool test_bool(int a, int b) {
        return a==b;
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_bool(TRUE, 1), "Expect test_arch TRUE(1) ok") &&
               expect(test_bool(FALSE, 0), "Expect test_arch FALSE(0) ok");
    }

};