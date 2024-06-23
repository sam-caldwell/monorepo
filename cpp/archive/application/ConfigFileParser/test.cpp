/**
 * @name ConfigFileParser/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../../../networking/TcpCommon/src/PipeResponse/main.h"
/**
 * Base tests
 */
#include "tests/TestInitialStateHappyPath.cpp"
#include "tests/TestParseYamlConfig.cpp"

using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();

    /**
     * Base tests
     */
    t.add(new TestInitialStateHappyPath("TestInitialStateHappyPath"), RUN_TEST);
    t.add(new TestParseYamlConfig("TestParseYamlConfig"), RUN_TEST);

    /**
     *  END OF TESTS
     */
    return t.run();
}