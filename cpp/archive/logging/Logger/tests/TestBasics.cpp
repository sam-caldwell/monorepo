/**
 * @name Logger/tests/TestBasics.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "../../../../networking/TcpCommon/src/PipeResponse/main.h"

#include <iostream>
#include "../src/Logger.h"

class TestBasics : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestBasics(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestBasics() {}

    /**
     * @name instantiation_scenario_1
     * @brief instantiate Logger with defaults (as pointer)
     *
     * @return bool
     */
    bool instantiation_scenario_1() {
        Logger *log;
        log = new Logger();
        return expect(log, "expect log not null");
    }

    /**
     * @name instantiation_scenario_2
     * @brief instantiate logger with defaults (as object)
     *
     * @return bool
     */
    bool instantiation_scenario_2() {
        Logger log;
        return true;
    }

    /**
     * @name instantiation_scenario_3
     * @brief instantiate logger with severity, writerId
     *
     * @return bool
     */
    bool instantiation_scenario_3() {
        Logger log(Logging::Severity::debug, Logging::Writer::Id::stdout);
        return true;
    }

    /**
     * @name instantiation_scenario_4
     * @brief instantiate logger with severity, writerId, tags
     *
     * @return bool
     */
    bool instantiation_scenario_4() {
        Logger log(Logging::Severity::debug,
                   Logging::Writer::Id::stdout,
                   "\"Foo\",\"bar\"");
        return true;
    }

    /**
     * @name instantiation_scenario_5
     * @brief instantiate logger with severity, writerId, tags and destination
     *
     * @return bool
     */
    bool instantiation_scenario_5() {
        Logger log(Logging::Severity::debug,
                   Logging::Writer::Id::stdout,
                   "\"Foo:bar\",\"Moo:mar\"",
                   "/var/log/myLog.log");
        return true;
    }

    /**
     * @name instantiation_scenario_6
     * @brief instantiate logger with all parameters but an invalid tag.
     *        expect out_of_range exception.
     *
     * @return bool
     */
    bool instantiation_scenario_6() {
        try {
            Logger log(Logging::Severity::debug,
                       Logging::Writer::Id::stdout,
                       "\"Foo\":\"bar\",\"Moo\":\"mar\"",
                       "/var/log/myLog.log");
            return expect(false, "Expect out_of_range exception");
        } catch (out_of_range &e) {
            const string errMsg = "Logger Configuration failed(3): Logger cannot validate given tag set.";
            return expect(true, "Expect out_of_range exception") &&
                   expect(string(e.what()).compare(errMsg) == 0, "Expected error message ok: " + string(e.what()));
        }
        return expect(false, "Expected out_of_range exception.  Got something unexpected.");
    }

    /**
     * @name instantiation_scenario_7
     * @brief instantiate logger with all parameters but an invalid tag and no destination.
     *        expect out_of_range exception.
     *
     * @return bool
     */
    bool instantiation_scenario_7() {
        try {
            Logger log(Logging::Severity::debug,
                       Logging::Writer::Id::stdout,
                       "\"Foo\":\"bar\",\"Moo\":\"mar\"");
            return expect(false, "Expect out_of_range exception");
        } catch (out_of_range &e) {
            const string errMsg = "Logger Configuration failed(2): Logger cannot validate given tag set.";
            return expect(true, "Expect out_of_range exception") &&
                   expect(string(e.what()).compare(errMsg) == 0, "Expected error message ok: " + string(e.what()));
        }
        return expect(false, "Expected out_of_range exception.  Got something unexpected.");
    }

    /**
     * @name instantiation_scenario_8
     * @brief instantiate logger with Configuration object
     *        expect out_of_range exception.
     *
     * @return bool
     */
    bool instantiation_scenario_8() {
        ConfigStateMap map;
        map.add("log.level", "log.level", String);
        map.add("log.output", "log.output", String);
        map.add("log.destination", "log.destination", String);
        map.add("log.tags", "log.tags", String);

        Configuration cfg;
        cfg.loadData(&map, "log.level", "debug");
        cfg.loadData(&map, "log.output", "stdout");
        cfg.loadData(&map, "log.tags", "\"Foo:bar\",\"Moo:mar\"");
        cfg.loadData(&map, "log.destination", "");
        Logger log(&cfg);
        return true;
    }

    /**
     * @name instantiation_scenario_9
     * @brief instantiate logger with Configuration object
     *        expect out_of_range exception.
     *
     * @return bool
     */
    bool instantiation_scenario_9() {
        ConfigStateMap map;
        map.add("log.level", "log.level", String);
        map.add("log.output", "log.output", String);
        map.add("log.destination", "log.destination", String);
        map.add("log.tags", "log.tags", String);

        Configuration cfg;
        cfg.loadData(&map, "log.level", "debug");
        cfg.loadData(&map, "log.output", "stdout");
        cfg.loadData(&map, "log.tags", "\"Foo\":\"bar\",\"Moo\":\"mar\"");
        cfg.loadData(&map, "log.destination", "");

        try {
            Logger log(&cfg);
            return expect(false, "Expect out_of_range exception");
        } catch (out_of_range &e) {
            const string errMsg = "Logger configuration failed(1): Logger cannot validate given tag set.";
            return expect(true, "Expect out_of_range exception") &&
                   expect(string(e.what()).compare(errMsg) == 0, "Expected error message ok: " + string(e.what()));
        }
        return expect(false, "Expected out_of_range exception.  Got something unexpected.");
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(instantiation_scenario_1(), "instantiation_scenario_1() ok") &&
               expect(instantiation_scenario_2(), "instantiation_scenario_2() ok") &&
               expect(instantiation_scenario_3(), "instantiation_scenario_3() ok") &&
               expect(instantiation_scenario_4(), "instantiation_scenario_4() ok") &&
               expect(instantiation_scenario_5(), "instantiation_scenario_5() ok") &&
               expect(instantiation_scenario_6(), "instantiation_scenario_6() ok") &&
               expect(instantiation_scenario_7(), "instantiation_scenario_7() ok") &&
               expect(instantiation_scenario_8(), "instantiation_scenario_8() ok") &&
               expect(instantiation_scenario_9(), "instantiation_scenario_9() ok");
    }

};