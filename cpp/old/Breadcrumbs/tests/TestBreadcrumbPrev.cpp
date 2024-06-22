/**
 * @name Breadcrumbs/tests/TestBreadcrumbPrev.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../../../common/networking/TcpCommon/src/PipeResponse/main.h"
#include "../src/Breadcrumbs.h"
#include "../../../common/exceptions/exceptions.h"

class TestBreadcrumbPrev : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestBreadcrumbPrev(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestBreadcrumbPrev() {}

    /**
     * @name test_char
     * @brief test prev() with char
     */
    bool test_char() {
        debug(name + "::test_char()");
        Breadcrumbs<char> tags;
        return expect(tags.empty(), "expect empty") &&
               expect(tags.push('0'), "push tag") &&
               expect(tags.count() == 1, "count 1") &&
               expect(tags.push('1'), "push tag") &&
               expect(tags.count() == 2, "count 2") &&
               expect(tags.prev() == '0', "expect prev() ok") &&
               expect(tags.push('2'), "push tag") &&
               expect(tags.count() == 3, "count 3") &&
               expect(tags.prev() == '1', "expect prev() ok") &&
               expect(tags.push('3'), "push tag") &&
               expect(tags.count() == 4, "count 4") &&
               expect(tags.push('4'), "push tag") &&
               expect(tags.count() == 5, "count 5") &&
               expect(tags.push('5'), "push tag") &&
               expect(tags.count() == 6, "count 6") &&
               expect(tags.push('6'), "push tag") &&
               expect(tags.count() == 7, "count 7") &&
               expect(tags.prev() == '5', "expect prev() ok") &&
               expect(tags.clear(), "expect clear() ok") &&
               expect(tags.prev() == 0x00, "expect prev() ok") &&
               expect(tags.count() == 0, "count 0");
    }

    /**
     * @name test_double
     * @brief test prev() with double
     */
    bool test_double() {
        debug(name + "::test_double()");
        Breadcrumbs<double> tags;
        return expect(tags.empty(), "expect empty") &&
               expect(tags.push(0), "push tag") &&
               expect(tags.count() == 1, "count 1") &&
               expect(tags.push(1), "push tag") &&
               expect(tags.count() == 2, "count 2") &&
               expect(tags.prev() == 0, "expect prev() ok") &&
               expect(tags.push(2), "push tag") &&
               expect(tags.count() == 3, "count 3") &&
               expect(tags.prev() == 1, "expect prev() ok") &&
               expect(tags.push(3), "push tag") &&
               expect(tags.count() == 4, "count 4") &&
               expect(tags.push(4), "push tag") &&
               expect(tags.count() == 5, "count 5") &&
               expect(tags.push(5), "push tag") &&
               expect(tags.count() == 6, "count 6") &&
               expect(tags.push(6), "push tag") &&
               expect(tags.count() == 7, "count 7") &&
               expect(tags.prev() == 5, "expect prev() ok") &&
               expect(tags.clear(), "expect clear() ok") &&
               expect(tags.prev() == 0, "expect prev() ok") &&
               expect(tags.count() == 0, "count 0");
    }

    /**
     * @name test_float
     * @brief test prev() with float
     */
    bool test_float() {
        debug(name + "::test_float()");
        Breadcrumbs<double> tags;
        return expect(tags.empty(), "expect empty") &&
               expect(tags.push(0), "push tag") &&
               expect(tags.count() == 1, "count 1") &&
               expect(tags.push(1), "push tag") &&
               expect(tags.count() == 2, "count 2") &&
               expect(tags.prev() == 0, "expect prev() ok") &&
               expect(tags.push(2), "push tag") &&
               expect(tags.count() == 3, "count 3") &&
               expect(tags.prev() == 1, "expect prev() ok") &&
               expect(tags.push(3), "push tag") &&
               expect(tags.count() == 4, "count 4") &&
               expect(tags.push(4), "push tag") &&
               expect(tags.count() == 5, "count 5") &&
               expect(tags.push(5), "push tag") &&
               expect(tags.count() == 6, "count 6") &&
               expect(tags.push(6), "push tag") &&
               expect(tags.count() == 7, "count 7") &&
               expect(tags.prev() == 5, "expect prev() ok") &&
               expect(tags.clear(), "expect clear() ok") &&
               expect(tags.prev() == 0, "expect prev() ok") &&
               expect(tags.count() == 0, "count 0");
    }

    /**
     * @name test_int
     * @brief test prev() with int
     */
    bool test_int() {
        debug(name + "::test_int()");
        Breadcrumbs<int> tags;
        return expect(tags.empty(), "expect empty") &&
               expect(tags.push(-1), "push tag") &&
               expect(tags.count() == 1, "count 1") &&
               expect(tags.push(1), "push tag") &&
               expect(tags.count() == 2, "count 2") &&
               expect(tags.prev() == -1, "expect prev() ok") &&
               expect(tags.push(2), "push tag") &&
               expect(tags.count() == 3, "count 3") &&
               expect(tags.prev() == 1, "expect prev() ok") &&
               expect(tags.push(3), "push tag") &&
               expect(tags.count() == 4, "count 4") &&
               expect(tags.push(4), "push tag") &&
               expect(tags.count() == 5, "count 5") &&
               expect(tags.push(5), "push tag") &&
               expect(tags.count() == 6, "count 6") &&
               expect(tags.push(6), "push tag") &&
               expect(tags.count() == 7, "count 7") &&
               expect(tags.prev() == 5, "expect prev() ok") &&
               expect(tags.clear(), "expect clear() ok") &&
               expect(tags.prev() == 0, "expect prev() ok") &&
               expect(tags.count() == 0, "count 0");
    }

    /**
     * @name test_string
     * @brief test prev() with string
     */
    bool test_string() {
        debug(name + "::test_string()");
        Breadcrumbs <string> tags;
        return expect(tags.empty(), "expect empty") &&
               expect(tags.push("0"), "push tag") &&
               expect(tags.count() == 1, "count 1") &&
               expect(tags.push("1"), "push tag") &&
               expect(tags.count() == 2, "count 2") &&
               expect(tags.prev() == "0", "expect prev() ok") &&
               expect(tags.push("2"), "push tag") &&
               expect(tags.count() == 3, "count 3") &&
               expect(tags.prev() == "1", "expect prev() ok") &&
               expect(tags.push("3"), "push tag") &&
               expect(tags.count() == 4, "count 4") &&
               expect(tags.push("4"), "push tag") &&
               expect(tags.count() == 5, "count 5") &&
               expect(tags.push("5"), "push tag") &&
               expect(tags.count() == 6, "count 6") &&
               expect(tags.push("6"), "push tag") &&
               expect(tags.count() == 7, "count 7") &&
               expect(tags.prev() == "5", "expect prev() ok") &&
               expect(tags.clear(), "expect clear() ok") &&
               expect(tags.prev() == "", "expect prev() ok") &&
               expect(tags.count() == 0, "count 0");
    }

    /**
     * @name test_unsigned
     * @brief test prev() with unsigned
     */
    bool test_unsigned() {
        debug(name + "::test_unsigned()");
        Breadcrumbs<unsigned> tags;
        return expect(tags.empty(), "expect empty") &&
               expect(tags.push(0), "push tag") &&
               expect(tags.count() == 1, "count 1") &&
               expect(tags.push(1), "push tag") &&
               expect(tags.count() == 2, "count 2") &&
               expect(tags.prev() == 0, "expect prev() ok") &&
               expect(tags.push(2), "push tag") &&
               expect(tags.count() == 3, "count 3") &&
               expect(tags.prev() == 1, "expect prev() ok") &&
               expect(tags.push(3), "push tag") &&
               expect(tags.count() == 4, "count 4") &&
               expect(tags.push(4), "push tag") &&
               expect(tags.count() == 5, "count 5") &&
               expect(tags.push(5), "push tag") &&
               expect(tags.count() == 6, "count 6") &&
               expect(tags.push(6), "push tag") &&
               expect(tags.count() == 7, "count 7") &&
               expect(tags.prev() == 5, "expect prev() ok") &&
               expect(tags.clear(), "expect clear() ok") &&
               expect(tags.prev() == 0, "expect prev() ok") &&
               expect(tags.count() == 0, "count 0");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_char(), "Expect test_char() ok") &&
               expect(test_double(), "Expect test_double() ok") &&
               expect(test_float(), "Expect test_double() ok") &&
               expect(test_int(), "Expect test_int() ok") &&
               expect(test_string(), "Expect test_string() ok") &&
               expect(test_unsigned(), "Expect test_unsigned() ok");
    }
};
