/**
 * @name ArbitraryKeyValue/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <stack>
#include <iostream>
#include "projects/Tester/Tester/main.h"
/**
 * Base tests
 */
#include "tests/TestArbitraryKeyValue.cpp"
#include "tests/TestArbitraryKeyValueBool.cpp"
#include "tests/TestArbitraryKeyValueDouble.cpp"
#include "tests/TestArbitraryKeyValueFloat.cpp"
#include "tests/TestArbitraryKeyValueInt.cpp"
#include "tests/TestArbitraryKeyValueString.cpp"
#include "tests/TestArbitraryKeyValueUint.cpp"

using namespace std;

/**
 * @brief Test coordinator function.
 * @author Sam Caldwell <mail@samcaldwell.net>
 * @return int return code
 */
int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    //t.enable_debug();
    /**
     * Base tests
     */
    t.add(new TestArbitraryKeyValue("TestArbitraryKeyValue"), RUN_TEST);
    t.add(new TestArbitraryKeyValueBool("TestArbitraryKeyValueBool"), RUN_TEST);
    t.add(new TestArbitraryKeyValueDouble("TestArbitraryKeyValueDouble"), RUN_TEST);
    t.add(new TestArbitraryKeyValueFloat("TestArbitraryKeyValueFloat"), RUN_TEST);
    t.add(new TestArbitraryKeyValueInt("TestArbitraryKeyValueInt"), RUN_TEST);
    t.add(new TestArbitraryKeyValueString("TestArbitraryKeyValueString"), RUN_TEST);
    t.add(new TestArbitraryKeyValueUint("TestArbitraryKeyValueUint"), RUN_TEST);
    /**
     *  END OF TESTS
     */
    return t.run();
}
