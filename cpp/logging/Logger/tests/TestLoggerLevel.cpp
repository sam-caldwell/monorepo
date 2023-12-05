/**
 * @name Logger/tests/TestLoggerLevel.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */


#include <iostream>
#include "projects/logging/Logger/src/Logger.h"

class TestLoggerLevel : public TestBase {
private:
    const Logging::Writer::Id writerNull = Logging::Writer::null;
    const Logging::Writer::Id writerStdout = Logging::Writer::stdout;
    const Logging::Writer::Id writerStderr = Logging::Writer::stderr;
    const Logging::Writer::Id writerFile = Logging::Writer::file;
    const Logging::Writer::Id writerSyslog = Logging::Writer::syslog;
    const Logging::Writer::Id writerHttps = Logging::Writer::https;

public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestLoggerLevel(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestLoggerLevel() {/*noop*/}

    /**
     * @name scenario0
     * @brief Test log severity when logger is created with defaults.
     *
     * @return bool
     */
    bool scenario0() {
        Logger log;
        return expect(log.severity() == Logging::Severity::info, "severity ok() ok");
    }

    /**
     * @name scenario1
     * @brief Test log severity when logger is created with only sev, writer.
     *
     * @return bool
     */
    bool scenario1(Logging::Severity sev, Logging::Writer::Id writer) {
        Logger log(sev, writer);
        return expect(log.severity() == sev, "severity ok() ok");
    }

    /**
     * @name scenario2
     * @brief Test log severity when logger is created with only sev, writer and tags.
     *
     * @return bool
     */
    bool scenario2(Logging::Severity sev, Logging::Writer::Id writer, string *tags) {
        Logger log(sev, writer, *tags);
        return expect(log.severity() == sev, "severity ok() ok");
    }

    /**
     * @name scenario3
     * @brief Test log severity when logger is created with all parameters.
     *
     * @return bool
     */
    bool scenario3(Logging::Severity sev, Logging::Writer::Id writer, string *tags, string *destination) {
        Logger log(sev, writer, *tags, *destination);
        return expect(log.severity() == sev, "severity ok() ok");
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");

        string tags = "\"tag:1\",\"tag:2\",\"tag:3\",\"tag:4\"";
        string destination = "file:///dev/null";

        return expect(scenario0(), "Expect scenario0 ok") &&
               expect(scenario1(Logging::Severity::alert, writerNull), "Expect scenario 1.0 ok") &&
               expect(scenario1(Logging::Severity::critical, writerNull), "Expect scenario 1.1 ok") &&
               expect(scenario1(Logging::Severity::error, writerNull), "Expect scenario 1.2 ok") &&
               expect(scenario1(Logging::Severity::warning, writerNull), "Expect scenario 1.3 ok") &&
               expect(scenario1(Logging::Severity::notice, writerNull), "Expect scenario 1.4 ok") &&
               expect(scenario1(Logging::Severity::info, writerNull), "Expect scenario 1.5 ok") &&
               expect(scenario1(Logging::Severity::debug, writerNull), "Expect scenario 1.6 ok") &&

               expect(scenario2(Logging::Severity::alert, writerNull, &tags), "Expect scenario 2.0 ok") &&
               expect(scenario2(Logging::Severity::critical, writerNull, &tags), "Expect scenario 2.1 ok") &&
               expect(scenario2(Logging::Severity::error, writerNull, &tags), "Expect scenario 2.2 ok") &&
               expect(scenario2(Logging::Severity::warning, writerNull, &tags), "Expect scenario 2.3 ok") &&
               expect(scenario2(Logging::Severity::notice, writerNull, &tags), "Expect scenario 2.4 ok") &&
               expect(scenario2(Logging::Severity::info, writerNull, &tags), "Expect scenario 2.5 ok") &&
               expect(scenario2(Logging::Severity::debug, writerNull, &tags), "Expect scenario 2.6 ok") &&

               expect(scenario3(Logging::Severity::alert, writerNull, &tags, &destination), "scenario 3.0 ok") &&
               expect(scenario3(Logging::Severity::critical, writerNull, &tags, &destination), "scenario 3.1 ok") &&
               expect(scenario3(Logging::Severity::error, writerNull, &tags, &destination), "scenario 3.2 ok") &&
               expect(scenario3(Logging::Severity::warning, writerNull, &tags, &destination), "scenario 3.3 ok") &&
               expect(scenario3(Logging::Severity::notice, writerNull, &tags, &destination), "scenario 3.4 ok") &&
               expect(scenario3(Logging::Severity::info, writerNull, &tags, &destination), "scenario 3.5 ok") &&
               expect(scenario3(Logging::Severity::debug, writerNull, &tags, &destination), "scenario 3.6 ok");
    }
};