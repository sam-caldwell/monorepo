/**
 * @name networking/TcpClient/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/Tester/Tester/main.h"

/**
 * TcpCommon tests
 */
#include "../TcpCommon/tests/TestBasics.cpp"
#include "../TcpCommon/tests/TestFileDescriptor.cpp"
#include "../TcpCommon/tests/TestWaitFor.cpp"
#include "../TcpCommon/tests/TestPipeResponse.cpp"
#include "../TcpCommon/tests/TestClient.cpp"
/**
 * TcpClient tests
 */
#include "tests/TestTcpClientBasics.cpp"


using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();

    /**
     * TcpCommon tests
     */
    t.add(new TestBasics("TcpCommon/TestBasics"), RUN_TEST);
    t.add(new TestFileDescriptor("TcpCommon/TestFileDescriptor"), RUN_TEST);
    t.add(new TestWaitFor("TcpCommon/TestWaitFor"), RUN_TEST);
    t.add(new TestPipeResponse("TcpCommon/TestPipeResponse"), RUN_TEST);

    /**
     * TcpClient tests
     */
    t.add(new TestTcpClientBasics("TestTcpClientBasics"),RUN_TEST);

    /**
     *  END OF TESTS
     */
    return t.run();
}
