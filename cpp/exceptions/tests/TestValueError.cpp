/**
 * @name exceptions/tests/TestValueError.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/exceptions/exceptions.h"

using namespace std;


class TestValueError : public TestBase {
public:
    TestValueError(string n) { name = n; }

    ~TestValueError() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw ValueError("test");
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (ValueError e) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/