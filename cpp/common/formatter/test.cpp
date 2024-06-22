/**
 * @name formatter/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */

#include <stack>
#include <iostream>
#include "../../networking/TcpCommon/src/PipeResponse/main.h"
/**
 * Base tests
 */
#include "tests/TestIntToString.cpp"

#include "tests/TestToUpper.cpp"
#include "tests/TestToLower.cpp"

#include "tests/TestTrim.cpp"
#include "tests/TestLeftTrim.cpp"
#include "tests/TestRightTrim.cpp"

#include "tests/TestRightTrimChar.cpp"

#include "tests/TestSetToString.cpp"
#include "tests/TestStringToSet.cpp"

using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();

    /**
     * Base tests
     */
    t.add(new TestIntToString("TestIntToString"), RUN_TEST);
    t.add(new TestToUpper("TestToUpper"), RUN_TEST);
    t.add(new TestToLower("TestToLower"), RUN_TEST);
    t.add(new TestTrim("TestTrim"), RUN_TEST);
    t.add(new TestLeftTrim("TestLeftTrim"), RUN_TEST);
    t.add(new TestRightTrim("TestRightTrim"), RUN_TEST);
    t.add(new TestRightTrimChar("TestRightTrimChar"), RUN_TEST);
    t.add(new TestSetToString("TestSetToString"), RUN_TEST);
    t.add(new TestStringToSet("TestStringToSet"), RUN_TEST);

    /**
     *  END OF TESTS
     */
    return t.run();
}
