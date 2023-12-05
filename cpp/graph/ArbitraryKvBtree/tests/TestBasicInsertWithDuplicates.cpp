/**
 * @name ArbitraryKvBtree/tests/TestBasicInsertWithDuplicates.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "projects/graph/ArbitraryKvBtree/src/ArbitraryKvBtree.h"

#ifndef PI
#define PI 3.141529
#endif

class TestBasicInsertWithDuplicates : public TestBase {
public:
    TestBasicInsertWithDuplicates(string n) {
        name = n;
    }

    ~TestBasicInsertWithDuplicates() {/*noop*/}

    bool main() {
        debug(name + "::main()");
        ArbitraryKvBtree tree(true); // unique is turned on...
        return expect(tree.insert("node1", true), "Expect bool insert ok") &&
               expect(tree.exists("node1"), "Expect node1 exists.") &&
               expect(tree.count() == 1, "Expect 1 nodes.  Encountered: " + to_string(tree.count())) &&
               expect(tree.balance() == 0, "expect balance 0.") &&
               expect(tree.duplicates() == 0, "expect duplicates 0") &&
               expect(tree.insert("node2", (double) (PI)), "Expect double insert ok") &&
               expect(tree.count() == 2, "Expect 2 nodes.  Encountered: " + to_string(tree.count())) &&
               expect(tree.balance() == 1, "expect balance 1") &&
               expect(tree.duplicates() == 0, "expect duplicates 0") &&
               expect(tree.insert("node3", (float) (PI)), "Expect float insert ok") &&
               expect(tree.count() == 3, "Expect 3 nodes.  Encountered: " + to_string(tree.count())) &&
               expect(tree.balance() == 2, "expect balance 2") &&
               expect(tree.duplicates() == 0, "expect duplicates 0") &&
               expect(tree.insert("node4", -1), "Expect int insert ok") &&
               expect(tree.count() == 4, "Expect 4 nodes.  Encountered: " + to_string(tree.count())) &&
               expect(tree.balance() == 3, "expect balance 3") &&
               expect(tree.duplicates() == 0, "expect duplicates 0") &&
               expect(tree.insert("node5", (string) "node5"), "Expect string insert ok") &&
               expect(tree.count() == 5, "Expect 5 nodes.  Encountered: " + to_string(tree.count())) &&
               expect(tree.balance() == 4, "expect balance 4") &&
               expect(tree.duplicates() == 0, "expect duplicates 0") &&
               expect(tree.insert("node6", (uint) (3)), "Expect uint insert ok") &&
               expect(tree.count() == 6, "Expect 6 nodes.  Encountered: " + to_string(tree.count())) &&
               expect(tree.balance() == 5, "expect balance 5") &&
               expect(tree.duplicates() == 0, "expect duplicates 0") &&
               /*
                * unique is turned off...adding duplicates.
                */
               expect(!tree.unique(false), "Expect we can turn off unique") &&
               expect(tree.insert("node1", false), "Expect we can create duplicate bool node") &&
               expect(tree.count() == 7, "Expect 7 nodes") &&
               expect(tree.balance() == 6, "expect balance 6. actual:" + to_string(tree.balance())) &&
               expect(tree.insert("node1", 3), "Expect we can create duplicate bool node") &&
               expect(tree.count() == 8, "Expect 8 nodes") &&
               expect(tree.balance() == 7, "expect balance 7. actual:" + to_string(tree.balance())) &&
               expect(tree.insert("node7", 3), "Expect we can create duplicate bool node") &&
               expect(tree.count() == 9, "Expect 9 nodes") &&
               expect(tree.balance() == 8, "expect balance 8. actual:" + to_string(tree.balance()));
    }

};/*end of class*/