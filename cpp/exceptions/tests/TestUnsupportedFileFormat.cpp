/**
 * @name exceptions/tests/TestUnsupportedFileFormat.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#pragma GCC diagnostic ignored "-Wnonportable-include-path"
#include "projects/Tester/TestBase/main.h"
#include "projects/common/exceptions/exceptions.h"

class TestUnsupportedFileFormat : public TestBase {
public:
    TestUnsupportedFileFormat(string n) { name = n; }

    ~TestUnsupportedFileFormat() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw UnsupportedFileFormat("test");
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (UnsupportedFileFormat) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/
