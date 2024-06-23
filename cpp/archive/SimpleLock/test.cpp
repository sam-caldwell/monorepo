/**
 * @name SimpleLock/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <stack>
#include <iostream>
#include "../networking/TcpCommon/src/PipeResponse/main.h"
/**
 * Base tests
 */

#include "tests/TestOnOff.cpp"
#include "tests/TestDoubleLock.cpp"
#include "tests/TestLockTimeout.cpp"

using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();

    /**
     * Base tests
     */
    t.add(new TestOnOff("TestOnOff"), RUN_TEST);
    t.add(new TestDoubleLock("TestDoubleLock"), RUN_TEST);
    t.add(new TestLockTimeout("TestLockTimeout"), RUN_TEST);
    /**
     *  END OF TESTS
     */
    return t.run();
}
