/**
 * @name formatter/tests/TestToUpper.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/formatter/formatter.h"

class TestToUpper : public TestBase {
public:
    TestToUpper(string n) {
        name = n;
    }

    ~TestToUpper() {/*noop*/}

    bool main() {
        debug(name + "::main()");
        const string actual="qwertyuiopasdfghjklzxcvbnm1234567890()-_=+[{]}\\|;:',<.>/?`~";
        const string expected="QWERTYUIOPASDFGHJKLZXCVBNM1234567890()-_=+[{]}\\|;:',<.>/?`~";
        return expect(toUpper(actual)==expected, "Expect uppercase switch ok");
    }

};/*end of class*/