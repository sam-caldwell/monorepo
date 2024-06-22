/**
 * @name formatter/tests/TestToLower.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../formatter.h"

class TestToLower : public TestBase {
public:
    TestToLower(string n) {
        name = n;
    }

    ~TestToLower() {/*noop*/}


    bool main() {
        debug(name + "::main()");
        const string actual="QWERTYUIOPASDFGHJKLZXCVBNM1234567890()-_=+[{]}\\|;:',<.>/?`~";
        const string expected="qwertyuiopasdfghjklzxcvbnm1234567890()-_=+[{]}\\|;:',<.>/?`~";

        return expect(toLower(actual)==expected, "Expect lowercase switch ok");
    }

};/*end of class*/