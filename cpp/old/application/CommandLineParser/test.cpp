/**
 * @name CommandLineParser/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */

//This will turn on debug logging and other related tooling...
#define ARBITRARY_KV_LIST_DEBUGGING

#include <stack>
#include <iostream>
#include "../../../networking/TcpCommon/src/PipeResponse/main.h"
/**
 * Base tests
 */
#include "tests/TestBasic.cpp"
#include "tests/TestParse.cpp"

using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();

    /**
     * Base tests
     */
    t.add(new TestBasic("TestBasic"), RUN_TEST);
    t.add(new TestParse("TestParse"), RUN_TEST);
    /**
     *  END OF TESTS
     */
    return t.run();
}
