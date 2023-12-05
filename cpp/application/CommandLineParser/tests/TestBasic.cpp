/**
 * @name CommandLineParser/tests/TestVerifyData.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/application/CommandLineParser/src/CommandLineParser.h"

class TestBasic : public TestBase {
private:
    ConfigStateMap *map;
    Configuration *cfg;
    CommandLineParser *parser;
public:
    /**
     * @name class constructor
     * @brief initialize the test
     *
     * @param n string
     */
    TestBasic(string n) {
        debug(n + "::TestParse()");
        name = n;
        map = new ConfigStateMap;
        cfg = new Configuration;

        map->add("--help", "option.help", Bool, false, true);
        map->add("-h", "option.help", Bool, false, true);
        map->add("--option1", "option.a", String, false, false);
        map->add("--option2", "option.b", String, false, false);
        map->add("--option3", "option.c", Int, false, false);
    }

    /**
     * @name class destructor
     * @brief clean up the test
     */
    ~TestBasic() {
        debug(name + "::~TestParse()");
        if (parser) delete parser;
        if (map) delete map;
        if (cfg) delete cfg;
    }

    /**
     * @name test_instantiation
     * @brief can we instantiate the commandline parser object?
     *
     * @return bool
     */
    bool test_instantiation() {
        int argc = 3;
        char *argv[9] = {
                (char *) "TestParse",
                (char *) "--help",
                (char *) "--option1", (char *) "OptionValue1",
                (char *) "--option2", (char *) "OptionValue2",
                (char *) "--option3", (char *) "1337",
        };
        parser = new CommandLineParser(cfg, map, argc, argv);
        return expect(parser, "expect parser not null");
    }

    /**
     * @name main
     * @brief coordinate tests.
     *
     * @return bool
     */
    bool main() {
        return expect(test_instantiation(), "test_instantiation()");
    }
};