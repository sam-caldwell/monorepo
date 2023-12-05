/**
 * @name system/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <filesystem>
#include <iostream>
#include "projects/Tester/Tester/main.h"
/**
 * Base tests
 */
#include "projects/common/system/tests/TestGetEnvironment.cpp"
#include "projects/common/system/tests/TestSetEnvironment.cpp"
#include "projects/common/system/tests/TestUnsetEnvironment.cpp"
#include "projects/common/system/tests/TestFileExists.cpp"
#include "projects/common/system/tests/TestDetectDirectoryTraversal.cpp"
#include "projects/common/system/tests/TestExec.cpp"


using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();
    /**
     * Base tests
     */
    t.add(new TestGetEnvironment("TestGetEnvironment"), RUN_TEST);
    t.add(new TestSetEnvironment("TestSetEnvironment"), RUN_TEST);
    t.add(new TestUnsetEnvironment("TestUnsetEnvironment"), RUN_TEST);
    t.add(new TestFileExists("TestFileExists"), RUN_TEST);
    t.add(new TestDetectDirectoryTraversal("TestDetectDirectoryTraversal"), RUN_TEST);
    t.add(new TestExec("TestExec"));
    /**
     *  END OF TESTS
     */
    return t.run();
}
