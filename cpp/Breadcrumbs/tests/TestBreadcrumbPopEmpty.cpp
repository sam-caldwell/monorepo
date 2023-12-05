/**
 * @name Breadcrumbs/tests/TestBreadcrumbPopEmpty.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/Tester/TestBase/main.h"
#include "projects/common/Breadcrumbs/src/Breadcrumbs.h"
#include "projects/common/exceptions/exceptions.h"

class TestBreadcrumbPopEmpty : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestBreadcrumbPopEmpty(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestBreadcrumbPopEmpty() {}

    /**
     * @name test_pop_empty
     * @brief test that we can safely pop an empty list.
     */
    bool test_pop_empty() {
        debug(name + "::test_char()");
        Breadcrumbs<char> tags;
        return expect(tags.empty(), "expect empty") &&
               expect(tags.pop(), "pop() ok");
    }

   /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_pop_empty(), "Expect test_pop_empty() ok");
    }
};
