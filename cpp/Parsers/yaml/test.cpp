/**
 * @name Parsers/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#define TESTING_ONLY

#include <iostream>
#include "../../../Tester/Tester/main.h"

/**
 * Base tests
 */
#include "tests/TestParserInitialState.cpp"
#include "tests/TestParserStripComment.cpp"
#include "tests/TestParserExtractTagOnly.cpp"
#include "tests/TestParserTokenType.cpp"
#include "tests/TestParserGetIndentation.cpp"
#include "tests/TestParserFetchNextLine.cpp"
#include "tests/TestParserFetchNextLinePreserve.cpp"
#include "tests/TestParserValidate.cpp"
#include "tests/TestParserGetToken.cpp"
/**
 * Parsers Tests
 */
#include "tests/TestParserRouteAny.cpp"
#include "tests/TestParserRouteHeader.cpp"
#include "tests/TestParserRouteKeyValue.cpp"
#include "tests/TestParserRouteMultiLineTag.cpp"
#include "tests/TestParserRouteSequenceItem.cpp"
#include "tests/TestParserRouteSequenceKeyValue.cpp"
#include "tests/TestParserRouteTagMap.cpp"
/**
 * Token Iterator
 */
#include "tests/TestParserTokenCount.cpp"
#include "tests/TestParserTokenIterator.cpp"
/**
 * Parsers end-to-end
 */
#include "tests/TestParserYaml.cpp"

using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();

    /**
     * Node Tests
     */
    t.add(new TestParserInitialState("TestParserInitialState"), RUN_TEST);
    t.add(new TestParserStripComment("TestParserStripComment"), RUN_TEST);
    t.add(new TestParserExtractTagOnly("TestParserExtractTagOnly"), RUN_TEST);
    t.add(new TestParserTokenType("TestParserTokenType"), RUN_TEST);
    t.add(new TestParserGetIndentation("TestParserGetIndentation"), RUN_TEST);
    t.add(new TestParserFetchNextLine("TestParserFetchNextLine"), RUN_TEST);
    t.add(new TestParserFetchNextLinePreserve("TestParserFetchNextLinePreserve"), RUN_TEST);
    t.add(new TestParserValidate("TestParserValidate"), RUN_TEST);
    t.add(new TestParserGetToken("TestParserGetToken"), RUN_TEST);
    /**
     * Parsers Tests
     */
    t.add(new TestParserRouteAny("TestParserRouteAny"), RUN_TEST);
    t.add(new TestParserRouteHeader("TestParserRouteHeader"), RUN_TEST);
    t.add(new TestParserRouteKeyValue("TestParserRouteKeyValue"), RUN_TEST);
    t.add(new TestParserRouteTagMap("TestParserRouteTagMap"), RUN_TEST);
    t.add(new TestParserRouteSequenceItem("TestParserRouteSequenceItem"), RUN_TEST);
    t.add(new TestParserRouteSequenceKeyValue("TestParserRouteSequenceKeyValue"), RUN_TEST);
    t.add(new TestParserRouteMultiLineTag("TestParserRouteMultiLineTag"), RUN_TEST);
    /**
     * Token Iterator
     */
    t.add(new TestParserTokenCount("TestParserTokenCount"), RUN_TEST);
    t.add(new TestParserTokenIterator("TestParserTokenIterator"), RUN_TEST);
    /**
     * End-to-End ParserTest
     */
    t.add(new TestParserYaml("TestParserYaml"), RUN_TEST);


    /**
     *  END OF TESTS
     */
    return t.run();
}
