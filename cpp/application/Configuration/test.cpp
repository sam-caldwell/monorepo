/**
 * @name Configuration/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

//This will turn on debug logging and other related tooling...
#define ARBITRARY_KV_LIST_DEBUGGING

#include <stack>
#include <iostream>
#include "projects/Tester/Tester/main.h"


//

/**
 * Base tests
 */
#include "tests/TestConstants.cpp"
#include "tests/TestBasicConfiguration.cpp"
#include "tests/TestCount.cpp"

#include "tests/TestParseBool.cpp"
#include "tests/TestParseDouble.cpp"
#include "tests/TestParseFloat.cpp"
#include "tests/TestParseInt.cpp"
#include "tests/TestParseUint.cpp"

#include "tests/TestLoadData.cpp"

#include "tests/TestGetType.cpp"
#include "tests/TestGetBool.cpp"
#include "tests/TestGetDouble.cpp"
#include "tests/TestGetFloat.cpp"
#include "tests/TestGetInt.cpp"
#include "tests/TestGetString.cpp"
#include "tests/TestGetUint.cpp"

using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();

    /**
     * Base tests
     */
    t.add(new TestConstants("TestConstants"), RUN_TEST);
    t.add(new TestBasicConfiguration("TestBasicConfiguration"), RUN_TEST);
    t.add(new TestCount("TestCount"), RUN_TEST);
    /**
     * parser method tests
     */
    t.add(new TestParseBool("TestParseBool"), RUN_TEST);
    t.add(new TestParseDouble("TestParseDouble"), RUN_TEST);
    t.add(new TestParseFloat("TestParseFloat"), RUN_TEST);
    t.add(new TestParseInt("TestParseInt"), RUN_TEST);
    t.add(new TestParseUint("TestParseUint"), RUN_TEST);
    /**
     * data load tests
     */
    t.add(new TestLoadData("TestLoadData"), RUN_TEST);
    /**
     * data getter tests
     */
    t.add(new TestGetType("TestGetType"), RUN_TEST);
    t.add(new TestGetBool("TestGetBool"), RUN_TEST);
    t.add(new TestGetDouble("TestGetDouble"), RUN_TEST);
    t.add(new TestGetFloat("TestGetFloat"), RUN_TEST);
    t.add(new TestGetInt("TestGetInt"), RUN_TEST);
    t.add(new TestGetString("TestGetString"), RUN_TEST);
    t.add(new TestGetUint("TestGetUint"), RUN_TEST);
    /**
     *  END OF TESTS
     */
    return t.run();
}
