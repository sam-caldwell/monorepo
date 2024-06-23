/**
 * @name ArbitraryKvBtree/tests/TestDuplicates.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryKvBtree.h"

class TestDuplicates : public TestBase {
private:
    ArbitraryKvBtree *tree;
public:
    /**
     * @name class constructor
     * @brief setup state
     *
     * @param n string
     */
    TestDuplicates(string n) {
        name = n;
        tree = new ArbitraryKvBtree(false);
    }

    /**
     * @name class destructor
     * @brief destroy state
     */
    ~TestDuplicates() {
        if (tree) delete tree;
    }

    /**
     * @name test_duplicate_ok
     * @brief with uniqueness flag at false, create duplicates without error.
     *
     * @return bool
     */
    bool test_duplicate_ok() {
        debug(name + "::main() starting");
        return expect(tree->insert("node1", true), ">>Expect we can insert a node") &&
               expect(tree->count("node1") == 1, ">>Expect node1 count ==1") &&
               expect(tree->insert("node1", true), ">>Expect we can insert a node") &&
               expect(tree->count("node1") == 2, ">>Expect node1 count ==2") &&
               expect(tree->insert("node1", true), ">>Expect we can insert a node") &&
               expect(tree->count("node1") == 3, ">>Expect node1 count ==3") &&
               expect(tree->insert("node1", true), ">>Expect we can insert a node") &&
               expect(tree->count("node1") == 4, ">>Expect node1 count ==4") &&
               expect(tree->count() == 4, ">>Expect count ==4");
    }

    /**
     * @name test_duplicate_fails
     * @brief with uniqueness enabled verify that a duplicate insert() will return false.
     *
     * @return bool
     */
    bool test_duplicate_fails() {
        tree->unique(true);
        return expect(!tree->insert("node1", true), ">>Expect we can insert a node1") &&
               expect(tree->count("node1") == 4, ">>Expect node1 count ==1");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_duplicate_ok(), "Expect test_duplicate_ok() ok.") &&
               expect(test_duplicate_fails(), "Expect test_duplicate_fails() ok.");
    }

};/*end of class*/
