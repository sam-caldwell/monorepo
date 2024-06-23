/**
 * @name Logger/tests/TestLoggerSeverity.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */


#include <iostream>
#include "../src/Logger.h"

class TestLoggerSeverity : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestLoggerSeverity(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestLoggerSeverity() {}

    /**
     * @name test_severity
     * @brief Test the validity of the severity enumerated type for both values and data size (1 bytes).
     *
     * @return bool
     */
    bool test_severity() {
        return expect(int(Logging::Severity::emergency) == 0, "Expect Severity::emergency 0") &&
               expect(int(Logging::Severity::alert) == 1, "Expect Severity::alert 1") &&
               expect(int(Logging::Severity::critical) == 2, "Expect Severity::critical 2") &&
               expect(int(Logging::Severity::error) == 3, "Expect Severity::error 3") &&
               expect(int(Logging::Severity::warning) == 4, "Expect Severity::warning 4") &&
               expect(int(Logging::Severity::notice) == 5, "Expect Severity::notice 5") &&
               expect(int(Logging::Severity::info) == 6, "Expect Severity::info 6") &&
               expect(int(Logging::Severity::debug) == 7, "Expect Severity::debug 7") &&
               expect(sizeof(Logging::Severity) == 1, "Expect Logging::Severity::Severity is 1 byte");
    }

    /**
     * @name test_severity_strings
     * @brief Test the validity of the severity strings table.
     *
     * @return bool
     */
    bool test_severity_strings() {
        return expect(Logging::ReverseSeverity.at(Logging::Severity::emergency).compare("emergency") == 0,
                      "Expect emergency ok") &&
               expect(Logging::ReverseSeverity.at(Logging::Severity::alert).compare("alert") == 0, "Expect alert ok") &&
               expect(Logging::ReverseSeverity.at(Logging::Severity::critical).compare("critical") == 0,
                      "Expect critical ok") &&
               expect(Logging::ReverseSeverity.at(Logging::Severity::error).compare("error") == 0, "Expect error ok") &&
               expect(Logging::ReverseSeverity.at(Logging::Severity::warning).compare("warning") == 0,
                      "Expect warning ok") &&
               expect(Logging::ReverseSeverity.at(Logging::Severity::notice).compare("notice") == 0,
                      "Expect notice ok") &&
               expect(Logging::ReverseSeverity.at(Logging::Severity::info).compare("info") == 0, "Expect info ok") &&
               expect(Logging::ReverseSeverity.at(Logging::Severity::debug).compare("debug") == 0, "Expect debug ok");
    }

    /**
     * @name test_severity_from_string
     * @brief verify that Logging::Severity::SeverityFromString() will return a Logging::Severity::Severity
     *        with a given string.
     *
     * @return bool
     */
    bool test_severity_from_string() {
        return expect(Logging::SeverityStrings.at("emergency") == Logging::Severity::emergency,
                      "Expect emergency ok.") &&
               expect(Logging::SeverityStrings.at("alert") == Logging::Severity::alert, "Expect alert ok.") &&
               expect(Logging::SeverityStrings.at("critical") == Logging::Severity::critical, "Expect critical ok.") &&
               expect(Logging::SeverityStrings.at("error") == Logging::Severity::error, "Expect error ok.") &&
               expect(Logging::SeverityStrings.at("warning") == Logging::Severity::warning, "Expect warning ok.") &&
               expect(Logging::SeverityStrings.at("notice") == Logging::Severity::notice, "Expect notice ok.") &&
               expect(Logging::SeverityStrings.at("info") == Logging::Severity::info, "Expect info ok.") &&
               expect(Logging::SeverityStrings.at("debug") == Logging::Severity::debug, "Expect debug ok.");
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_severity(), "test_severity() ok") &&
               expect(test_severity_strings(), "test_severity_strings() ok") &&
               expect(test_severity_from_string(), "test_severity_from_string() ok");
    }

};