/**
 * @name exceptions/tests/TestLockTimeoutExpired.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/exceptions/exceptions.h"

using namespace std;


class TestLockTimeoutExpired : public TestBase {
public:
    TestLockTimeoutExpired(string n) { name = n; }
    ~TestLockTimeoutExpired() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw LockTimeoutExpired();
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (LockTimeoutExpired e) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/