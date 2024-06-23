/**
 * @name Logger/tests/TestLoggerAddContextWithValues.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */


#include <iostream>
#include "../src/Logger.h"

class TestLoggerAddContextWithValues : public TestBase {

private:
    const Logging::Writer::Id writerNull = Logging::Writer::null;
    const Logging::Writer::Id writerStdout = Logging::Writer::stdout;
    const Logging::Writer::Id writerStderr = Logging::Writer::stderr;
    const Logging::Writer::Id writerFile = Logging::Writer::null; // disabling file writer
    const Logging::Writer::Id writerSyslog = Logging::Writer::null; // disabling file writer
    const Logging::Writer::Id writerHttps = Logging::Writer::null; // disabling file writer

public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestLoggerAddContextWithValues(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestLoggerAddContextWithValues() {}

    /**
     * @name test_instantiation
     * @brief instantiate the Application class
     *
     * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
     * @param writer - Log writer Id (e.g. stdout, stderr, file, syslog, https)
     * @param tags - a comma-delimited set of quoted strings used to map attributes
     *               of logs across our observability tooling.
     * @param destination - a String url in the form <protocol>:://<addr[:<port>]> used to
     *                      identify the log receiver (e.g. file://path/filename or syslog:://<ip:port>).
     * @return bool
     */
    bool test_instantiation(Logging::Severity sev, Logging::Writer::Id writer, string *tags, string *destination) {
        debug("test_instantiation()");
        if (tags) {
            if (destination) {
                debug("test_instantiation() with tags and destination");
                Logger log(sev, writer, *tags, *destination);
                return expect(log.contextCount() == 1, "Expect log(1) context count is 1");
            } else {
                debug("test_instantiation() with tags");
                Logger log(sev, writer, *tags);
                return expect(log.contextCount() == 1, "Expect log(2) context count is 1");
            }
        } else {
            debug("test_instantiation() with basics");
            Logger log(sev, writer);
            return expect(log.contextCount() == 1, "Expect log(3) context count is 1");
        }
    }

    /**
     * @name test_will_not_revert_root_context
     * @brief verify that given root context only popContext() will throw out_of_range exception.
     *
     * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
     * @param writer - Log writer Id (e.g. stdout, stderr, file, syslog, https)
     * @param tags - a comma-delimited set of quoted strings used to map attributes
     *               of logs across our observability tooling.
     * @param destination - a String url in the form <protocol>:://<addr[:<port>]> used to
     *                      identify the log receiver (e.g. file://path/filename or syslog:://<ip:port>).
     * @return bool
     */
    bool test_will_not_revert_root_context(Logging::Severity sev,
                                           Logging::Writer::Id writer,
                                           string *tags,
                                           string *destination) {
        debug("test_will_not_revert_root_context()");
        Logger log(sev, writer, (tags) ? *tags : "", (destination) ? *destination : "");
        try {
            log.popContext();
            return expect(false, "expect out_of_range exception when revert attempts to remove root context.");
        } catch (out_of_range &e) {
            return expect(true, "expect out_of_range exception when revert attempts to remove root context.") &&
                   expect(log.contextCount() == 1, "Expect context count is 1") &&
                   expect(string(e.what()).compare("logger ContextId 0 cannot be removed.") == 0,
                          "expect correct error message");
        }
    }

    /**
     * @name test_add_revert_context
     * @brief Verify we can add and revert a non-root context.
     *
     * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
     * @param writer - Log writer Id (e.g. stdout, stderr, file, syslog, https)
     * @param tags - a comma-delimited set of quoted strings used to map attributes
     *               of logs across our observability tooling.
     * @param destination - a String url in the form <protocol>:://<addr[:<port>]> used to
     *                      identify the log receiver (e.g. file://path/filename or syslog:://<ip:port>).
     * @return bool
     */
    bool test_add_revert_context(Logging::Severity sev, Logging::Writer::Id writer, string *tags, string *destination) {
        Logger log(sev, writer, (tags) ? *tags : "", (destination) ? *destination : "");
        try {
            return expect(log.newContext(), "expect new context can be created.") &&
                   expect(log.contextCount() == 2, "expect count is 2") &&
                   expect(log.popContext(), "expect new context can be destroyed.") &&
                   expect(log.contextCount() == 1, "expect count is 1");
        } catch (exception) {
            return false;
        }
    }

    /**
     * @name test_add_context_limit
     * @brief test that we get an out_of_range exception if we newContext() more than 256 contexts.
     *
     * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
     * @param writer - Log writer Id (e.g. stdout, stderr, file, syslog, https)
     * @param tags - a comma-delimited set of quoted strings used to map attributes
     *               of logs across our observability tooling.
     * @param destination - a String url in the form <protocol>:://<addr[:<port>]> used to
     *                      identify the log receiver (e.g. file://path/filename or syslog:://<ip:port>).
     * @return bool
     */
    bool test_add_context_limit(Logging::Severity sev, Logging::Writer::Id writer, string *tags, string *destination) {
        Logger log(sev, writer, (tags) ? *tags : "", (destination) ? *destination : "");
        for (unsigned count = 0; count < 3000; count++) {
            try {
                log.newContext();
                if (count >= 256)
                    throw runtime_error("newContext() expected exception where count>=256. Got:" + to_string(count));
            } catch (out_of_range &e) {
                return expect(true, "Expect out_of_range exception") &&
                       expect(count < 256, "Expect out_of_range exception if count>256. Actual:" + to_string(count));
            } catch (runtime_error &fe) {
                return expect(false, string(fe.what()));
            }
        }
        return expect(false, "Expect out_of_range exception when context count >= 256");
    }

    /**
     * @name test_context_counts
     * @brief Verify that we can create a number of contexts, revert them and not leak any context objects.
     *
     * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
     * @param writer - Log writer Id (e.g. stdout, stderr, file, syslog, https)
     * @param tags - a comma-delimited set of quoted strings used to map attributes
     *               of logs across our observability tooling.
     * @param destination - a String url in the form <protocol>:://<addr[:<port>]> used to
     *                      identify the log receiver (e.g. file://path/filename or syslog:://<ip:port>).
     * @return bool
     */
    bool test_context_counts(Logging::Severity sev, Logging::Writer::Id writer, string *tags, string *destination) {
        Logger log(sev, writer, (tags) ? *tags : "", (destination) ? *destination : "");
        for (unsigned char i = 0; i < 255; i++) log.newContext();
        for (unsigned char i = 0; i < 255; i++) log.popContext();
        try {
            log.popContext();
            return expect(false, "Expected out_of_range exception since we should have removed all contexts.");
        } catch (out_of_range &e) {
            return expect(true, "Expected out_of_range exception because our last revert would have removed root.");
        }
    }

    /**
     * @name scenario
     * @brief test with given conditions
     * 
     * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
     * @param writer - Log writer Id (e.g. stdout, stderr, file, syslog, https)
     * @param tags - a comma-delimited set of quoted strings used to map attributes
     *               of logs across our observability tooling.
     * @param destination - a String url in the form <protocol>:://<addr[:<port>]> used to
     *                      identify the log receiver (e.g. file://path/filename or syslog:://<ip:port>).
     * @return bool
     */
    bool scenario(Logging::Severity sev, Logging::Writer::Id writer, string *tags, string *destination) {
        debug("scenario starting");
        return expect(test_instantiation(sev, writer, tags, destination), "test_instantiation() ok") &&
               expect(test_will_not_revert_root_context(sev, writer, tags, destination),
                      "test_will_not_revert_root_context() ok") &&
               expect(test_add_revert_context(sev, writer, tags, destination), "test_add_revert_context() ok") &&
               expect(test_add_context_limit(sev, writer, tags, destination), "test_add_context_limit() ok") &&
               expect(test_context_counts(sev, writer, tags, destination), "test_context_counts() ok");
    }


    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        string tags = "\"tag:1\",\"tag:2\",\"tag:3\",\"tag:4\"";
        string testLogFile = "./artifacts/TestLoggerAddContextWithValues.log.txt";
        string syslogEndpoint = "syslog://127.0.0.1:514";
        string httpsEndpoint = "https://127.0.0.1:8514";

        debug(name + "::main()");
        return expect(scenario(Logging::Severity::emergency, writerNull, NULL, NULL), "scenario 1.0 ok") &&
               expect(scenario(Logging::Severity::alert, writerNull, NULL, NULL), "scenario 1.1 ok") &&
               expect(scenario(Logging::Severity::critical, writerNull, NULL, NULL), "scenario 1.2 ok") &&
               expect(scenario(Logging::Severity::error, writerNull, NULL, NULL), "scenario 1.3 ok") &&
               expect(scenario(Logging::Severity::warning, writerNull, NULL, NULL), "scenario 1.4 ok") &&
               expect(scenario(Logging::Severity::notice, writerNull, NULL, NULL), "scenario 1.5 ok") &&
               expect(scenario(Logging::Severity::info, writerNull, NULL, NULL), "scenario 1.6 ok") &&
               expect(scenario(Logging::Severity::debug, writerNull, NULL, NULL), "scenario 1.7 ok") &&

               expect(scenario(Logging::Severity::emergency, writerStdout, &tags, NULL), "scenario 2.0 ok") &&
               expect(scenario(Logging::Severity::alert, writerStdout, &tags, NULL), "scenario 2.1 ok") &&
               expect(scenario(Logging::Severity::critical, writerStdout, &tags, NULL), "scenario 2.2 ok") &&
               expect(scenario(Logging::Severity::error, writerStdout, &tags, NULL), "scenario 2.3 ok") &&
               expect(scenario(Logging::Severity::warning, writerStdout, &tags, NULL), "scenario 2.4 ok") &&
               expect(scenario(Logging::Severity::notice, writerStdout, &tags, NULL), "scenario 2.5 ok") &&
               expect(scenario(Logging::Severity::info, writerStdout, &tags, NULL), "scenario 2.6 ok") &&
               expect(scenario(Logging::Severity::debug, writerStdout, &tags, NULL), "scenario 2.7 ok") &&

               expect(scenario(Logging::Severity::emergency, writerStderr, &tags, NULL), "scenario 3.0 ok") &&
               expect(scenario(Logging::Severity::alert, writerStderr, &tags, NULL), "scenario 3.1 ok") &&
               expect(scenario(Logging::Severity::critical, writerStderr, &tags, NULL), "scenario 3.2 ok") &&
               expect(scenario(Logging::Severity::error, writerStderr, &tags, NULL), "scenario 3.3 ok") &&
               expect(scenario(Logging::Severity::warning, writerStderr, &tags, NULL), "scenario 3.4 ok") &&
               expect(scenario(Logging::Severity::notice, writerStderr, &tags, NULL), "scenario 3.5 ok") &&
               expect(scenario(Logging::Severity::info, writerStderr, &tags, NULL), "scenario 3.6 ok") &&
               expect(scenario(Logging::Severity::debug, writerStderr, &tags, NULL), "scenario 3.7 ok") &&

               expect(scenario(Logging::Severity::emergency, writerFile, &tags, &testLogFile), "scenario 4.0 ok") &&
               expect(scenario(Logging::Severity::alert, writerFile, &tags, &testLogFile), "scenario 4.1 ok") &&
               expect(scenario(Logging::Severity::critical, writerFile, &tags, &testLogFile), "scenario 4.2 ok") &&
               expect(scenario(Logging::Severity::error, writerFile, &tags, &testLogFile), "scenario 4.3 ok") &&
               expect(scenario(Logging::Severity::warning, writerFile, &tags, &testLogFile), "scenario 4.4 ok") &&
               expect(scenario(Logging::Severity::notice, writerFile, &tags, &testLogFile), "scenario 4.5 ok") &&
               expect(scenario(Logging::Severity::info, writerFile, &tags, &testLogFile), "scenario 4.6 ok") &&
               expect(scenario(Logging::Severity::debug, writerFile, &tags, &testLogFile), "scenario 4.7 ok") &&

               expect(scenario(Logging::Severity::emergency, writerSyslog, &tags, &syslogEndpoint),
                      "scenario 5.0 ok") &&
               expect(scenario(Logging::Severity::alert, writerSyslog, &tags, &syslogEndpoint), "scenario 5.1 ok") &&
               expect(scenario(Logging::Severity::critical, writerSyslog, &tags, &syslogEndpoint), "scenario 5.2 ok") &&
               expect(scenario(Logging::Severity::error, writerSyslog, &tags, &syslogEndpoint), "scenario 5.3 ok") &&
               expect(scenario(Logging::Severity::warning, writerSyslog, &tags, &syslogEndpoint), "scenario 5.4 ok") &&
               expect(scenario(Logging::Severity::notice, writerSyslog, &tags, &syslogEndpoint), "scenario 5.5 ok") &&
               expect(scenario(Logging::Severity::info, writerSyslog, &tags, &syslogEndpoint), "scenario 5.6 ok") &&
               expect(scenario(Logging::Severity::debug, writerSyslog, &tags, &syslogEndpoint), "scenario 5.7 ok") &&

               expect(scenario(Logging::Severity::emergency, writerHttps, &tags, &httpsEndpoint), "scenario 6.0 ok") &&
               expect(scenario(Logging::Severity::alert, writerHttps, &tags, &httpsEndpoint), "scenario 6.1 ok") &&
               expect(scenario(Logging::Severity::critical, writerHttps, &tags, &httpsEndpoint), "scenario 6.2 ok") &&
               expect(scenario(Logging::Severity::error, writerHttps, &tags, &httpsEndpoint), "scenario 6.3 ok") &&
               expect(scenario(Logging::Severity::warning, writerHttps, &tags, &httpsEndpoint), "scenario 6.4 ok") &&
               expect(scenario(Logging::Severity::notice, writerHttps, &tags, &httpsEndpoint), "scenario 6.5 ok") &&
               expect(scenario(Logging::Severity::info, writerHttps, &tags, &httpsEndpoint), "scenario 6.6 ok") &&
               expect(scenario(Logging::Severity::debug, writerHttps, &tags, &httpsEndpoint), "scenario 6.7 ok");
    }

};