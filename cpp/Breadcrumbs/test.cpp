/**
 * @name projects/common/Breadcrumbs/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <stack>
#include <iostream>
#include "projects/Tester/Tester/main.h"
/**
 *
 */
#include "tests/TestBreadcrumbString.cpp"
#include "tests/TestBreadcrumbInt.cpp"
#include "tests/TestBreadcrumbUnsigned.cpp"
#include "tests/TestBreadcrumbFloat.cpp"
#include "tests/TestBreadcrumbDouble.cpp"
#include "tests/TestBreadcrumbClear.cpp"
#include "tests/TestBreadcrumbEmpty.cpp"
#include "tests/TestBreadcrumbPrev.cpp"


using namespace std;

int main(int argc, char *argv[]) {
    Tester t(argc, argv);
    t.enable_debug();
    /**
     * Base tests
     */
    t.add(new TestBreadcrumbString("TestBreadcrumbString"), RUN_TEST);
    t.add(new TestBreadcrumbInt("TestBreadcrumbInt"), RUN_TEST);
    t.add(new TestBreadcrumbUnsigned("TestBreadcrumbUnsigned"), RUN_TEST);
    t.add(new TestBreadcrumbFloat("TestBreadcrumbFloat"), RUN_TEST);
    t.add(new TestBreadcrumbDouble("TestBreadcrumbDouble"), RUN_TEST);
    t.add(new TestBreadcrumbClear("TestBreadcrumbClear"), RUN_TEST);
    t.add(new TestBreadcrumbEmpty("TestBreadcrumbEmpty"), RUN_TEST);
    t.add(new TestBreadcrumbPrev("TestBreadcrumbPrev"), RUN_TEST);

    /**
     *  END OF TESTS
     */
    return t.run();
}
