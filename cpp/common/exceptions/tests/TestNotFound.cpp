/**
 * @name exceptions/tests/TestNotFound.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../exceptions.h"

using namespace std;


class TestNotFound : public TestBase {
public:
    TestNotFound(string n) { name = n; }
    ~TestNotFound() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw NotFound();
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (NotFound e) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/