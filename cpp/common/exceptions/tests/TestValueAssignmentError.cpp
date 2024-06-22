/**
 * @name exceptions/tests/TestValueAssignmentError.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../exceptions.h"

using namespace std;


class TestValueAssignmentError : public TestBase {
public:
    TestValueAssignmentError(string n) { name = n; }
    ~TestValueAssignmentError() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw ValueAssignmentError();
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (ValueAssignmentError e) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/