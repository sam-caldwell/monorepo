/**
 * @name SimpleLock/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <stack>
#include <iostream>
#include "projects/Tester/Tester/main.h"
/**
 * Base tests
 */

#include "projects/common/SimpleLock/tests/TestOnOff.cpp"
#include "projects/common/SimpleLock/tests/TestDoubleLock.cpp"
#include "projects/common/SimpleLock/tests/TestLockTimeout.cpp"

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
