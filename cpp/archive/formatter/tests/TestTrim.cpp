/**
 * @name formatter/tests/TestTrim.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../formatter.h"

class TestTrim : public TestBase {
public:
    TestTrim(string n) { name = n; }

    ~TestTrim() {/*noop*/}

    bool main() {
        return expect(Strings::trim("  test  ").compare("test")==0, "trim ok 1") &&
               expect(Strings::trim("  test  me  ").compare("test  me")==0, "trim ok 2") &&
               expect(Strings::trim("\ttest\t").compare("test")==0, "trim ok 3");
    }
};/*end of class*/