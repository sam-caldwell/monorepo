/**
 * @name ArbitraryKvBtree/tests/TestReBalance.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "projects/graph/ArbitraryKvBtree/src/ArbitraryKvBtree.h"

class TestReBalance : public TestBase {
public:
    TestReBalance(string n) {
        name = n;
    }

    ~TestReBalance() {/*noop*/}

    bool main() {
        debug(name + "::main() starting");
        ArbitraryKvBtree tree(true); // unique is turned on...
        return expect(tree.insert("node1", true), ">>Expect initial node to insert ok.") &&
               expect(!tree.insert("node1", false), ">>Expect duplicate to fail on insert.");
    }
};/*end of class*/
