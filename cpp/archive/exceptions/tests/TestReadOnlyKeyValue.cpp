/**
 * @name exceptions/tests/TestReadOnlyKeyValue.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../exceptions.h"

using namespace std;


class TestReadOnlyKeyValue : public TestBase {
public:
    TestReadOnlyKeyValue(string n) { name = n; }
    ~TestReadOnlyKeyValue() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        try {
            throw ReadOnlyKeyValue();
            debug(name + "::main() exception expected but not thrown");
            return false;
        } catch (ReadOnlyKeyValue e) {
            debug(name + "::main() exception expected and thrown");
        }
        return true;
    }
};/*end of class*/