/**
 * @name TesterFramework/TesterFramework.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef TESTER_TESTER_H
#define TESTER_TESTER_H

#include <cstdlib>
#include <filesystem>
#include <iostream>
#include <fstream>
#include <queue>
#include "../../common/AnsiColors/AnsiCodes.h"
#include "../TestBase/main.h"

#define SKIP_TEST true
#define RUN_TEST false

using namespace std;
using std::ofstream;

class Tester {
public:
    /**
     * Class constructor.
     *  - set state to zero.
     *  @param bool debug
     */
    Tester(int argc, char **argv);

    /**
     * Class destructor (noop)
     */
    ~Tester();

    /**
     * add a test object to the queue.
     * @param TestBase *t
     * @param bool skip (default: false
     * @return bool (pass/fail)
     */
    bool add(TestBase *t, bool skip);

    /**
     * run() method executes the tests registered in the queue.
     * @return bool (pass/fail)
     */
    int run();

    /**
     * enable debug_flag
     */
    void enable_debug();

    /**
     * count method returns the number of tests.
     * @return int
     */
    int count();

    /**
     * return number of tests passing.
     * @return int
     */
    int number_passed();

    /**
     * return number of tests failing.
     * @return int
     */
    int number_failed();

    /**
     * return number skipped.
     */
    int number_skipped();

private:
    ofstream log;
    bool debug_flag;
    string test_group_name;
    int test_count;
    int passed;
    int failed;
    int skipped;
    queue<TestBase *> tests;
};

#include "exitCodetoString.cpp"
#include "add.cpp"
#include "constructor.cpp"
#include "count.cpp"
#include "destructor.cpp"
#include "enable_debug.cpp"
#include "number_failed.cpp"
#include "number_passed.cpp"
#include "number_skipped.cpp"
#include "run.cpp"

#endif /* TESTER_TESTER_H */