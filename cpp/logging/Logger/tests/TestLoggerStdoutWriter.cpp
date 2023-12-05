/**
 * @name Logger/tests/TestLoggerStdoutWriter.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/common/types/ConfigStateMap.h"
#include "projects/application/Configuration/src/Configuration.h"
#include "projects/logging/Logger/tests/validateLine.cpp"
#include "projects/logging/Logger/src/Logger.h"
#include "projects/Tester/Mocks/MockStdout.h"


class TestLoggerStdoutWriter : public TestBase {
private:
    Logging::Writer::Id thisWriter;
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestLoggerStdoutWriter(string n) {
        name = n;
        thisWriter = Logging::Writer::stdout;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestLoggerStdoutWriter() {/*noop*/}

    /**
     * @name scenario
     * @brief Test and expect we can set our log level as needed.
     *
     * @param sev1 - Log severity of the log context.
     * @param sev2 - Log severity of the log message.
     * @params writeFlag - bool - expected return from Context::write()
     * @return bool
     */
    bool scenario(Logging::Severity sev1, Logging::Severity sev2, bool writeFlag = true) {
        string tags = "\"tag:1\",\"tag:2\",\"tag:3\",\"tag:4\"";
        string destination = "artifacts/test.log.txt";
        string msg = "test_message";
        Mock::Stdout mock;

        Context logContext(1, sev1, thisWriter, &tags, &destination);

        mock.enable();
        bool writeOutcome = logContext.write(sev2, &msg);
        mock.capture();
        string output = mock.last();

        if (writeFlag) {
            return expect(writeOutcome, "Expect write ok") &&
                   expect(validateLine(sev2, &tags, &output, &msg), "expect log is valid.  line: "+output);
        } else {
            return expect(true, "Expect no-write ok");
        }
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        Logging::Severity emergency = Logging::Severity::emergency;
        Logging::Severity alert = Logging::Severity::alert;
        Logging::Severity critical = Logging::Severity::critical;
        Logging::Severity error = Logging::Severity::error;
        Logging::Severity warning = Logging::Severity::warning;
        Logging::Severity notice = Logging::Severity::notice;
        Logging::Severity info = Logging::Severity::info;
        Logging::Severity debug = Logging::Severity::debug;

        return expect(scenario(emergency, emergency, true), "scenario 1.0 ok") &&
               expect(scenario(emergency, alert, false), "scenario 1.1 ok") &&
               expect(scenario(emergency, critical, false), "scenario 1.2 ok") &&
               expect(scenario(emergency, error, false), "scenario 1.3 ok") &&
               expect(scenario(emergency, warning, false), "scenario 1.4 ok") &&
               expect(scenario(emergency, notice, false), "scenario 1.5 ok") &&
               expect(scenario(emergency, info, false), "scenario 1.6 ok") &&
               expect(scenario(emergency, debug, false), "scenario 1.7 ok") &&

               expect(scenario(alert, emergency, true), "scenario 2.0 ok") &&
               expect(scenario(alert, alert, true), "scenario 2.1 ok") &&
               expect(scenario(alert, critical, false), "scenario 2.2 ok") &&
               expect(scenario(alert, error, false), "scenario 2.3 ok") &&
               expect(scenario(alert, warning, false), "scenario 2.4 ok") &&
               expect(scenario(alert, notice, false), "scenario 2.5 ok") &&
               expect(scenario(alert, info, false), "scenario 2.6 ok") &&
               expect(scenario(alert, debug, false), "scenario 2.7 ok") &&

               expect(scenario(critical, emergency, true), "scenario 3.0 ok") &&
               expect(scenario(critical, alert, true), "scenario 3.1 ok") &&
               expect(scenario(critical, critical, true), "scenario 3.2 ok") &&
               expect(scenario(critical, error, false), "scenario 3.3 ok") &&
               expect(scenario(critical, warning, false), "scenario 3.4 ok") &&
               expect(scenario(critical, notice, false), "scenario 3.5 ok") &&
               expect(scenario(critical, info, false), "scenario 3.6 ok") &&
               expect(scenario(critical, debug, false), "scenario 3.7 ok") &&

               expect(scenario(error, emergency, true), "scenario 4.0 ok") &&
               expect(scenario(error, alert, true), "scenario 4.1 ok") &&
               expect(scenario(error, critical, true), "scenario 4.2 ok") &&
               expect(scenario(error, error, true), "scenario 4.3 ok") &&
               expect(scenario(error, warning, false), "scenario 4.4 ok") &&
               expect(scenario(error, notice, false), "scenario 4.5 ok") &&
               expect(scenario(error, info, false), "scenario 4.6 ok") &&
               expect(scenario(error, debug, false), "scenario 4.7 ok") &&

               expect(scenario(warning, emergency, true), "scenario 5.0 ok") &&
               expect(scenario(warning, alert, true), "scenario 5.1 ok") &&
               expect(scenario(warning, critical, true), "scenario 5.2 ok") &&
               expect(scenario(warning, error, true), "scenario 5.3 ok") &&
               expect(scenario(warning, warning, true), "scenario 5.4 ok") &&
               expect(scenario(warning, notice, false), "scenario 5.5 ok") &&
               expect(scenario(warning, info, false), "scenario 5.6 ok") &&
               expect(scenario(warning, debug, false), "scenario 5.7 ok") &&

               expect(scenario(notice, emergency, true), "scenario 6.0 ok") &&
               expect(scenario(notice, alert, true), "scenario 6.1 ok") &&
               expect(scenario(notice, critical, true), "scenario 6.2 ok") &&
               expect(scenario(notice, error, true), "scenario 6.3 ok") &&
               expect(scenario(notice, warning, true), "scenario 6.4 ok") &&
               expect(scenario(notice, notice, true), "scenario 6.5 ok") &&
               expect(scenario(notice, info, false), "scenario 6.6 ok") &&
               expect(scenario(notice, debug, false), "scenario 6.7 ok") &&

               expect(scenario(info, emergency, true), "scenario 7.0 ok") &&
               expect(scenario(info, alert, true), "scenario 7.1 ok") &&
               expect(scenario(info, critical, true), "scenario 7.2 ok") &&
               expect(scenario(info, error, true), "scenario 7.3 ok") &&
               expect(scenario(info, warning, true), "scenario 7.4 ok") &&
               expect(scenario(info, notice, true), "scenario 7.5 ok") &&
               expect(scenario(info, info, true), "scenario 7.6 ok") &&
               expect(scenario(info, debug, false), "scenario 7.7 ok") &&

               expect(scenario(debug, emergency, true), "scenario 8.0 ok") &&
               expect(scenario(debug, alert, true), "scenario 8.1 ok") &&
               expect(scenario(debug, critical, true), "scenario 8.2 ok") &&
               expect(scenario(debug, error, true), "scenario 8.3 ok") &&
               expect(scenario(debug, warning, true), "scenario 8.4 ok") &&
               expect(scenario(debug, notice, true), "scenario 8.5 ok") &&
               expect(scenario(debug, info, true), "scenario 8.6 ok") &&
               expect(scenario(debug, debug, true), "scenario 8.7 ok");
    }
};