/**
 * @name networking/TcpCommon/tests/TestWaitFor.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "projects/Tester/TestBase/main.h"

#include <iostream>
#include "projects/networking/TcpCommon/src/main.h"

class TestWaitFor : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestWaitFor(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestWaitFor() {}

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        networking::FileDescriptor fd;
        fd.set(1);
        return expect(networking::FdWait::waitFor(fd, 1) == networking::FdWait::Result::TIMEOUT,
                      "expect waitFor expire ok (TIMEOUT) ...1s");
    }

};