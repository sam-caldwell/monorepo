/**
 * @name formatter/tests/TestIntToString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/formatter/formatter.h"

class TestIntToString : public TestBase {
public:
    TestIntToString(string n) {
        name = n;
    }

    ~TestIntToString() {/*noop*/}

    bool main() {
        debug(name + "::main()");
        return expect(intToString(123, 4) == "0123",
                      "Expect 3-digit int with width four ok:" + intToString(123, 4)) &&
               expect(intToString(-123, 4) == "-123",
                      "Expect neg 3-digit int with width four ok:" + intToString(-123, 4)) &&
               expect(intToString(+12, 4) == "0012",
                      "Expect 2-digit int with width four ok:" + intToString(+12, 4));
    }

};/*end of class*/