/** ArbitraryKvLinkedList/test.cpp
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
#include "tests/TestDetectCircularReferences.cpp"
#include "tests/TestGreaterThan.cpp"
#include "tests/TestLessThan.cpp"
#include "tests/TestIsAdjacent.cpp"
#include "tests/TestResetNodeLeft.cpp"
#include "tests/TestResetNodeRight.cpp"
#include "tests/TestSwapPointers.cpp"
#include "tests/TestSwapNodes.cpp"
#include "tests/TestLinkListUnique.cpp"
#include "tests/TestLinkListInsert.cpp"
#include "tests/TestLinkListResetLeft.cpp"
#include "tests/TestLinkListResetRight.cpp"
#include "tests/TestLinkListExists.cpp"
#include "tests/TestLinkListCountMatches.cpp"
#include "tests/TestLinkListSearch.cpp"
#include "tests/TestLinkListGetValue.cpp"
#include "tests/TestLinkListDumpKeys.cpp"
#include "tests/TestLinkListRemoveKey.cpp"
#include "tests/TestLinkListRemoveDuplicates.cpp"
#include "tests/TestLinkListInsertMoveLeftOnInsert.cpp"
#include "tests/TestLinkListSwapSeparateNodes.cpp"
#include "tests/TestLinkListSwapAdjacent.cpp"
#include "tests/TestLinkListMoveLeftRight.cpp"
#include "tests/TestLinkListSortNoop.cpp"
#include "tests/TestLinkListSort.cpp"

using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();
    /**
     * Base tests
     */
    t.add(new TestDetectCircularReferences("TestDetectCircularReferences"), RUN_TEST);
    t.add(new TestGreaterThan("TestGreaterThan"), RUN_TEST);
    t.add(new TestLessThan("TestLessThan"), RUN_TEST);
    t.add(new TestIsAdjacent("TestIsAdjacent"), RUN_TEST);
    t.add(new TestResetNodeLeft("TestResetNodeLeft"), RUN_TEST);
    t.add(new TestResetNodeRight("TestResetNodeRight"), RUN_TEST);
    t.add(new TestSwapPointers("TestSwapPointers"), RUN_TEST);
    t.add(new TestSwapNodes("TestSwapNodes"), RUN_TEST);
    t.add(new TestLinkListUnique("TestLinkListUnique"), RUN_TEST);
    t.add(new TestLinkListInsert("TestLinkListInsert"), RUN_TEST);
    t.add(new TestLinkListResetLeft("TestLinkListResetLeft"), RUN_TEST);
    t.add(new TestLinkListResetRight("TestLinkListResetRight"), RUN_TEST);
    t.add(new TestLinkListExists("TestLinkListExists"), RUN_TEST);
    t.add(new TestLinkListCountMatches("TestLinkListCountMatches"), RUN_TEST);
    t.add(new TestLinkListGetValue("TestLinkListGetValue"), RUN_TEST);
    t.add(new TestLinkListSearch("TestLinkListSearch"), RUN_TEST);
    t.add(new TestLinkListDumpKeys("TestLinkListDumpKeys"), RUN_TEST);
    t.add(new TestLinkListRemoveKey("TestLinkListRemoveKey"), RUN_TEST);
    t.add(new TestLinkListRemoveDuplicates("TestLinkListRemoveDuplicates"), RUN_TEST);
    t.add(new TestLinkListInsertMoveLeftOnInsert("TestLinkListInsertMoveLeftOnInsert"), RUN_TEST);
    t.add(new TestLinkListSwapSeparateNodes("TestLinkListSwapSeparateNodes"), RUN_TEST);
    t.add(new TestLinkListSwapAdjacent("TestLinkListSwapAdjacent"), RUN_TEST);
    t.add(new TestLinkListMoveLeftRight("TestLinkListMoveLeftRight"), RUN_TEST);
    t.add(new TestLinkListSortNoop("TestLinkListSortNoop"), RUN_TEST);
    t.add(new TestLinkListSort("TestLinkListSort"), RUN_TEST);
    /**
     *  END OF TESTS
     */
    return t.run();
}
