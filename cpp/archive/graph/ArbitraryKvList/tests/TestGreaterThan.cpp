/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.

 */
#include <iostream>
#include <stdlib.h>
#include <time.h>
#include "../src/ArbitraryKvList.h"

class TestGreaterThan : public TestBase {
public:
    TestGreaterThan(string n) { name = n; }

    ~TestGreaterThan() {}

    bool test1() {
        ArbitraryKvNode *lhs = new ArbitraryKvNode("key_Z", 1);
        ArbitraryKvNode *rhs = new ArbitraryKvNode("key_A", 1);
        return GreaterThan(lhs, rhs);
    }

    bool test2() {
        ArbitraryKvNode *lhs = new ArbitraryKvNode("key_Z", 1);
        ArbitraryKvNode *rhs = new ArbitraryKvNode("key_A", 1);
        return !GreaterThan(rhs, lhs);
    }

    bool main() {
        return expect(test1(), "Expect that LHS greater than RHS is true.") &&
               expect(test2(), "Expect that LHS greater than RHS is not true");
    }

};