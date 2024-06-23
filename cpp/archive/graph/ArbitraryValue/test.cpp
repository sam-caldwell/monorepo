/**
 * @name ArbitraryValue/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <stack>
#include <iostream>
#include "../../../networking/TcpCommon/src/PipeResponse/main.h"




/**
 * Base tests
 */
#include "tests/TestArbitraryValue.cpp"
#include "tests/TestArbitraryValueType.cpp"
#include "tests/TestArbitraryValueTypeCheck.cpp"
/**
 * Bool to * tests
 */
#include "tests/TestArbitraryValueSetBoolGetBool.cpp"
#include "tests/TestArbitraryValueSetBoolGetDouble.cpp"
#include "tests/TestArbitraryValueSetBoolGetFloat.cpp"
#include "tests/TestArbitraryValueSetBoolGetInt.cpp"
#include "tests/TestArbitraryValueSetBoolGetUint.cpp"
#include "tests/TestArbitraryValueSetBoolGetString.cpp"
/**
 * Double to * tests
 */
#include "tests/TestArbitraryValueSetDoubleGetBool.cpp"
#include "tests/TestArbitraryValueSetDoubleGetDouble.cpp"
#include "tests/TestArbitraryValueSetDoubleGetFloat.cpp"
#include "tests/TestArbitraryValueSetDoubleGetInt.cpp"
#include "tests/TestArbitraryValueSetDoubleGetString.cpp"
#include "tests/TestArbitraryValueSetDoubleGetUint.cpp"
/**
 * Float to * tests
 */
#include "tests/TestArbitraryValueSetFloatGetBool.cpp"
#include "tests/TestArbitraryValueSetFloatGetDouble.cpp"
#include "tests/TestArbitraryValueSetFloatGetFloat.cpp"
#include "tests/TestArbitraryValueSetFloatGetInt.cpp"
#include "tests/TestArbitraryValueSetFloatGetString.cpp"
#include "tests/TestArbitraryValueSetFloatGetUint.cpp"
/**
 * Int to * tests
 */
#include "tests/TestArbitraryValueSetIntGetBool.cpp"
#include "tests/TestArbitraryValueSetIntGetDouble.cpp"
#include "tests/TestArbitraryValueSetIntGetFloat.cpp"
#include "tests/TestArbitraryValueSetIntGetInt.cpp"
#include "tests/TestArbitraryValueSetIntGetString.cpp"
#include "tests/TestArbitraryValueSetIntGetUint.cpp"
/**
 * String to * tests
 */
#include "tests/TestArbitraryValueSetStringGetBool.cpp"
#include "tests/TestArbitraryValueSetStringGetDouble.cpp"
#include "tests/TestArbitraryValueSetStringGetFloat.cpp"
#include "tests/TestArbitraryValueSetStringGetInt.cpp"
#include "tests/TestArbitraryValueSetStringGetString.cpp"
#include "tests/TestArbitraryValueSetStringGetUint.cpp"
/**
 * Uint to * tests
 */
#include "tests/TestArbitraryValueSetUintGetBool.cpp"
#include "tests/TestArbitraryValueSetUintGetDouble.cpp"
#include "tests/TestArbitraryValueSetUintGetFloat.cpp"
#include "tests/TestArbitraryValueSetUintGetInt.cpp"
#include "tests/TestArbitraryValueSetUintGetString.cpp"
#include "tests/TestArbitraryValueSetUintGetUint.cpp"


