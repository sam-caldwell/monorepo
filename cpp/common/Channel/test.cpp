/**
 * @name Channel/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */

//This will turn on debug logging and other related tooling...
#define CHANNEL_DEBUGGING

#include <stack>
#include <iostream>
#include "../../networking/TcpCommon/src/PipeResponse/main.h"
/**
 * Base tests
 */
#include "tests/TestBasics.h"
#include "tests/TestConcurrentUse.h"
#include "tests/TestWriteBlocked.h"
#include "tests/TestClosedConnectionException.h"

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
