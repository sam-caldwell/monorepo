/**
 * @name EnvironmentVariableParser/tests/TestEvParserDataLoad.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include <stdlib.h>
#include "projects/common/system/environment.h"
#include "projects/application/EnvironmentVariableParser/src/EnvironmentVariableParser.h"

class TestEvParserDataLoad : public TestBase {
private:
    Configuration *cfg;
    ConfigStateMap *map;
public:
    /**
     * @name class constructor
     * @brief initialize the test by setting two environment variables.
     *        then setup a Configuration and ConfigStateMap object.
     *
     * @param n string
     */
    TestEvParserDataLoad(string n) {
        name = n;
        Environment::set("TEST_ENV_A", "test_value");
        Environment::set("TEST_ENV_B", "1");

        map = new ConfigStateMap;
        map->add("TEST_ENV_A", "test.env.a", String, true, false);
        map->add("TEST_ENV_B", "test.env.b", Int, true, false);

        cfg = new Configuration;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestEvParserDataLoad() {
        Environment::unset("TEST_ENV_A");
        Environment::unset("TEST_ENV_B");
        if (cfg) delete cfg;
        if (map) delete map;
    }

    /**
     * @name main
     * @brief create a map for our test environment variables, then parse them
     *        into the EnvironmentVariableParser object.  Test only that we can
     *        then pop two objects into the current cache.
     *
     * @warning This does not validate that the data is properly stored/retrieved.
     *          This only (intentionally) tests that we can perform the operations
     *          without any obvious errors.
     *
     * @return bool
     */
    bool main() {
        /**
         * Create our environment variables.
         */
        bool result = true;


        EnvironmentVariableParser parser(cfg, map);
        return expect(parser.parse(), "Expect parser ok") &&
               expect(cfg->count() == 2, "Expect parser has 2 elements (actual:" + to_string(cfg->count()) + ")");
    }
};