/**
 * @name formatter/tests/TestRightTrimChar.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../formatter.h"

class TestRightTrimChar : public TestBase {
public:
    TestRightTrimChar(string n) { name = n; }

    ~TestRightTrimChar() {/*noop*/}

    bool main() {
        return expect(Strings::rtrimChar("  test  ", ' ') == "  test", "rtrimChar ok 1") &&
               expect(Strings::rtrimChar("  test  me  ", ' ') == "  test  me", "rtrimChar ok 2") &&
               expect(Strings::rtrimChar("test.", '.') == "test", "rtrimChar ok 3") &&
               expect(Strings::rtrimChar(".test", '.') == ".test", "rtrimChar ok 3") &&
               expect(Strings::rtrimChar("test..", '.') == "test", "rtrimChar ok 3") &&
               expect(Strings::rtrimChar("..test", '.') == "..test", "rtrimChar ok 3") &&
               expect(Strings::rtrimChar("....", '.') == "", "rtrimChar ok 3");
    }
};/*end of class*/