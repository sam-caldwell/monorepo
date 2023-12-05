/**
 * @name Signals/tests/TestBasics.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/types/ConfigStateMap.h"
#include "projects/Signals/src/Signals.h"
#include "projects/Tester/Mocks/MockStdout.h"

/**
 * @namespace SignalTest
 * Create some test signal handlers.
 */
namespace SignalTest {
    void mockSignal(int i) { cout << i << endl; }
}

class TestBasics : public TestBase {
private:
    SignalHandler::Table signals;
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n - string
     */
    TestBasics(string n) { name = n; }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestBasics() {/*noop*/}

    /**
     * @name test_signal_map
     * @brief test the ability to generate a signal map
     *
     * @return bool
     */
    bool test_signal_map() {
        signals = {
                {SIGHUP,    &SignalTest::mockSignal},
                {SIGINT,    &SignalTest::mockSignal},
                {SIGQUIT,   &SignalTest::mockSignal},
                {SIGILL,    &SignalTest::mockSignal},
                {SIGTRAP,   &SignalTest::mockSignal},
                {SIGABRT,   &SignalTest::mockSignal},
                //{SIGBUS,    &SignalTest::mockSignal},
                {SIGFPE,    &SignalTest::mockSignal},
                //{SIGKILL, &SignalTest::mockSignal}, // - Kill cannot be captured...it kills the test program.
                //{SIGUSR1,   &SignalTest::mockSignal},
                {SIGSEGV,   &SignalTest::mockSignal},
                //{SIGUSR2,   &SignalTest::mockSignal},
                {SIGPIPE,   &SignalTest::mockSignal},
                {SIGALRM,   &SignalTest::mockSignal},
                {SIGTERM,   &SignalTest::mockSignal},
                //{SIGSTKFLT, &SignalTest::mockSignal}, // - undefined - Stack fault on coprocessor (unused)
                //{SIGCHLD,   &SignalTest::mockSignal},
                //{SIGCONT,   &SignalTest::mockSignal},
                //{SIGSTOP,   &SignalTest::mockSignal},
                //{SIGTSTP,   &SignalTest::mockSignal},
                {SIGTTIN,   &SignalTest::mockSignal},
                {SIGTTOU,   &SignalTest::mockSignal},
                //{SIGURG,    &SignalTest::mockSignal},
                {SIGXCPU,   &SignalTest::mockSignal},
                {SIGXFSZ,   &SignalTest::mockSignal},
                {SIGVTALRM, &SignalTest::mockSignal},
                {SIGPROF,   &SignalTest::mockSignal},
                {SIGWINCH,  &SignalTest::mockSignal},
                //{SIGIO,     &SignalTest::mockSignal},
                //{SIGPWR, &SignalTest::mockSignal}, // - undefined - Power failure (System V)
                //{SIGSYS,    &SignalTest::mockSignal}
        };
        return expect(signals.size() == 18, "signal table size ok (got "+to_string(signals.size())+")");
    }

    /**
     * @name test_signal_handler_init
     * @brief Test signal handler initializer
     *
     * @return bool
     */
    bool test_signal_handler_init() {
        SignalHandler::init(&signals);
        return true;
    }

    /**
     * @name test_signal
     * @brief Test the signal execution
     *
     * @param signal - int
     * @param expected - string
     * @return bool
     */
    bool test_signal(int signal, int expected) {
        Mock::Stdout mock;
        mock.enable();
        raise(signal);
        mock.capture();
        string line = mock.last();
        return expect(line.compare(to_string(expected)) == 0, "signal execution ok");
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_signal_map(), "test_signal_map() ok") &&
               expect(test_signal_handler_init(), "test_signal_handle_init() ok") &&
               expect(test_signal(SIGHUP, 1), "test_signal(SIGHUP) ok") &&
               expect(test_signal(SIGINT, 2), "test_signal(SIGINT) ok") &&
               expect(test_signal(SIGQUIT, 3), "test_signal(SIGQUIT) ok") &&
               expect(test_signal(SIGILL, 4), "test_signal(SIGILL) ok") &&
               expect(test_signal(SIGTRAP, 5), "test_signal(SIGTRAP) ok") &&
               expect(test_signal(SIGABRT, 6), "test_signal(SIGABRT) ok") &&
               expect(test_signal(SIGFPE, 8), "test_signal(SIGFPE) ok") &&
               expect(test_signal(SIGSEGV, 11), "test_signal(SIGSEGV) ok") &&
               expect(test_signal(SIGPIPE, 13), "test_signal(SIGPIPE) ok") &&
               expect(test_signal(SIGALRM, 14), "test_signal(SIGALRM) ok") &&
               expect(test_signal(SIGTERM, 15), "test_signal(SIGTERM) ok") &&
               expect(test_signal(SIGTTIN, 21), "test_signal(SIGTTIN) ok") &&
               expect(test_signal(SIGTTOU, 22), "test_signal(SIGTTOU) ok") &&
               expect(test_signal(SIGXCPU, 24), "test_signal(SIGXCPU) ok") &&
               expect(test_signal(SIGXFSZ, 25), "test_signal(SIGXFSZ) ok") &&
               expect(test_signal(SIGVTALRM, 26), "test_signal(SIGVTALRM) ok") &&
               expect(test_signal(SIGPROF, 27), "test_signal(SIGPROF) ok") &&
               expect(test_signal(SIGWINCH, 28), "test_signal(SIGWINCH) ok");
    }

};