using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();
    /**
     * Base tests
     */
    t.add(new TestArbitraryValueType("TestArbitraryValueType"), RUN_TEST);
    t.add(new TestArbitraryValue("TestArbitraryValue (Class Constructor)"), RUN_TEST);
    t.add(new TestArbitraryValueTypeCheck("TestArbitraryValueTypeCheck"), RUN_TEST);
    /**
     * Bool to * tests
     */
    t.add(new TestArbitraryValueSetBoolGetBool("TestArbitraryValueSetBoolGetBool"), RUN_TEST);
    t.add(new TestArbitraryValueSetBoolGetDouble("TestArbitraryValueSetFloatGetDouble"), RUN_TEST);
    t.add(new TestArbitraryValueSetBoolGetFloat("TestArbitraryValueSetBoolGetFloat"), RUN_TEST);
    t.add(new TestArbitraryValueSetBoolGetInt("TestArbitraryValueSetBoolGetInt"), RUN_TEST);
    t.add(new TestArbitraryValueSetBoolGetString("TestArbitraryValueSetBoolGetString"), RUN_TEST);
    t.add(new TestArbitraryValueSetBoolGetUint("TestArbitraryValueSetBoolGetUint"), RUN_TEST);
    /**
     * Double to * tests
     */
    t.add(new TestArbitraryValueSetDoubleGetBool("TestArbitraryValueSetDoubleGetBool"), RUN_TEST);
    t.add(new TestArbitraryValueSetDoubleGetDouble("TestArbitraryValueSetDoubleGetDouble"), RUN_TEST);
    t.add(new TestArbitraryValueSetDoubleGetFloat("TestArbitraryValueSetDoubleGetFloat"), RUN_TEST);
    t.add(new TestArbitraryValueSetDoubleGetInt("TestArbitraryValueSetDoubleGetInt"), RUN_TEST);
    t.add(new TestArbitraryValueSetDoubleGetString("TestArbitraryValueSetDoubleGetString"), RUN_TEST);
    t.add(new TestArbitraryValueSetDoubleGetUint("TestArbitraryValueSetDoubleGetUint"), RUN_TEST);
    /**
     * Float to * tests
     */
    t.add(new TestArbitraryValueSetFloatGetBool("TestArbitraryValueSetFloatGetBool"), RUN_TEST);
    t.add(new TestArbitraryValueSetFloatGetDouble("TestArbitraryValueSetFloatGetDouble"), RUN_TEST);
    t.add(new TestArbitraryValueSetFloatGetFloat("TestArbitraryValueSetFloatGetFloat"), RUN_TEST);
    t.add(new TestArbitraryValueSetFloatGetInt("TestArbitraryValueSetFloatGetInt"), RUN_TEST);
    t.add(new TestArbitraryValueSetFloatGetString("TestArbitraryValueSetFloatGetString"), RUN_TEST);
    t.add(new TestArbitraryValueSetFloatGetUint("TestArbitraryValueSetFloatGetUint"), RUN_TEST);
    /**
     * Int to * tests
     */
    t.add(new TestArbitraryValueSetIntGetBool("TestArbitraryValueSetIntGetBool"), RUN_TEST);
    t.add(new TestArbitraryValueSetIntGetDouble("TestArbitraryValueSetIntGetDouble"), RUN_TEST);
    t.add(new TestArbitraryValueSetIntGetFloat("TestArbitraryValueSetIntGetFloat"), RUN_TEST);
    t.add(new TestArbitraryValueSetIntGetInt("TestArbitraryValueSetIntGetInt"), RUN_TEST);
    t.add(new TestArbitraryValueSetIntGetString("TestArbitraryValueSetIntGetString"), RUN_TEST);
    t.add(new TestArbitraryValueSetIntGetUint("TestArbitraryValueSetIntGetUint"), RUN_TEST);
    /**
     * String to * tests
     */
    t.add(new TestArbitraryValueSetStringGetBool("TestArbitraryValueSetStringGetBool"), RUN_TEST);
    t.add(new TestArbitraryValueSetStringGetDouble("TestArbitraryValueSetStringGetDouble"), RUN_TEST);
    t.add(new TestArbitraryValueSetStringGetFloat("TestArbitraryValueSetStringGetFloat"), RUN_TEST);
    t.add(new TestArbitraryValueSetStringGetInt("TestArbitraryValueSetStringGetInt"), RUN_TEST);
    t.add(new TestArbitraryValueSetStringGetString("TestArbitraryValueSetStringGetString"), RUN_TEST);
    t.add(new TestArbitraryValueSetStringGetUint("TestArbitraryValueSetStringGetUint"), RUN_TEST);
    /**
     * Uint to * tests
     */
    t.add(new TestArbitraryValueSetUintGetBool("TestArbitraryValueSetUintGetBool"), RUN_TEST);
    t.add(new TestArbitraryValueSetUintGetDouble("TestArbitraryValueSetUintGetDouble"), RUN_TEST);
    t.add(new TestArbitraryValueSetUintGetFloat("TestArbitraryValueSetUintGetFloat"), RUN_TEST);
    t.add(new TestArbitraryValueSetUintGetInt("TestArbitraryValueSetUintGetInt"), RUN_TEST);
    t.add(new TestArbitraryValueSetUintGetString("TestArbitraryValueSetUintGetString"), RUN_TEST);
    t.add(new TestArbitraryValueSetUintGetUint("TestArbitraryValueSetUintGetUint"), RUN_TEST);
    /**
     *  END OF TESTS
     */
    return t.run();
}
