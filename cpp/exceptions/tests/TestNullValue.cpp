/**
 * @name exceptions/tests/TestNullValue.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/exceptions/exceptions.h"

using namespace std;


class TestNullValue : public TestBase {
public:
    TestNullValue(string n) { name = n; }
    ~TestNullValue() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw NullValue();
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (NullValue e) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/