/**
 * @name Application/tests/TestConfigure.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/system/file.h"
#include "projects/common/system/environment.h"
#include "projects/application/Core/src/configure.cpp"

class TestConfigure : public TestBase {
private:
    ConfigStateMap *map;
    const string cfg_file = "test_config.yaml";
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestConfigure(string n) {
        name = n;

        map = new ConfigStateMap;
        /*
         * Environment Variables
         */
        map->add("CONFIG_FILE", ConfigFile, String, false, true);
        map->add("LOG_LEVEL", ConfigLogLevel, String, false, true);
        map->add("LOG_OUTPUT", ConfigLogOutput, String, false, true);
        /*
         * Command-line Arguments
         */
        map->add("--config", ConfigFile, String, false, true);
        map->add("--logLevel", ConfigLogLevel, String, false, true);
        map->add("--logOutput", ConfigLogOutput, String, false, true);
        /*
         * Config File
         */
        map->add(ConfigLogOutput, ConfigLogOutput, Bool, false, true);
        map->add(ConfigLogLevel, ConfigLogLevel, String, false, true);
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestConfigure() {
        Environment::unset("CONFIG_FILE");
        Environment::unset("LOG_LEVEL");
        Environment::unset("LOG_OUTPUT");
        //if (map) delete map;
    }

    /**
     * @name verify_file
     * @brief verify that the given file exists
     *
     * @param filename - string
     * @return bool
     */
    bool verify_file(string filename) {
        return File::exists(filename);
    }

    /**
     * @name test_configure
     * @brief expect that configure() will run as expected with mock Environment variables.
     *
     * @return bool
     */
    bool test_configure() {
        debug("test_configure_with_env_vars() start");

        Configuration cfg;
        const string initial_config = "test_config.yaml";

        namespace env = Environment;

        int argc = 7;
        char *argv[7] = {
                (char *) "TestConfigure",
                (char *) "--config", (char *) cfg_file.c_str(),
                (char *) "--logLevel", (char *) "debug",
                (char *) "--logOutput", (char *) "stdout"
        };

        return expect(env::set("CONFIG_FILE", initial_config), "set CONFIG_FILE") &&
               expect(env::set("LOG_LEVEL", "critical"), "set LOG_LEVEL") &&
               expect(env::set("LOG_OUTPUT", "stderr"), "set LOG_OUTPUT") &&
               expect(env::get("CONFIG_FILE").compare(initial_config) == 0,
                      "CONFIG_FILE ok:" + env::get("CONFIG_FILE")) &&

               expect(env::get("LOG_LEVEL").compare("critical") == 0, "LOG_LEVEL ok:" + env::get("LOG_LEVEL")) &&
               expect(env::get("LOG_OUTPUT").compare("stderr") == 0, "LOG_OUTPUT ok:" + env::get("LOG_OUTPUT"));

        expect(configure(&cfg, map, argc, argv), "configuration ok") &&

        expect(true, cfg.getString(ConfigFile)) &&
        expect(cfg.getString(ConfigFile).compare(cfg_file) == 0, "Expect config file ok");// &&
        expect(verify_file(cfg.getString(ConfigFile)), "Expect config file exists ok") &&
        expect(cfg.getString(ConfigLogLevel).compare("debug") == 0, "expect log level ok") &&
        expect(cfg.getString(ConfigLogOutput).compare("stdout") == 0, "expect log level ok");
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_configure(), "test_configure() ok");
    }
};