/**
 * @name exceptions/tests/TestTypeError.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../exceptions.h"

using namespace std;

class TestTypeError : public TestBase {
public:
    TestTypeError(string n) { name = n; }

    ~TestTypeError() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw TypeError("test");
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (TypeError e) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/