/**
 * @name ConfigFileParser/tests/TestInitialState.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../../../../system/src/exists.h"
#include "../../../../system/environment.h"
#include "../src/ConfigFileParser.h"

using namespace File;

class TestInitialStateHappyPath : public TestBase {
private:
    Configuration *cfg;
    ConfigStateMap *map;
    const string cfg_file = "good_basic_config.yaml";
public:
    /**
     * @name class constructor
     * @brief setup the test.
     *
     * @param n string
     */
    TestInitialStateHappyPath(string n) {
        name = n;

        map = new ConfigStateMap;
        map->add("CONFIG_FILE", "config.file", String, true);
        map->add("secA.key", "secA.key", String, true);
        map->add("secB.key", "secB.key", Int, true);

        cfg = new Configuration;
        cfg->loadData(map, "CONFIG_FILE", cfg_file);
    }

    /**
     * @name class destructor
     * @brief destroy class
     */
    ~TestInitialStateHappyPath() {
        if (cfg) delete cfg;
        if (map) delete map;
    }

    /**
     * @name test_initial_state
     * @brief verify that the initial state of the test is okay.
     *
     * @return bool
     */
    bool test_initial_state() {
        debug("       filename: " + cfg_file);
        debug("config filename: " + cfg->getString("config.file"));
        return expect(cfg->count() == 1, "expect one record") &&
               expect(cfg->getString("config.file").compare(cfg_file) == 0,
                      "Verify the cfg object has the config filename.");
    }

    /**
     * @name test_instantiation
     * @brief expect that ConfigParser will instantiate without error if given proper inputs.
     * @return bool
     */
    bool test_instantiation() {
        try {
            ConfigFileParser parser(cfg, map);
            return expect(true, "Expected no errors in ConfigFileParser instantiation");
        } catch (exception &e) {
            error("error: " + string(e.what()));
            return expect(false, "Expected no errors in ConfigFileParser instantiation. Error:" + string(e.what()));
        }
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
        return expect(test_initial_state(), "Expect test_initial_state() ok") &&
               expect(test_instantiation(), "Expect test_instantiation() ok");
    }

}; /*end of class*/