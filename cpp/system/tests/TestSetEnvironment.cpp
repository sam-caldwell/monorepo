/**
 * @name system/tests/TestSetEnvironment.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include <stdlib.h>
#include "projects/common/system/environment.h"

class TestSetEnvironment : public TestBase {
public:
    TestSetEnvironment(string n) { name = n; }

    ~TestSetEnvironment() {}

    bool main() {
        debug(name + "::main()");
        return expect(Environment::set("TEST_ENV_VAR", "TEST_VALUE"),
                      "Expect we can set an Environment variable.") &&
               expect(Environment::get("TEST_ENV_VAR") == "TEST_VALUE",
                      "Expect we can retrieve the Environment variable.");

    }

};/*end of class*/