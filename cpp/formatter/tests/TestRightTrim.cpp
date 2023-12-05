/**
 * @name formatter/tests/TestRightTrim.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/formatter/formatter.h"

class TestRightTrim : public TestBase {
public:
    TestRightTrim(string n) { name = n; }

    ~TestRightTrim() {/*noop*/}

    bool main() {
        return expect(Strings::rtrim("  test  ") == "  test", "rtrim ok 1") &&
               expect(Strings::rtrim("  test  me  ") == "  test  me", "rtrim ok 2") &&
               expect(Strings::rtrim("\ttest\t") == "\ttest", "rtrim ok 3");
    }
};/*end of class*/