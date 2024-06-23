/**
 * @name exceptions/tests/TestValueGetStringError.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../exceptions.h"

using namespace std;


class TestValueGetStringError : public TestBase {
public:
    TestValueGetStringError(string n) { name = n; }

    ~TestValueGetStringError() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw ValueGetStringError();
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (ValueGetStringError e) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/