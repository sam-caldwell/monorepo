/**
 * @name types/tests/TestSortOrder.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../SortOrder.h"

class TestSortOrder : public TestBase {
public:
    TestSortOrder(string n) {
        name = n;
    }

    ~TestSortOrder() {/*noop*/}

    bool test_ascending() {
        debug(name + "::test_ascending()");
        SortOrder o;
        o = Ascending;
        return expect(SortOrderToString(o) == "ascending", "Expect SortOrderToString() to return correct result.");
    }

    bool test_descending() {
        debug(name + "::test_descending()");
        SortOrder o;
        o = Descending;
        return expect(SortOrderToString(o) == "descending", "Expect .SortOrderToString() to return correct result.");
    }

    bool test_undefined() {
        debug(name + "::test_undefined()");
        SortOrder o;
        o = Undefined;
        return expect(SortOrderToString(o) == "undefined", "Expect SortOrderToString() to return correct result.");
    }

    bool main() {
        debug(name + "::main()");
        return expect(test_ascending(), "Expect Ascending OK") &&
               expect(test_descending(), "Expect Descending OK") &&
               expect(test_undefined(), "Expect Undefined OK");
    }

};/*end of class*/