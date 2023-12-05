/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.

 */
#include <iostream>
#include <stdlib.h>
#include <time.h>
#include "projects/graph/ArbitraryKvList/src/ArbitraryKvList.h"

class TestLessThan : public TestBase {
public:
    TestLessThan(string n) { name = n; }

    ~TestLessThan() {}

    bool test1() {
        ArbitraryKvNode *lhs = new ArbitraryKvNode("key_A", 1);
        ArbitraryKvNode *rhs = new ArbitraryKvNode("key_Z", 1);
        return LessThan(lhs, rhs);
    }

    bool test2() {
        ArbitraryKvNode *lhs = new ArbitraryKvNode("key_Z", 1);
        ArbitraryKvNode *rhs = new ArbitraryKvNode("key_A", 1);
        return !LessThan(lhs, rhs);
    }

    bool main() {
        return expect(test1(), "Expect that LHS less than RHS is true.") &&
               expect(test2(), "Expect that LHS less than RHS is not true");
    }

};