/**
 * @name exceptions/tests/TestUnexpectedType.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/exceptions/exceptions.h"

using namespace std;


class TestUnexpectedType : public TestBase {
public:
    TestUnexpectedType(string n) { name = n; }
    ~TestUnexpectedType() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw UnexpectedType();
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (UnexpectedType e) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/