/**
 * @name ArbitraryKvBtree/tests/TestDataString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../src/ArbitraryKvBtree.h"

class TestDataString : public TestBase {
private:
    ArbitraryKvBtree *tree;
public:
    TestDataString(string n) {
        name = n;
        tree = new ArbitraryKvBtree(true);
    }

    ~TestDataString() {
        if (tree) delete tree;
    }

    bool test_positive() {
        return expect(tree->insert("NodeM", (string)("Hi")), "Expect we can insert NodeM(Hi)") &&
               expect(tree->count() == 1, "Expect count is 1") &&
               expect(tree->getType("NodeM") == String, "Expect getType ok(1)") &&
               expect(tree->getString("NodeM") == "Hi", "Expect we can read the value from NodeM(true)");
    }

    bool test_negative() {
        try {
            tree->getString("NonExistent");
            return expect(false, "Expect that non-existent node throws out_of_range error.");
        } catch (out_of_range &e) {
            return expect(true, "Expect that non-existent node throws out_of_range error.");
        }
    }

    bool main() {
        ArbitraryKvBtree tree(true);
        return expect(test_positive(), "test_positive() ok") &&
               expect(test_negative(), "test_negative() ok");
    }

};/*end of class*/