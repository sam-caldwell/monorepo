/**
 * @name formatter/tests/TestLeftTrim.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../formatter.h"

class TestLeftTrim : public TestBase {
public:
    TestLeftTrim(string n) { name = n; }

    ~TestLeftTrim() {/*noop*/}

    bool main() {
        return expect(Strings::ltrim("  test  ") == "test  ", "ltrim ok 1") &&
               expect(Strings::ltrim("  test  me  ") == "test  me  ", "ltrim ok 2") &&
               expect(Strings::ltrim("\ttest\t") == "test\t", "ltrim ok 3");
    }
};/*end of class*/