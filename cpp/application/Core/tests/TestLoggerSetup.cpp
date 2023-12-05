/**
 * @name Application/tests/TestLoggerSetup.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/system/file.h"
#include "projects/Tester/Mocks/MockStdout.h"
#include "projects/logging/Logger/src/Logger.h"
#include "projects/common/system/environment.h"
#include "projects/application/Core/src/configure.cpp"

class TestLoggerSetup : public TestBase {
private:
    string msg;
    Logger *log;
    Configuration *cfg;
    ConfigStateMap *map;
    string tags;
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestLoggerSetup(string n) {
        name = n;
        msg = "test_message";
        tags = "\"tag:1\"";
    }

    ~TestLoggerSetup() {}

    bool setup_logger() {
        map = new ConfigStateMap;
        cfg = new Configuration();

        map->add(ConfigLogLevel, ConfigLogLevel, String, false, true);
        map->add(ConfigLogOutput, ConfigLogOutput, String, false, true);
        map->add(ConfigLogTags, ConfigLogTags, String, false, true);

        cfg->loadData(map, ConfigLogLevel, "debug");
        cfg->loadData(map, ConfigLogOutput, "stdout");
        cfg->loadData(map, ConfigLogTags, tags);

        log = new Logger(cfg);
        return true;
    }

    bool log_by_value() {
        Mock::Stdout mock;
        mock.enable();
        log->debug(msg);
        mock.capture();
        string line = mock.last();
        return expect(line.length() == 71, "line(" + to_string(line.length()) + "):" + line);
    }

    bool log_by_reference() {
        Mock::Stdout mock;
        mock.enable();
        log->debug(&msg);
        mock.capture();
        string line = mock.last();
        return expect(line.length() == 71, "line(" + to_string(line.length()) + "):" + line);
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");

        return expect(setup_logger(), "logger ok") &&
               expect(log_by_value(), "log_by_value() ok") &&
               expect(log_by_reference(), "log_by_reference() ok");
    }
};