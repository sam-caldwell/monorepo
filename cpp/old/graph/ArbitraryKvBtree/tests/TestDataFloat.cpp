/**
 * @name ArbitraryKvBtree/tests/TestDataFloat.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../src/ArbitraryKvBtree.h"

#ifdef PI
#undef PI
#endif /*PI*/
#define PI 3.141529

class TestDataFloat : public TestBase {
private:
    ArbitraryKvBtree *tree;
public:
    TestDataFloat(string n) {
        name = n;
        tree = new ArbitraryKvBtree(true);
    }

    ~TestDataFloat() {
        if (tree) delete tree;
    }

    bool test_positive() {
        return expect(tree->insert("NodeM", (float)(PI)), "Expect we can insert NodeM(PI)") &&
               expect(tree->count() == 1, "Expect count is 1") &&
               expect(tree->getType("NodeM") == Float, "Expect getType ok") &&
               expect(tree->getFloat("NodeM")==(float)(PI), "Expect we can read the value from NodeM(PI)");
    }

    bool test_negative() {
        try {
            tree->getFloat("NonExistent");
            return expect(false, "Expect that non-existent node throws out_of_range error.");
        } catch (out_of_range &e) {
            return expect(true, "Expect that non-existent node throws out_of_range error.");
        }
    }

    bool main() {
        ArbitraryKvBtree tree(true);
        return expect(tree.insert("NodeM", (float)(PI)), "Expect we can insert NodeM(PI)") &&
               expect(tree.count() == 1, "Expect count is 1") &&
               expect(tree.getType("NodeM") == Float, "Expect getType ok(1)") &&
               expect(tree.getFloat("NodeM")==(float)(PI), "Expect we can read the value from NodeM(PI)");
    }
};/*end of class*/