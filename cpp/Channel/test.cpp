/**
 * @name Channel/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */

//This will turn on debug logging and other related tooling...
#define CHANNEL_DEBUGGING

#include <stack>
#include <iostream>
#include "projects/Tester/Tester/main.h"
/**
 * Base tests
 */
#include "tests/TestBasics.cpp"
#include "tests/TestConcurrentUse.cpp"
#include "tests/TestWriteBlocked.cpp"
#include "tests/TestClosedConnectionException.cpp"


using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();
    /**
     * Base tests
     */
    t.add(new TestBasics("TestBasics"), RUN_TEST);
    t.add(new TestWriteBlocked("TestWriteBlocked"), RUN_TEST);
    t.add(new TestClosedConnectionException("TestClosedConnectionException"), RUN_TEST);
    t.add(new TestConcurrentUse("TestConcurrentUse"), RUN_TEST);

    /**
     *  END OF TESTS
     */
    return t.run();
}
