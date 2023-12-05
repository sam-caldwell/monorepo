/**
 * @name common/system/tests/TestExec.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "projects/Tester/TestBase/main.h"

#include <iostream>
#include "projects/common/system/exec.h"

class TestExec : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestExec(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestExec() {}

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(exec::run("echo 'test'").compare("test\n") == 0, "expect simple echo ok");
    }

};