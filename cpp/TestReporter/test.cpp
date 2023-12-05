/**
 * @name TestReporter/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <stack>
#include <iostream>
#include "projects/Tester/Tester/main.h"
/**
 * Base tests
 */
#include "projects/TestReporter/tests/TestTestReporter.cpp"
#include <filesystem>

using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();
    /**
     * Base tests
     */
    t.add(new TestTestReporter("TestTestReporter"), RUN_TEST);
    /**
     *  END OF TESTS
     */
    return t.run();
}
