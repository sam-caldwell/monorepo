/**
 * @name ConfigFileParser/tests/TestParseYamlConfig.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../src/ConfigFileParser.h"
#include "../../../../formatter/formatter.h"

class TestParseYamlConfig : public TestBase {
private:
    Configuration *cfg;
    ConfigStateMap *map;
    ConfigFileParser *parser;
    const string cfg_file = "good_basic_config.yaml";
public:
    /**
     * @name class constructor
     * @brief setup the test.
     *
     * @param n string
     */
    TestParseYamlConfig(string n) {
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
    ~TestParseYamlConfig() {
        if (cfg) delete cfg;
        if (map) delete map;
    }

    /**
     * @name test_initial_state
     * @brief Expect that in the initial state parser is NULL, map is configured, configuration object is
     *        configured.
     *
     * @return bool
     */
    bool test_initial_state() {
        return expect(cfg, "expect cfg is not null") &&
               expect(map, "expect map is not null") &&
               expect(cfg->count() == 1, "expect initial config count = 1. Actual:" + to_string(cfg->count()));
    }

    /**
     * @name test_parse
     * @brief Expect that parser will parse file on instantiation.
     *
     * @return bool
     */
    bool test_parse() {
        string fileName = string(cfg_file);
        debug("       filename: " + fileName);
        debug("config filename: " + cfg->getString("config.file"));
        parser = new ConfigFileParser(cfg, map);
        return expect(parser, "Expect parser is not null") &&
               expect(cfg->count() == 3, "expect count 3 in configuration object.  Got:" + to_string(cfg->count())) &&
               expect(cfg->getString("secA.key").compare("value0")==0, "Expect secA.key contains string 'value0'");
    }

    bool main() {
        debug(name + "::main()");
        return expect(test_parse(), "Expect test_parse() ok");
    }

}; /*end of class*/