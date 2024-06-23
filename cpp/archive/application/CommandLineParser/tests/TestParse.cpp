/**
 * @name CommandLineParser/tests/TestParse.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../src/CommandLineParser.h"

class TestParse : public TestBase {
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
    TestParse(string n) {
        debug(n + "::TestParse()");
        name = n;
        map = new ConfigStateMap;
        cfg = new Configuration;

        map->add("--help", "option.help", Bool, false, true);
        map->add("-h", "option.help", Bool, false, true);
        map->add("--option1", "option.a", String, true, false);
        map->add("--option2", "option.b", String, true, false);
        map->add("--option3", "option.c", Int, true, false);
    }

    /**
     * @name class destructor
     * @brief clean up the test.
     */
    ~TestParse() {
        debug(name + "::~TestParse()");
        if (parser) delete parser;
        if (map) delete map;
        if (cfg) delete cfg;
    }

    /**
     * @name test_initial_state
     * @brief test the initial state post-construction.
     *
     * @return bool
     */
    bool test_initial_state() {
        return expect(cfg, "expect Configuration object not null") &&
               expect(map, "expect ConfigStateMap object not null") &&
               expect(parser == NULL, "expect CommandLineParser is null (not initialized)") &&
               expect(map->count() == 5, "expect map has five(5) entries.") &&
               expect(map->getRhsString("--help") == "option.help", "expect OK") &&
               expect(map->getRhsString("-h") == "option.help", "expect OK") &&
               expect(map->getRhsString("--option1") == "option.a", "expect option.a OK") &&
               expect(map->getRhsString("--option2") == "option.b", "expect option.b OK") &&
               expect(map->getRhsString("--option3") == "option.c", "expect option.c OK") &&
               expect([this]() -> bool {
                   auto key = this->map->getLhsBegin();
                   if (!expect(key->first == "--help", "Expect --help in pos 0. Actual:" + key->first)) return false;
                   key++;
                   if (!expect(key != map->getLhsEnd(), "Expect map iterator not at end 1.")) return false;
                   if (!expect(key->first == "--option1", "Expect --option1 in pos 1 Actual:" + key->first))
                       return false;
                   key++;
                   if (!expect(key != map->getLhsEnd(), "Expect map iterator not at end 2.")) return false;
                   if (!expect(key->first == "--option2", "Expect --option2 in pos 2 Actual:" + key->first))
                       return false;
                   key++;
                   if (!expect(key != map->getLhsEnd(), "Expect map iterator not at end 3.")) return false;
                   if (!expect(key->first == "--option3", "Expect --option3 in pos 3 Actual:" + key->first))
                       return false;
                   key++;
                   if (!expect(key != map->getLhsEnd(), "Expect map iterator not at end 4.")) return false;
                   if (!expect(key->first == "-h", "Expect -h in pos 4 Actual:" + key->first)) return false;
                   return true;
               }(), "map elements exist when manually iterating over the map.") &&
               expect([this]() -> bool {
                   int count = 0;
                   for (auto key = map->getLhsBegin(); key != map->getLhsEnd(); ++key) {
                       debug("lhs:" + key->first + " rhs:" + key->second);
                       count++;
                   }
                   return count == 5;
               }(), "expect iterator to return count of 5 nodes.");
    }

    /**
     * @name test_parse
     * @brief can we instantiate the commandline parser object and parse a given commandline argument set.?
     *
     * @return bool
     */
    bool test_parse() {
        //Note: if argc and argv don't match the actual count, you're going to get a segfault.
        int argc = 8;
        char *argv[8] = {
                (char *) "TestParse",
                (char *) "--help",
                (char *) "--option1", (char *) "OptionValue1",
                (char *) "--option2", (char *) "OptionValue2",
                (char *) "--option3", (char *) "1337",
        };
        try {
            parser = new CommandLineParser(cfg, map, argc, argv);
            debug("parser initialization completed");
            return expect(parser, "expect parser not null");
        } catch (exception &e) {
            return expect(false, "expect parser to experience no errors");
        }
        return true;
    }

    /**
     * @name test_post_parse_state
     * @brief Expect that after parsing is complete cfg will contain the expected data.
     *
     * @return bool
     */
    bool test_post_parse_state() {
        return expect(cfg->count() == 4, "Expect four objects in Configuration object") &&
               expect(cfg->getBool("option.help"), "Expect option.help is set to true") &&
               expect(cfg->getString("option.a").compare("OptionValue1")==0, "Expect option.a ok") &&
               expect(cfg->getString("option.b").compare("OptionValue2")==0, "Expect option.b ok") &&
               expect(cfg->getInt("option.c")==1337, "Expect option.c ok");
    }

    /**
     * @name main
     * @brief coordinate tests.
     *
     * @return bool
     */
    bool main() {
        return expect(test_initial_state(), "expect test_initial_state() ok") &&
               expect(test_parse(), "test_parse() ok") &&
               expect(test_post_parse_state(), "test_post_parse_state() ok");
    }
};