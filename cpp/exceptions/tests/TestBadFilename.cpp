/**
 * @name exceptions/tests/TestBadFilename.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/Tester/TestBase/main.h"
#include "projects/common/exceptions/exceptions.h"

class TestBadFilename : public TestBase {
public:
    TestBadFilename(string n) { name = n; }

    ~TestBadFilename() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw BadFilename("test");
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (BadFilename) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/
`