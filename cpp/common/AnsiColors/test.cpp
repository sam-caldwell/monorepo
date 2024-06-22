/**
 * @name AnsiColors/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../../networking/TcpCommon/src/PipeResponse/main.h"

/**
 * Base tests
 */
#include "tests/TestAnsiColors.cpp"


using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();

    /**
     * Base tests
     */
    t.add(new TestAnsiColors("TestAnsiColors"), RUN_TEST);

    /**
     *  END OF TESTS
     */
    return t.run();
}
