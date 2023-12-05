/**
 * @name EnvironmentVariableParser/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */

//This will turn on debug logging and other related tooling...
#define ARBITRARY_KV_LIST_DEBUGGING

#include <stack>
#include <iostream>
#include "projects/Tester/Tester/main.h"
/**
 * Base tests
 */
#include "tests/TestEvParserBasic.cpp"
#include "tests/TestEvParserDataLoad.cpp"
#include "tests/TestVerifyData.cpp"

using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();

    /**
     * Base tests
     */
    t.add(new TestEvParserBasic("TestEvParserBasic"), RUN_TEST);
    t.add(new TestEvParserDataLoad("TestEvParserDataLoad"), RUN_TEST);
    t.add(new TestVerifyData("TestVerifyData"), SKIP_TEST);
    /**
     *  END OF TESTS
     */
    return t.run();
}
