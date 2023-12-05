/**
 * @name ArbitraryKvBtree/tests/TestBasicInsertUnique.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "projects/graph/ArbitraryKvBtree/src/ArbitraryKvBtree.h"

#ifndef PI
#define PI 3.141529
#endif

class TestBasicInsertUnique : public TestBase {
private:
    ArbitraryKvBtree *tree;
public:
    TestBasicInsertUnique(string n) {
        name = n;
        tree = new ArbitraryKvBtree(true);
    }

    ~TestBasicInsertUnique() {
        debug("destroying class instance " + name);
        if(tree) delete tree;
    }

    bool main() {
        debug(name + "::main() starting");
        return expect(tree->insert("node1", true), "Expect bool insert ok") &&
               expect(tree->exists("node1"), "Expect node1 exists.") &&
               expect(tree->count() == 1, "Expect 1 nodes.  Encountered: " + to_string(tree->count())) &&
               expect(tree->balance() == 0, "expect balance 0.") &&
               expect(tree->duplicates() == 0, "expect duplicates 0") &&
               expect(tree->insert("node2", (double) (PI)), "Expect double insert ok") &&
               expect(tree->count() == 2, "Expect 2 nodes.  Encountered: " + to_string(tree->count())) &&
               expect(tree->balance() == 1, "expect balance 1") &&
               expect(tree->duplicates() == 0, "expect duplicates 0") &&
               expect(tree->insert("node3", (float) (PI)), "Expect float insert ok") &&
               expect(tree->count() == 3, "Expect 3 nodes.  Encountered: " + to_string(tree->count())) &&
               expect(tree->balance() == 2, "expect balance 2") &&
               expect(tree->duplicates() == 0, "expect duplicates 0") &&
               expect(tree->insert("node4", -1), "Expect int insert ok") &&
               expect(tree->count() == 4, "Expect 4 nodes.  Encountered: " + to_string(tree->count())) &&
               expect(tree->balance() == 3, "expect balance 3") &&
               expect(tree->duplicates() == 0, "expect duplicates 0") &&
               expect(tree->insert("node5", (string) "node5"), "Expect string insert ok") &&
               expect(tree->count() == 5, "Expect 5 nodes.  Encountered: " + to_string(tree->count())) &&
               expect(tree->balance() == 4, "expect balance 4") &&
               expect(tree->duplicates() == 0, "expect duplicates 0") &&
               expect(tree->insert("node6", (uint)(3)), "Expect uint insert ok") &&
               expect(tree->count() == 6, "Expect 6 nodes.  Encountered: " + to_string(tree->count())) &&
               expect(tree->balance() == 5, "expect balance 5") &&
               expect(tree->duplicates() == 0, "expect duplicates 0") &&
               expect(tree->getBool("node1"), "Expect getBool() for node1 will pass") &&
               expect(tree->getDouble("node2") == (double) (PI), "Expect getDouble() for node2 will pass") &&
               expect(tree->getFloat("node3") == (float) (PI), "Expect getFloat() for node3 will pass") &&
               expect(tree->getInt("node4") == (int) (-1), "Expect getInt() for node4 will pass") &&
               expect(tree->getString("node5") == "node5", "Expect getString() for node5 will pass.") &&
               expect(tree->getUint("node6") == (uint)(3), "Expect getUint() for node6 will pass");
    }

};/*end of class*/