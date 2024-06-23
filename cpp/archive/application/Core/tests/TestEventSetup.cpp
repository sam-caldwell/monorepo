/**
 * @name Application/tests/TestEventSetup.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../../../../system/file.h"
#include "../../../../Tester/Mocks/MockStdout.h"
#include "../../../../logging/Logger/src/Logger.h"
#include "../../../../system/environment.h"
#include "../src/configure.cpp"

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