/**
 * @name exceptions/tests/TestBoundsCheckError.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/exceptions/exceptions.h"

using namespace std;


class TestBoundsCheckError : public TestBase {
public:
    TestBoundsCheckError(string n) { name = n; }
    ~TestBoundsCheckError() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw BoundsCheckError();
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (BoundsCheckError e) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/