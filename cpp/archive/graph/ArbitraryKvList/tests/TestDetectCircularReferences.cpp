/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * Test a linked list node for a circular self-reference.
 */
#include <iostream>
#include <stdlib.h>
#include <time.h>
#include "../src/ArbitraryKvList.h"

class TestDetectCircularReferences : public TestBase {
public:
    TestDetectCircularReferences(string n) { name = n; }

    ~TestDetectCircularReferences() {}

    bool test_no_circular_references() {
        ArbitraryKvNode *node;
        node = new ArbitraryKvNode("center", 1);
        node->left = new ArbitraryKvNode("left", 1);
        node->right = new ArbitraryKvNode("right", 1);
        node->left->right = node;
        node->right->left = node;
        return !detectCircular(node);
    }

    bool test_circular_reference1() {
        ArbitraryKvNode *node;
        node = new ArbitraryKvNode("self", 1);
        node->left = node;
        return detectCircular(node);
    }

    bool test_circular_reference2() {
        ArbitraryKvNode *node;
        node = new ArbitraryKvNode("self", 1);
        node->right = node;
        return detectCircular(node);
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_no_circular_references(), "Expect false when no circular loop exists.") &&
               expect(test_circular_reference1(), "Expect that circular reference (left) is detected") &&
               expect(test_circular_reference2(), "Expect that circular reference (right) is detected");
    }
};