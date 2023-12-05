/**
 * @name system/tests/TestGetEnvironment.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include <stdlib.h>
#include "projects/common/system/environment.h"

class TestGetEnvironment : public TestBase {
private:
    char const *e = "TEST_ENV_VAR";
    char const *v = "TEST_VALUE";
public:
    TestGetEnvironment(string n) {
        name = n;
        setenv(e, v, 1);
    }

    ~TestGetEnvironment() {
        unsetenv(e);
    }

    bool main() {
        debug(name + "::main()");
        return expect(Environment::get("TEST_ENV_VAR") == "TEST_VALUE",
                      "Expect we can read an Environment variable.");
    }

};/*end of class*/