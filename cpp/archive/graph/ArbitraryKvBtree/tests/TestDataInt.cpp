/**
 * @name ArbitraryKvBtree/tests/TestDataInt.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../src/ArbitraryKvBtree.h"

class TestDataInt : public TestBase {
private:
    ArbitraryKvBtree *tree;
public:
    TestDataInt(string n) {
        name = n;
        tree = new ArbitraryKvBtree(true);
    }

    ~TestDataInt() {
        if (tree) delete tree;
    }

    bool test_positive() {
        return expect(tree->insert("NodeM", -1048576), "Expect we can insert NodeM(-1048576)") &&
               expect(tree->count() == 1, "Expect count is 1") &&
               expect(tree->getType("NodeM") == Int, "Expect getType ok") &&
               expect(tree->getInt("NodeM") == -1048576, "Expect we can read the value from NodeM(-1048576)");
    }

    bool test_negative() {
        try {
            tree->getInt("NonExistent");
            return expect(false, "Expect that non-existent node throws out_of_range error.");
        } catch (out_of_range &e) {
            return expect(true, "Expect that non-existent node throws out_of_range error.");
        }
    }

    bool main() {
        ArbitraryKvBtree tree(true);
        return expect(tree.insert("NodeM", (int) (-1048576)), "Expect we can insert NodeM(-1048576)") &&
               expect(tree.count() == 1, "Expect count is 1") &&
               expect(tree.getType("NodeM") == Int, "Expect getType ok(1)") &&
               expect(tree.getInt("NodeM") == (int) (-1048576), "Expect we can read value from NodeM(-1048576)");
    }
};/*end of class*/