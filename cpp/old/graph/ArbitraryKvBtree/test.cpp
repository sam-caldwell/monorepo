/**
 * @name ArbitraryKvBtree/test.cpp
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
#include "tests/CreateAndDestroy.cpp"
#include "tests/TestBasicInsertUnique.cpp"
#include "tests/TestBasicInsertWithDuplicates.cpp"
#include "tests/TestInitialState.cpp"
#include "tests/TestInsertExpectNoDuplicates.cpp"
#include "tests/TestRemoveBasic.cpp"
#include "tests/TestCountMatches.cpp"
#include "tests/TestDuplicates.cpp"
#include "tests/TestFind.cpp"
#include "tests/TestDataBool.cpp"
#include "tests/TestDataDouble.cpp"
#include "tests/TestDataFloat.cpp"
#include "tests/TestDataInt.cpp"
#include "tests/TestDataString.cpp"
#include "tests/TestDataUint.cpp"

using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();

    /**
     * Base tests
     */
    t.add(new TestInitialState("TestInitialState"), RUN_TEST);
    t.add(new TestCreateAndDestroy("TestCreateAndDestroy"), RUN_TEST);
    t.add(new TestCountMatches("TestCountMatches"), RUN_TEST);
    t.add(new TestDuplicates("TestDuplicates"), RUN_TEST);
    t.add(new TestFind("TestFind"), RUN_TEST);

    /**
     * Insert scenarios
     */
    t.add(new TestBasicInsertUnique("TestBasicInsertUnique"), RUN_TEST);
    t.add(new TestBasicInsertWithDuplicates("TestBasicInsertWithDuplicates"), RUN_TEST);
    t.add(new TestInsertExpectNoDuplicates("TestInsertExpectNoDuplicates"), RUN_TEST);

    /**
     * Removal scenarios
     */
    t.add(new TestRemoveBasic("TestRemoveBasic"), RUN_TEST);

    /**
     * DataGetter/Setter scenarios
     */
    t.add(new TestDataBool("TestDataBool"), RUN_TEST);
    t.add(new TestDataDouble("TestDataDouble"), RUN_TEST);
    t.add(new TestDataFloat("TestDataFloat"), RUN_TEST);
    t.add(new TestDataInt("TestDataInt"), RUN_TEST);
    t.add(new TestDataString("TestDataString"), RUN_TEST);
    t.add(new TestDataUint("TestDataUint"), RUN_TEST);

    /**
     *  END OF TESTS
     */
    return t.run();
}
