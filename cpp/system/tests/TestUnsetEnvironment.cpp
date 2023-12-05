/**
 * @name system/tests/TestUnsetEnvironment.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/system/environment.h"

class TestUnsetEnvironment : public TestBase {
public:
    TestUnsetEnvironment(string n) { name = n; }

    ~TestUnsetEnvironment() {}

    bool main() {
        debug(name + "::main()");
        return expect(Environment::set("TEST_ENV_VAR", "TEST_VALUE"),
                      "Expect we can set an Environment variable.") &&
               expect(Environment::get("TEST_ENV_VAR") == "TEST_VALUE",
                      "Expect we can retrieve the Environment variable.") &&
               expect(Environment::unset("TEST_ENV_VAR"), "expect we can unset the env var") &&
               expect([this]() -> bool {
                   Environment::unset("TEST_ENV_VAR");
                   try {
                       Environment::get("TEST_ENV_VAR");
                       return false; //Expect that we deleted the env var.
                   } catch (runtime_error e) {
                       return true;
                   }
               }(), "Expect the env var will no longer exist");
    }

};/*end of class*/