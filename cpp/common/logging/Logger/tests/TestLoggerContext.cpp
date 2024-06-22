/**
 * @name Logger/tests/TestLoggerContext.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../src/Logger.h"

class TestLoggerContext : public TestBase {
    const string tags = "\"tag:1\",\"tag:2\",\"tag:3\",\"tag:4\"";
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestLoggerContext(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestLoggerContext() {}

    /**
     * @name test_context_nullWriter
     * @brief Test that we can instantiate Logger class can instantiate with no error.
     *
     * @return bool
     */
    bool test_context_nullWriter() {
        Logging::Writer::Id id = Logging::Writer::null;

        Logger c0(Logging::Severity::emergency, id);
        Logger c1(Logging::Severity::alert, id);
        Logger c2(Logging::Severity::critical, id);
        Logger c3(Logging::Severity::error, id);
        Logger c4(Logging::Severity::warning, id);
        Logger c5(Logging::Severity::notice, id);
        Logger c6(Logging::Severity::info, id);
        Logger c7(Logging::Severity::debug, id);

        return expect(c0.severity() == Logging::Severity::emergency, "c0 emergency ok") &&
               expect(c1.severity() == Logging::Severity::alert, "c1 alert ok") &&
               expect(c2.severity() == Logging::Severity::critical, "c2 critical ok") &&
               expect(c3.severity() == Logging::Severity::error, "c3 error ok") &&
               expect(c4.severity() == Logging::Severity::warning, "c4 warning ok") &&
               expect(c5.severity() == Logging::Severity::notice, "c5 notice ok") &&
               expect(c6.severity() == Logging::Severity::info, "c6 info ok") &&
               expect(c7.severity() == Logging::Severity::debug, "c7 debug ok") &&

                expect(c0.writerId() == id, "expect c0 writerId ok") &&
                expect(c1.writerId() == id, "expect c1 writerId ok") &&
                expect(c2.writerId() == id, "expect c2 writerId ok") &&
                expect(c3.writerId() == id, "expect c3 writerId ok") &&
                expect(c4.writerId() == id, "expect c4 writerId ok") &&
                expect(c5.writerId() == id, "expect c5 writerId ok") &&
                expect(c6.writerId() == id, "expect c6 writerId ok") &&
                expect(c7.writerId() == id, "expect c7 writerId ok") &&

                expect(c0.tags().compare("") == 0, "expect c0 tags ok") &&
                expect(c1.tags().compare("") == 0, "expect c1 tags ok") &&
                expect(c2.tags().compare("") == 0, "expect c2 tags ok") &&
                expect(c3.tags().compare("") == 0, "expect c3 tags ok") &&
                expect(c4.tags().compare("") == 0, "expect c4 tags ok") &&
                expect(c5.tags().compare("") == 0, "expect c5 tags ok") &&
                expect(c6.tags().compare("") == 0, "expect c6 tags ok") &&
                expect(c7.tags().compare("") == 0, "expect c7 tags ok") &&

                expect(c0.destination().compare("") == 0, "expect c0 destination ok") &&
                expect(c1.destination().compare("") == 0, "expect c1 destination ok") &&
                expect(c2.destination().compare("") == 0, "expect c2 destination ok") &&
                expect(c3.destination().compare("") == 0, "expect c3 destination ok") &&
                expect(c4.destination().compare("") == 0, "expect c4 destination ok") &&
                expect(c5.destination().compare("") == 0, "expect c5 destination ok") &&
                expect(c6.destination().compare("") == 0, "expect c6 destination ok") &&
                expect(c7.destination().compare("") == 0, "expect c7 destination ok");
    }

    /**
     * @name test_context_stdoutWriter
     * @brief Test that we can instantiate Logger class can instantiate with no error.
     *
     * @return bool
     */
    bool test_context_stdoutWriter() {
        Logging::Writer::Id id = Logging::Writer::stdout;

        Logger c0(Logging::Severity::emergency, id, tags);
        Logger c1(Logging::Severity::alert, id, tags);
        Logger c2(Logging::Severity::critical, id, tags);
        Logger c3(Logging::Severity::error, id, tags);
        Logger c4(Logging::Severity::warning, id, tags);
        Logger c5(Logging::Severity::notice, id, tags);
        Logger c6(Logging::Severity::info, id, tags);
        Logger c7(Logging::Severity::debug, id, tags);

        return expect(c0.severity() == Logging::Severity::emergency, "c0 emergency ok") &&
               expect(c1.severity() == Logging::Severity::alert, "c1 alert ok") &&
               expect(c2.severity() == Logging::Severity::critical, "c2 critical ok") &&
               expect(c3.severity() == Logging::Severity::error, "c3 error ok") &&
               expect(c4.severity() == Logging::Severity::warning, "c4 warning ok") &&
               expect(c5.severity() == Logging::Severity::notice, "c5 notice ok") &&
               expect(c6.severity() == Logging::Severity::info, "c6 info ok") &&
               expect(c7.severity() == Logging::Severity::debug, "c7 debug ok") &&

               expect(c0.writerId() == id, "expect c0 writerId ok") &&
               expect(c1.writerId() == id, "expect c1 writerId ok") &&
               expect(c2.writerId() == id, "expect c2 writerId ok") &&
               expect(c3.writerId() == id, "expect c3 writerId ok") &&
               expect(c4.writerId() == id, "expect c4 writerId ok") &&
               expect(c5.writerId() == id, "expect c5 writerId ok") &&
               expect(c6.writerId() == id, "expect c6 writerId ok") &&
               expect(c7.writerId() == id, "expect c7 writerId ok") &&

                expect(c0.tags().compare(tags) == 0, "expect c0 tags ok") &&
                expect(c1.tags().compare(tags) == 0, "expect c1 tags ok") &&
                expect(c2.tags().compare(tags) == 0, "expect c2 tags ok") &&
                expect(c3.tags().compare(tags) == 0, "expect c3 tags ok") &&
                expect(c4.tags().compare(tags) == 0, "expect c4 tags ok") &&
                expect(c5.tags().compare(tags) == 0, "expect c5 tags ok") &&
                expect(c6.tags().compare(tags) == 0, "expect c6 tags ok") &&
                expect(c7.tags().compare(tags) == 0, "expect c7 tags ok") &&

                expect(c0.destination().compare("") == 0, "expect c0 destination ok") &&
                expect(c1.destination().compare("") == 0, "expect c1 destination ok") &&
                expect(c2.destination().compare("") == 0, "expect c2 destination ok") &&
                expect(c3.destination().compare("") == 0, "expect c3 destination ok") &&
                expect(c4.destination().compare("") == 0, "expect c4 destination ok") &&
                expect(c5.destination().compare("") == 0, "expect c5 destination ok") &&
                expect(c6.destination().compare("") == 0, "expect c6 destination ok") &&
                expect(c7.destination().compare("") == 0, "expect c7 destination ok");
    }

    /**
     * @name test_context_stderrWriter
     * @brief Test that we can instantiate Logger class can instantiate with no error.
     *
     * @return bool
     */
    bool test_context_stderrWriter() {
        Logging::Writer::Id id = Logging::Writer::stderr;

        Logger c0(Logging::Severity::emergency, id, tags);
        Logger c1(Logging::Severity::alert, id, tags);
        Logger c2(Logging::Severity::critical, id, tags);
        Logger c3(Logging::Severity::error, id, tags);
        Logger c4(Logging::Severity::warning, id, tags);
        Logger c5(Logging::Severity::notice, id, tags);
        Logger c6(Logging::Severity::info, id, tags);
        Logger c7(Logging::Severity::debug, id, tags);

        return expect(c0.severity() == Logging::Severity::emergency, "c0 emergency ok") &&
               expect(c1.severity() == Logging::Severity::alert, "c1 alert ok") &&
               expect(c2.severity() == Logging::Severity::critical, "c2 critical ok") &&
               expect(c3.severity() == Logging::Severity::error, "c3 error ok") &&
               expect(c4.severity() == Logging::Severity::warning, "c4 warning ok") &&
               expect(c5.severity() == Logging::Severity::notice, "c5 notice ok") &&
               expect(c6.severity() == Logging::Severity::info, "c6 info ok") &&
               expect(c7.severity() == Logging::Severity::debug, "c7 debug ok") &&

               expect(c0.writerId() == id, "expect c0 writerId ok") &&
               expect(c1.writerId() == id, "expect c1 writerId ok") &&
               expect(c2.writerId() == id, "expect c2 writerId ok") &&
               expect(c3.writerId() == id, "expect c3 writerId ok") &&
               expect(c4.writerId() == id, "expect c4 writerId ok") &&
               expect(c5.writerId() == id, "expect c5 writerId ok") &&
               expect(c6.writerId() == id, "expect c6 writerId ok") &&
               expect(c7.writerId() == id, "expect c7 writerId ok") &&

               expect(c0.tags().compare(tags) == 0, "expect c0 tags ok") &&
               expect(c1.tags().compare(tags) == 0, "expect c1 tags ok") &&
               expect(c2.tags().compare(tags) == 0, "expect c2 tags ok") &&
               expect(c3.tags().compare(tags) == 0, "expect c3 tags ok") &&
               expect(c4.tags().compare(tags) == 0, "expect c4 tags ok") &&
               expect(c5.tags().compare(tags) == 0, "expect c5 tags ok") &&
               expect(c6.tags().compare(tags) == 0, "expect c6 tags ok") &&
               expect(c7.tags().compare(tags) == 0, "expect c7 tags ok") &&

               expect(c0.destination().compare("") == 0, "expect c0 destination ok") &&
               expect(c1.destination().compare("") == 0, "expect c1 destination ok") &&
               expect(c2.destination().compare("") == 0, "expect c2 destination ok") &&
               expect(c3.destination().compare("") == 0, "expect c3 destination ok") &&
               expect(c4.destination().compare("") == 0, "expect c4 destination ok") &&
               expect(c5.destination().compare("") == 0, "expect c5 destination ok") &&
               expect(c6.destination().compare("") == 0, "expect c6 destination ok") &&
               expect(c7.destination().compare("") == 0, "expect c7 destination ok");
    }

    /**
     * @name test_context_fileWriter
     * @brief Test that we can instantiate Logger class can instantiate with no error.
     *
     * @return bool
     */
    bool test_context_fileWriter() {
        Logging::Writer::Id id = Logging::Writer::file;
        string testLogFile = "./artifacts/test.log.txt";

        Logger c0(Logging::Severity::emergency, id, tags, testLogFile);
        Logger c1(Logging::Severity::alert, id, tags, testLogFile);
        Logger c2(Logging::Severity::critical, id, tags, testLogFile);
        Logger c3(Logging::Severity::error, id, tags, testLogFile);
        Logger c4(Logging::Severity::warning, id, tags, testLogFile);
        Logger c5(Logging::Severity::notice, id, tags, testLogFile);
        Logger c6(Logging::Severity::info, id, tags, testLogFile);
        Logger c7(Logging::Severity::debug, id, tags, testLogFile);

        return expect(c0.severity() == Logging::Severity::emergency, "c0 emergency ok") &&
               expect(c1.severity() == Logging::Severity::alert, "c1 alert ok") &&
               expect(c2.severity() == Logging::Severity::critical, "c2 critical ok") &&
               expect(c3.severity() == Logging::Severity::error, "c3 error ok") &&
               expect(c4.severity() == Logging::Severity::warning, "c4 warning ok") &&
               expect(c5.severity() == Logging::Severity::notice, "c5 notice ok") &&
               expect(c6.severity() == Logging::Severity::info, "c6 info ok") &&
               expect(c7.severity() == Logging::Severity::debug, "c7 debug ok") &&

               expect(c0.writerId() == id, "expect c0 writerId ok") &&
               expect(c1.writerId() == id, "expect c1 writerId ok") &&
               expect(c2.writerId() == id, "expect c2 writerId ok") &&
               expect(c3.writerId() == id, "expect c3 writerId ok") &&
               expect(c4.writerId() == id, "expect c4 writerId ok") &&
               expect(c5.writerId() == id, "expect c5 writerId ok") &&
               expect(c6.writerId() == id, "expect c6 writerId ok") &&
               expect(c7.writerId() == id, "expect c7 writerId ok") &&

               expect(c0.tags().compare(tags) == 0, "expect c0 tags ok") &&
               expect(c1.tags().compare(tags) == 0, "expect c1 tags ok") &&
               expect(c2.tags().compare(tags) == 0, "expect c2 tags ok") &&
               expect(c3.tags().compare(tags) == 0, "expect c3 tags ok") &&
               expect(c4.tags().compare(tags) == 0, "expect c4 tags ok") &&
               expect(c5.tags().compare(tags) == 0, "expect c5 tags ok") &&
               expect(c6.tags().compare(tags) == 0, "expect c6 tags ok") &&
               expect(c7.tags().compare(tags) == 0, "expect c7 tags ok") &&

               expect(c0.destination().compare(testLogFile) == 0, "expect c0 destination ok") &&
               expect(c1.destination().compare(testLogFile) == 0, "expect c1 destination ok") &&
               expect(c2.destination().compare(testLogFile) == 0, "expect c2 destination ok") &&
               expect(c3.destination().compare(testLogFile) == 0, "expect c3 destination ok") &&
               expect(c4.destination().compare(testLogFile) == 0, "expect c4 destination ok") &&
               expect(c5.destination().compare(testLogFile) == 0, "expect c5 destination ok") &&
               expect(c6.destination().compare(testLogFile) == 0, "expect c6 destination ok") &&
               expect(c7.destination().compare(testLogFile) == 0, "expect c7 destination ok");
    }

    /**
     * @name test_context_syslogWriter
     * @brief Test that we can instantiate Logger class can instantiate with no error.
     *
     * @return bool
     */
    bool test_context_syslogWriter() {
        Logging::Writer::Id id = Logging::Writer::syslog;
        string syslogEndpoint = "syslog://127.0.0.1:514";

        debug("test_context_syslogWriter() start");
        Logger c0(Logging::Severity::emergency, id, tags, syslogEndpoint);
        Logger c1(Logging::Severity::alert, id, tags, syslogEndpoint);
        Logger c2(Logging::Severity::critical, id, tags, syslogEndpoint);
        Logger c3(Logging::Severity::error, id, tags, syslogEndpoint);
        Logger c4(Logging::Severity::warning, id, tags, syslogEndpoint);
        Logger c5(Logging::Severity::notice, id, tags, syslogEndpoint);
        Logger c6(Logging::Severity::info, id, tags, syslogEndpoint);
        Logger c7(Logging::Severity::debug, id, tags, syslogEndpoint);

        return expect(true, "syslog test starting") &&
               expect(c0.severity() == Logging::Severity::emergency, "c0 emergency ok") &&
               expect(c1.severity() == Logging::Severity::alert, "c1 alert ok") &&
               expect(c2.severity() == Logging::Severity::critical, "c2 critical ok") &&
               expect(c3.severity() == Logging::Severity::error, "c3 error ok") &&
               expect(c4.severity() == Logging::Severity::warning, "c4 warning ok") &&
               expect(c5.severity() == Logging::Severity::notice, "c5 notice ok") &&
               expect(c6.severity() == Logging::Severity::info, "c6 info ok") &&
               expect(c7.severity() == Logging::Severity::debug, "c7 debug ok") &&

               expect(c0.writerId() == id, "expect c0 writerId ok") &&
               expect(c1.writerId() == id, "expect c1 writerId ok") &&
               expect(c2.writerId() == id, "expect c2 writerId ok") &&
               expect(c3.writerId() == id, "expect c3 writerId ok") &&
               expect(c4.writerId() == id, "expect c4 writerId ok") &&
               expect(c5.writerId() == id, "expect c5 writerId ok") &&
               expect(c6.writerId() == id, "expect c6 writerId ok") &&
               expect(c7.writerId() == id, "expect c7 writerId ok") &&

               expect(c0.tags().compare(tags) == 0, "expect c0 tags ok") &&
               expect(c1.tags().compare(tags) == 0, "expect c1 tags ok") &&
               expect(c2.tags().compare(tags) == 0, "expect c2 tags ok") &&
               expect(c3.tags().compare(tags) == 0, "expect c3 tags ok") &&
               expect(c4.tags().compare(tags) == 0, "expect c4 tags ok") &&
               expect(c5.tags().compare(tags) == 0, "expect c5 tags ok") &&
               expect(c6.tags().compare(tags) == 0, "expect c6 tags ok") &&
               expect(c7.tags().compare(tags) == 0, "expect c7 tags ok") &&

               expect(c0.destination().compare(syslogEndpoint) == 0, "expect c0 destination ok") &&
               expect(c1.destination().compare(syslogEndpoint) == 0, "expect c1 destination ok") &&
               expect(c2.destination().compare(syslogEndpoint) == 0, "expect c2 destination ok") &&
               expect(c3.destination().compare(syslogEndpoint) == 0, "expect c3 destination ok") &&
               expect(c4.destination().compare(syslogEndpoint) == 0, "expect c4 destination ok") &&
               expect(c5.destination().compare(syslogEndpoint) == 0, "expect c5 destination ok") &&
               expect(c6.destination().compare(syslogEndpoint) == 0, "expect c6 destination ok") &&
               expect(c7.destination().compare(syslogEndpoint) == 0, "expect c7 destination ok");
    }

    /**
     * @name test_context_httpsWriter
     * @brief Test that we can instantiate Logger class can instantiate with no error.
     *
     * @return bool
     */
    bool test_context_httpsWriter() {
        Logging::Writer::Id id = Logging::Writer::https;
        string httpsEndpoint = "https://127.0.0.1:8514";

        Logger c0(Logging::Severity::emergency, id, tags, httpsEndpoint);
        Logger c1(Logging::Severity::alert, id, tags, httpsEndpoint);
        Logger c2(Logging::Severity::critical, id, tags, httpsEndpoint);
        Logger c3(Logging::Severity::error, id, tags, httpsEndpoint);
        Logger c4(Logging::Severity::warning, id, tags, httpsEndpoint);
        Logger c5(Logging::Severity::notice, id, tags, httpsEndpoint);
        Logger c6(Logging::Severity::info, id, tags, httpsEndpoint);
        Logger c7(Logging::Severity::debug, id, tags, httpsEndpoint);

        return expect(c0.severity() == Logging::Severity::emergency, "c0 emergency ok") &&
               expect(c1.severity() == Logging::Severity::alert, "c1 alert ok") &&
               expect(c2.severity() == Logging::Severity::critical, "c2 critical ok") &&
               expect(c3.severity() == Logging::Severity::error, "c3 error ok") &&
               expect(c4.severity() == Logging::Severity::warning, "c4 warning ok") &&
               expect(c5.severity() == Logging::Severity::notice, "c5 notice ok") &&
               expect(c6.severity() == Logging::Severity::info, "c6 info ok") &&
               expect(c7.severity() == Logging::Severity::debug, "c7 debug ok") &&

               expect(c0.writerId() == id, "expect c0 writerId ok") &&
               expect(c1.writerId() == id, "expect c1 writerId ok") &&
               expect(c2.writerId() == id, "expect c2 writerId ok") &&
               expect(c3.writerId() == id, "expect c3 writerId ok") &&
               expect(c4.writerId() == id, "expect c4 writerId ok") &&
               expect(c5.writerId() == id, "expect c5 writerId ok") &&
               expect(c6.writerId() == id, "expect c6 writerId ok") &&
               expect(c7.writerId() == id, "expect c7 writerId ok") &&

               expect(c0.tags().compare(tags) == 0, "expect c0 tags ok") &&
               expect(c1.tags().compare(tags) == 0, "expect c1 tags ok") &&
               expect(c2.tags().compare(tags) == 0, "expect c2 tags ok") &&
               expect(c3.tags().compare(tags) == 0, "expect c3 tags ok") &&
               expect(c4.tags().compare(tags) == 0, "expect c4 tags ok") &&
               expect(c5.tags().compare(tags) == 0, "expect c5 tags ok") &&
               expect(c6.tags().compare(tags) == 0, "expect c6 tags ok") &&
               expect(c7.tags().compare(tags) == 0, "expect c7 tags ok") &&

               expect(c0.destination().compare(httpsEndpoint) == 0, "expect c0 destination ok") &&
               expect(c1.destination().compare(httpsEndpoint) == 0, "expect c1 destination ok") &&
               expect(c2.destination().compare(httpsEndpoint) == 0, "expect c2 destination ok") &&
               expect(c3.destination().compare(httpsEndpoint) == 0, "expect c3 destination ok") &&
               expect(c4.destination().compare(httpsEndpoint) == 0, "expect c4 destination ok") &&
               expect(c5.destination().compare(httpsEndpoint) == 0, "expect c5 destination ok") &&
               expect(c6.destination().compare(httpsEndpoint) == 0, "expect c6 destination ok") &&
               expect(c7.destination().compare(httpsEndpoint) == 0, "expect c7 destination ok");
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return debug("---step 1----") &&
               expect(test_context_nullWriter(), "test_context_nullWriter() ok") &&
               debug("---step 2----") &&
               expect(test_context_stdoutWriter(), "test_context_stdoutWriter() ok") &&
                debug("---step 3----") &&
               expect(test_context_stderrWriter(), "test_context_stderrWriter() ok") &&
               debug("---step 4----") &&
               expect(test_context_fileWriter(), "test_context_fileWriter() ok") &&
               debug("---step 5----") &&
               expect(test_context_syslogWriter(), "test_context_syslogWriter() ok") &&
               debug("---step 6----") &&
               expect(test_context_httpsWriter(), "test_context_httpsWriter() ok") &&
               debug("----done-----");
    }

};