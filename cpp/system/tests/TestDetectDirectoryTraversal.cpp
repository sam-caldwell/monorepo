/**
 * @name system/tests/TestDetectDirectoryTraversal.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include <stdlib.h>
#include "projects/common/system/file.h"

using namespace File;

class TestDetectDirectoryTraversal : public TestBase {
private:
    char const *e = "TEST_ENV_VAR";
    char const *v = "TEST_VALUE";
public:
    TestDetectDirectoryTraversal(string n) {
        name = n;
        setenv(e, v, 1);
    }

    ~TestDetectDirectoryTraversal() {
        unsetenv(e);
    }

    bool main() {
        debug(name + "::main()");
        return expect(!detectDirectoryTraversal("/tmp/simple.txt"), "expect good file/path ok1.") &&
               expect(!detectDirectoryTraversal("./tmp/simple.txt"), "expect good file/path ok2.") &&
               expect(detectDirectoryTraversal("../etc/passwd"), "expect detection1") &&
               expect(detectDirectoryTraversal("../../etc/passwd"), "expect detection2") &&
               expect(detectDirectoryTraversal("../etc/.."), "expect detection3") &&
               expect(detectDirectoryTraversal("file:///%2e%2e/%2e%2e/etc/passwd"), "expect detection4");
    }

};/*end of class*/