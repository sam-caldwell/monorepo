/**
 * @name Breadcrumbs/tests/TestBreadcrumbString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/Tester/TestBase/main.h"
#include "projects/common/Breadcrumbs/src/Breadcrumbs.h"
#include "projects/common/exceptions/exceptions.h"

class TestBreadcrumbString : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestBreadcrumbString(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestBreadcrumbString() {}

    bool test_push_count() {
        debug(name + "::test_push_count()");
        Breadcrumbs <string> tags;
        return expect(tags.push(to_string(0)), "push tag") &&
               expect(tags.count() == 1, "count 1") &&
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
               expect(tags.count() == 7, "count 7");
    }

    bool test_push_pop() {
        debug(name + "::test_push_pop()");
        Breadcrumbs <string> tags;
        return expect(tags.push("A"), "push tag") &&
               expect(tags.count() == 1, "count 1") &&
               expect(tags.pop(), "pop tag") &&
               expect(tags.count() == 0, "count 0");
    }

    bool test_toString() {
        debug(name + "::test_toString()");
        Breadcrumbs <string> tags;
        return expect(tags.push("A"), "push tag") &&
               expect(tags.count() == 1, "count 1") &&
               expect(tags.push("B"), "push tag") &&
               expect(tags.count() == 2, "count 2") &&
               expect(tags.push("C"), "push tag") &&
               expect(tags.count() == 3, "count 3") &&
               expect(tags.push("D"), "push tag") &&
               expect(tags.count() == 4, "count 4") &&
               expect(tags.push("E"), "push tag") &&
               expect(tags.count() == 5, "count 5") &&
               expect(tags.push("F"), "push tag") &&
               expect(tags.count() == 6, "count 6") &&
               expect(tags.push("G"), "push tag") &&
               expect(tags.count() == 7, "count 7") &&
               expect(tags.toString() == "A.B.C.D.E.F.G", "expect toString() ok");
    }

    bool test_back() {
        debug(name + "::test_push_count()");
        Breadcrumbs <string> tags;
        return expect(tags.push("A"), "push tag") &&
               expect(tags.count() == 1, "count 1") &&
               expect(tags.current() == "A", "current() ok") &&
               expect(tags.push("B"), "push tag") &&
               expect(tags.count() == 2, "count 2") &&
               expect(tags.current() == "B", "current() ok");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_push_count(), "Expect test_push_count() ok") &&
               expect(test_push_pop(), "Expect test_push_pop() ok") &&
               expect(test_toString(), "Expect test_toString() ok");
    }
};
