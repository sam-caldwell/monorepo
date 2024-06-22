/**
 * @name ArbitraryKvBtree/tests/TestCountMatches.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryKvBtree.h"

class TestCountMatches : public TestBase {
private:
    ArbitraryKvBtree *tree;
public:
    /**
     * @name class constructor
     * @brief setup state
     *
     * @param n string
     */
    TestCountMatches(string n) {
        name = n;
        tree = new ArbitraryKvBtree(false); //Allow duplicates
    }

    /**
     * @name class destructor
     * @brief destroy state
     */
    ~TestCountMatches() {
        if (tree) delete tree;
    }

    /**
     * @name test_count_one
     * @brief Create a set of nodes and count matches and total count.  Each node created is unique so match count
     *        is always one (1)
     *
     * @return bool
     */
    bool test_count_one() {
        debug(name + "::test_count_byref() starting");
        return expect(tree->insert("node1", true),
                      ">>Expect we can insert a node:" + to_string(tree->count("node1"))) &&
               expect(tree->count("node1") == 1,
                      ">>Expect node1 count(1) Got " + to_string(tree->count("node1"))) &&
               expect(tree->insert("node2", "test"), ">>Expect we can insert a node") &&
               expect(tree->count("node2") == 1,
                      ">>Expect node2 count(1) Got " + to_string(tree->count("node1"))) &&
               expect(tree->insert("node3", (double) (1.0)), ">>Expect we can insert a node") &&
               expect(tree->count("node3") == 1,
                      ">>Expect node3 count(3) Got " + to_string(tree->count("node1"))) &&
               expect(tree->insert("node4", (uint)(1)), ">>Expect we can insert a node") &&
               expect(tree->count("node4") == 1,
                      ">>Expect node4 count(4) Got " + to_string(tree->count("node1"))) &&
               expect(tree->count() == 4, ">>Expect count ==4");
    }

    /**
     * @name test_count_many
     * @brief Create a set of duplicate nodes and expect duplicate count to increase
     *
     * @return bool
     */
    bool test_count_many() {
        return expect(tree->insert("node1", 3.0), "Create duplicate with number ok") &&
               expect(tree->count("node1") == 2,
                      ">>Expect node1 count(2) Got " + to_string(tree->count("node1"))) &&
               expect(tree->insert("node1", 3.1), "Create duplicate with number ok") &&
               expect(tree->count("node1") == 3,
                      ">>Expect node1 count(3) Got " + to_string(tree->count("node1"))) &&
               expect(tree->insert("node1", 3.2), "Create duplicate with number ok") &&
               expect(tree->count("node1") == 4,
                      ">>Expect node1 count(4) Got " + to_string(tree->count("node1"))) &&
               expect(tree->insert("node1", 3.2), "Create duplicate with number ok");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_count_one(), "expect test_count_one() ok") &&
               expect(test_count_many(), "expect test_count_many() ok");
    }
};/*end of class*/
