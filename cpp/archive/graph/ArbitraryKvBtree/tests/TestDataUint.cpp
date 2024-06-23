/**
 * @name ArbitraryKvBtree/tests/TestDataUint.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../src/ArbitraryKvBtree.h"

class TestDataUint : public TestBase {
private:
    ArbitraryKvBtree *tree;
public:
    TestDataUint(string n) {
        name = n;
        tree = new ArbitraryKvBtree(true);
    }

    ~TestDataUint() {
        if (tree) delete tree;
    }

    bool test_positive() {
        return expect(tree->insert("NodeM", (uint)(1048576)), "Expect we can insert NodeM(1048576)") &&
               expect(tree->count() == 1, "Expect count is 1") &&
               expect(tree->getType("NodeM") == Uint, "Expect getType ok(1)") &&
               expect(tree->getUint("NodeM") == (uint)(1048576), "Expect we can read value from NodeM(1048576)");
    }

    bool test_negative() {
        try {
            tree->getUint("NonExistent");
            return expect(false, "Expect that non-existent node throws runtime error.");
        } catch (out_of_range &e) {
            return expect(true, "Expect that non-existent node throws runtime error.");
        }
    }

    bool main() {
        ArbitraryKvBtree tree(true);
        return expect(test_positive(), "test_positive() ok") &&
               expect(test_negative(), "test_negative() ok");
    }

};/*end of class*/