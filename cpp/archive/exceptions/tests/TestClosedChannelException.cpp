/**
 * @name exceptions/tests/TestClosedChannelException.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../exceptions.h"

using namespace std;


class TestClosedChannelException : public TestBase {
public:
    TestClosedChannelException(string n) { name = n; }

    ~TestClosedChannelException() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw closedChannelException();
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (closedChannelException e) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/