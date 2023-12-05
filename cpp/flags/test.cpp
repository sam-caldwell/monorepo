/**
 * @name flags/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/Tester/Tester/main.h"

/**
 * Base tests
 */
#include "tests/TestArch.cpp"
#include "tests/TestBoolean.cpp"
#include "tests/TestOpsys.cpp"


using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();

    /**
     * Base tests
     */
    t.add(new TestArch("TestArch"), RUN_TEST);
    t.add(new TestBoolean("TestBoolean"), RUN_TEST);
    t.add(new TestOpsys("TestOpsys"), RUN_TEST);

    /**
     *  END OF TESTS
     */
    return t.run();
}
