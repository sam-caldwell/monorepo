/**
 * @name exceptions/tests/TestMissingInputs.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../../../networking/TcpCommon/src/PipeResponse/main.h"
#include "../exceptions.h"

class TestMissingInputs : public TestBase {
public:
    TestMissingInputs(string n) { name = n; }

    ~TestMissingInputs() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw MissingInputs("test");
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (MissingInputs) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/
