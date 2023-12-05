/**
 * @name Breadcrumbs/tests/TestBreadcrumbEmpty.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/Tester/TestBase/main.h"
#include "projects/common/Breadcrumbs/src/Breadcrumbs.h"
#include "projects/common/exceptions/exceptions.h"

class TestBreadcrumbEmpty : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestBreadcrumbEmpty(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestBreadcrumbEmpty() {}

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        debug(name + "::main()");
        Breadcrumbs <string> tags;
        return expect(tags.empty(), "expect empty") &&
               expect(tags.push(to_string(0)), "push tag") &&
               expect(tags.count() == 1, "count 1") &&
               expect(!tags.empty(), "Expect not empty") &&
               expect(tags.push(to_string(1)), "push tag") &&
               expect(tags.count() == 2, "count 2") &&
               expect(tags.push(to_string(-1)), "push tag") &&
               expect(tags.count() == 3, "count 3") &&
               expect(tags.push("A"), "push tag") &&
               expect(tags.count() == 4, "count 4") &&
               expect(tags.push("B"), "push tag") &&
               expect(tags.count() == 5, "count 5") &&
               expect(tags.push("C"), "push tag") &&
               expect(tags.count() == 6, "count 6") &&
               expect(tags.push("D"), "push tag") &&
               expect(tags.count() == 7, "count 7") &&
               expect(tags.clear(), "expect clear() ok") &&
               expect(tags.count() == 0, "count 0") &&
               expect(tags.empty(), "expect empty");
    }
};
