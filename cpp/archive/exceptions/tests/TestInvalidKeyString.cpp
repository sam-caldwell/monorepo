/**
 * @name exceptions/tests/TestInvalidKeyString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../exceptions.h"

using namespace std;


class TestInvalidKeyString : public TestBase {
public:
    TestInvalidKeyString(string n) { name = n; }
    ~TestInvalidKeyString() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw InvalidKeyString();
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (InvalidKeyString e) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/
