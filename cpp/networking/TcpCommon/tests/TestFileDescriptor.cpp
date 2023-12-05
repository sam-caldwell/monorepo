/**
 * @name networking/TcpCommon/tests/TestFileDescriptor.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "projects/Tester/TestBase/main.h"

#include <iostream>
#include <limits.h>
#include "projects/networking/TcpCommon/src/main.h"


class TestFileDescriptor : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestFileDescriptor(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestFileDescriptor() {}

    /**
     * @name test
     * @brief Verify we can set/get integer values from -max to +max as file descriptors.
     *
     * @return bool
     */
    bool test() {
        networking::FileDescriptor fd;
        for (int i = INT_MIN; i < INT_MAX; i++) {
            fd.set(i);
            if (fd.get() != i)
                return expect(false, "test i="+to_string(i)+" ok");
            else
                return expect(true, "test i="+to_string(i)+" ok");
        }
        return expect(true, "FileDescriptor range ok");
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test(), "Expect test() ok");
    }

};