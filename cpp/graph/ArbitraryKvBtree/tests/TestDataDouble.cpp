/**
 * @name ArbitraryKvBtree/tests/TestDataDouble.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/graph/ArbitraryKvBtree/src/ArbitraryKvBtree.h"

#ifdef PI
#undef PI
#endif /*PI*/
#define PI 3.141529

class TestDataDouble : public TestBase {
private:
    ArbitraryKvBtree *tree;
public:
    TestDataDouble(string n) {
        name = n;
        tree = new ArbitraryKvBtree(true);
    }

    ~TestDataDouble() {
        if (tree) delete tree;
    }

    bool test_positive() {
        return expect(tree->insert("NodeM", (double)(PI)), "Expect we can insert NodeM(PI)") &&
               expect(tree->count() == 1, "Expect count is 1") &&
               expect(tree->getType("NodeM") == Double, "Expect getType ok") &&
               expect(tree->getDouble("NodeM")==(double)(PI), "Expect we can read the value from NodeM(PI)");
    }

    bool test_negative() {
        try {
            tree->getDouble("NonExistent");
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