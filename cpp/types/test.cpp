/**
 * @name types/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/Tester/Tester/main.h"

/**
 * SortOrder tests
 */
#include "tests/TestSortOrder.cpp"
/**
 * ValueTypes tests
 */
#include "tests/TestValueTypes.cpp"
/**
 * ConfigStateMap tests
 */
#include "tests/ConfigStateMap/TestBasicFunctionality.cpp"
#include "tests/ConfigStateMap/TestAdd.cpp"
/**
 * FileFormats tests
 */
#include "tests/TestFileFormats.cpp"
/**
 * Mapper tests
 */
#include "tests/TestMapper.cpp"

using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();

    /**
     * SortOrder tests
     */
    t.add(new TestSortOrder("TestSortOrder"), RUN_TEST);
    /**
     * TestValueTypes
     */
    t.add(new TestValueTypes("TestValueTypes"), RUN_TEST);
    /**
     * FileFormats tests
     */
    t.add(new TestFileFormats("TestFileFormats"), RUN_TEST);
    /**
     * Mapper tests.
     */
    t.add(new TestMapper("TestMapper"), RUN_TEST);
    /**
     * ConfigStateMap tests
     */
    t.add(new TestBasicFunctionality("TestBasicFunctionality"), RUN_TEST);
    t.add(new TestAdd("TestAdd"), RUN_TEST);


    /**
     *  END OF TESTS
     */
    return t.run();
}
