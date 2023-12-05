/**
 * @name exceptions/tests/TestFileIoError.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#pragma GCC diagnostic ignored "-Wnonportable-include-path"
#include "projects/Tester/TestBase/main.h"
#include "projects/common/exceptions/exceptions.h"

class TestFileIoError : public TestBase {
public:
    TestFileIoError(string n) { name = n; }

    ~TestFileIoError() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw FileIoError("test");
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (FileIoError) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/
