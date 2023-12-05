/**
 * @name networking/TcpCommon/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/Tester/Tester/main.h"

/**
 * Base tests
 */
#include "tests/TestBasics.cpp"
#include "tests/TestFileDescriptor.cpp"
#include "tests/TestWaitFor.cpp"
#include "tests/TestPipeResponse.cpp"
#include "tests/TestClient.cpp"


using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();

    /**
     * Base tests
     */
    t.add(new TestBasics("TestBasics"), RUN_TEST);
    t.add(new TestFileDescriptor("TestFileDescriptor"), RUN_TEST);
    t.add(new TestWaitFor("TestWaitFor"), RUN_TEST);
    t.add(new TestPipeResponse("TestPipeResponse"), RUN_TEST);
    t.add(new TestClient("TestClient"), RUN_TEST);

    /**
     *  END OF TESTS
     */
    return t.run();
}
