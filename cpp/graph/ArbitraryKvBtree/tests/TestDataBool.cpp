/**
 * @name ArbitraryKvBtree/tests/TestDataBool.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/graph/ArbitraryKvBtree/src/ArbitraryKvBtree.h"

class TestDataBool : public TestBase {
private:
    ArbitraryKvBtree *tree;
public:
    TestDataBool(string n) {
        name = n;
        tree = new ArbitraryKvBtree(true);
    }

    ~TestDataBool() {
        if (tree) delete tree;
    }

    bool test_positive() {
        return expect(tree->insert("NodeM", true), "Expect we can insert NodeM(true)") &&
               expect(tree->insert("NodeG", false), "Expect we can insert NodeG(false)") &&
               expect(tree->count() == 2, "Expect count is 2") &&
               expect(tree->getType("NodeM") == Bool, "Expect getType ok(1)") &&
               expect(tree->getType("NodeG") == Bool, "Expect getType ok(2)") &&
               expect(tree->getBool("NodeM"), "Expect we can read the value from NodeM(true)") &&
               expect(!tree->getBool("NodeG"), "Expect we can read the value from NodeG(false)");
    }

    bool test_negative() {
        try {
            tree->getBool("NonExistent");
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