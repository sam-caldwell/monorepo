/**
 * @name TesterFramework/TestBase.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef TEST_BASE_H
#define TEST_BASE_H

#include <string>
#include <iostream>
#include "../../common/AnsiColors/AnsiCodes.h"

using namespace std;

class TestBase {
protected:
    bool debug_flag;
    string name;
public:
    /**
     * Default Class Constructor.
     */
    TestBase();

    /**
     * Class Constructor
     *   - Add test name to internal state.
     * @param string n
     */
    TestBase(string n);

    /**
     * Class Destructor (noop)
     */
    ~TestBase();

    /**
     * main() method is the test payload
     * we will overload this method when we implement a test.
     * @return bool (pass/fail)
     */
    virtual bool main();

    /**
     * Expect() a condition to be true or show the message and return false.
     * @param bool condition
     * @param string message
     * @return bool (pass/fail)
     */
    bool expect(bool condition, string message);

    /**
     * Set debug flag.
     * @return operational outcome (success/fail)
     */
    bool set_debug();

    /**
     * Debug message printing for tests
     * @param msg
     * @return operational outcome (success/fail)
     */
    bool debug(string msg);

    /**
     * Print a red error message for tests.
     * @param msg
     * @return operation outcome (success/fail)
     */
     bool error(string msg);

    /**
     * return the test name
     * @return string
     */
    string get_name();
};

#include "TestBase.cpp"

#endif // TEST_BASE_H