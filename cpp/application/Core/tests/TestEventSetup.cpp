/**
 * @name Application/tests/TestEventSetup.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/system/file.h"
#include "projects/Tester/Mocks/MockStdout.h"
#include "projects/logging/Logger/src/Logger.h"
#include "projects/common/system/environment.h"
#include "projects/application/Core/src/configure.cpp"

class TestEventSetup : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestEventSetup(string n) {
        name = n;
    }
    ~TestEventSetup(){}
};