/**
 * @name EnvironmentVariableParser/tests/TestVerifyData.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include <stdlib.h>
#include "../../../../system/environment.h"
#include "../src/EnvironmentVariableParser.h"

class TestVerifyData : public TestBase {
private:
    ConfigStateMap *map;
    Configuration *cfg;
    EnvironmentVariableParser *parser;
public:
    /**
     * @name class constructor
     * @brief initialize the test by setting two environment variables.
     *
     * @param n string
     */
    TestVerifyData(string n) {
        debug(n + "::TestParse()");
        name = n;
        Environment::set("TEST_ENV_A", "test_value");
        Environment::set("TEST_ENV_B", "1");
        map = new ConfigStateMap;
        cfg = new Configuration;
        parser = new EnvironmentVariableParser(cfg, map);

        map->add("TEST_ENV_A", "test.env.a", String, true, false);
        map->add("TEST_ENV_B", "test.env.b", Int, true, false);
    }

    /**
     * @name class destructor
     * @brief clean up the environment variables created for the test.
     */
    ~TestVerifyData() {
        debug(name + "::~TestParse()");
        Environment::unset("TEST_ENV_A");
        Environment::unset("TEST_ENV_B");
        if (parser) delete parser;
        if (map) delete map;
        if (cfg) delete cfg;
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
        debug(name + "::main()");
        /**
         * Create our environment variables.
         */
        return expect(parser->parse(), "Expect parser ok") &&
               expect(cfg->count() == 2, "Expect parser has 2 elements (actual:" + to_string(cfg->count()) + ")");
    }
};