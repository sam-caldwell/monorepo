/**
 * @name Breadcrumbs/tests/TestBreadcrumbInt.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../../../common/networking/TcpCommon/src/PipeResponse/main.h"
#include "../src/Breadcrumbs.h"
#include "../../../common/exceptions/exceptions.h"

class TestBreadcrumbInt : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestBreadcrumbInt(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestBreadcrumbInt() {}

    bool test_push_count() {
        debug(name + "::test_push_count()");
        Breadcrumbs<int> tags;
        return expect(tags.push(-1), "push tag") &&
               expect(tags.count() == 1, "count 1") &&
               expect(tags.push(0), "push tag") &&
               expect(tags.count() == 2, "count 2") &&
               expect(tags.push(1), "push tag") &&
               expect(tags.count() == 3, "count 3") &&
               expect(tags.push(2), "push tag") &&
               expect(tags.count() == 4, "count 4") &&
               expect(tags.push(3), "push tag") &&
               expect(tags.count() == 5, "count 5") &&
               expect(tags.push(4), "push tag") &&
               expect(tags.count() == 6, "count 6") &&
               expect(tags.push(5), "push tag") &&
               expect(tags.count() == 7, "count 7");
    }

    bool test_push_pop() {
        debug(name + "::test_push_pop()");
        Breadcrumbs <int> tags;
        return expect(tags.push(-1), "push tag") &&
               expect(tags.count() == 1, "count 1") &&
               expect(tags.pop(), "pop tag") &&
               expect(tags.count() == 0, "count 0");
    }

    bool test_toString() {
        debug(name + "::test_toString()");
        Breadcrumbs <int> tags;
        return expect(tags.push(1), "push tag") &&
               expect(tags.count() == 1, "count 1") &&
               expect(tags.push(2), "push tag") &&
               expect(tags.count() == 2, "count 2") &&
               expect(tags.push(3), "push tag") &&
               expect(tags.count() == 3, "count 3") &&
               expect(tags.push(4), "push tag") &&
               expect(tags.count() == 4, "count 4") &&
               expect(tags.push(5), "push tag") &&
               expect(tags.count() == 5, "count 5") &&
               expect(tags.push(6), "push tag") &&
               expect(tags.count() == 6, "count 6") &&
               expect(tags.push(7), "push tag") &&
               expect(tags.count() == 7, "count 7") &&
               expect(tags.toString() == "1.2.3.4.5.6.7", "expect toString() ok");
    }

    bool test_back() {
        debug(name + "::test_push_count()");
        Breadcrumbs <int> tags;
        return expect(tags.push(1), "push tag") &&
               expect(tags.count() == 1, "count 1") &&
               expect(tags.current() == 1, "current() ok") &&
               expect(tags.push(1), "push tag") &&
               expect(tags.count() == 2, "count 2") &&
               expect(tags.current() == 1, "current() ok");
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
