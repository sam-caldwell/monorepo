/**
 * @name Logger/tests/TestLoggerAddContextWithDefaults.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */


#include <iostream>
#include "../src/Logger.h"

class TestLoggerAddContextWithDefaults : public TestBase {
private:
    Logger *log;
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestLoggerAddContextWithDefaults(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestLoggerAddContextWithDefaults() {
        if(log) delete log;
    }

    /**
     * @name test_instantiation
     * @brief instantiate the Application class
     *
     * @return bool
     */
    bool test_instantiation() {
        log = new Logger;
        debug("log instantiated");
        return expect(log, "expect log not null") &&
               expect(log->contextCount() == 1, "Expect context count is 1");
    }

    /**
     * @name test_will_not_revert_root_context
     * @brief verify that given root context only popContext() will throw out_of_range exception.
     *
     * @return bool
     */
    bool test_will_not_revert_root_context() {
        try {
            debug("popContext()...");
            log->popContext();
            debug("context popped");
            return expect(false, "expect out_of_range exception when revert attempts to remove root context.");
        } catch (out_of_range &e) {
            return expect(true, "expect out_of_range exception when revert attempts to remove root context.") &&
                   expect(log->contextCount() == 1, "Expect context count is 1") &&
                   expect(string(e.what()).compare("logger ContextId 0 cannot be removed.") == 0,
                          "expect correct error message");
        }
    }

    /**
     * @name test_add_revert_context
     * @brief Verify we can add and revert a non-root context.
     *
     * @return bool
     */
    bool test_add_revert_context() {
        try {
            return expect(log->newContext(), "expect new context can be created.") &&
                   expect(log->contextCount()==2, "expect count is 2") &&
                   expect(log->popContext(), "expect new context can be destroyed.") &&
                   expect(log->contextCount()==1, "expect count is 1");
        } catch (exception) {
            return false;
        }
    }

    /**
     * @name test_add_context_limit
     * @brief test that we get an out_of_range exception if we newContext() more than 256 contexts.
     *
     * @return bool
     */
    bool test_add_context_limit() {
        debug("test_add_context_limit() start");
        unsigned count = 0;
        while (count < 3000) {
            try {
                count++;
                log->newContext();
                if (count >= 256)
                    throw runtime_error("newContext() expected exception where count>=256. Got:" + to_string(count));
            } catch (out_of_range &e) {
                return expect(true, "Expect out_of_range exception") &&
                       expect(count >= 256, "Expect out_of_range exception if more than 256 contexts are created.");
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
     * @return bool
     */
    bool test_context_counts() {
        Logger thisLog;
        for (unsigned char i = 0; i < 255; i++) thisLog.newContext();
        for (unsigned char i = 0; i < 255; i++) thisLog.popContext();
        try {
            thisLog.popContext();
            return expect(false, "Expected out_of_range exception since we should have removed all contexts.");
        } catch (out_of_range &e) {
            return expect(true, "Expected out_of_range exception because our last revert would have removed root.");
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
        return expect(test_instantiation(), "test_instantiation() ok") &&
               expect(test_will_not_revert_root_context(), "test_will_not_revert_root_context() ok") &&
               expect(test_add_revert_context(), "test_add_revert_context() ok") &&
               expect(test_add_context_limit(), "test_add_context_limit() ok") &&
               expect(test_context_counts(), "test_context_counts() ok");
    }

};