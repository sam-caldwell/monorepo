/**
 * @name ArbitraryKvBtree/tests/TestFind.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryKvBtree.h"

class TestFind : public TestBase {
private:
    ArbitraryKvBtree *tree;
public:
    /**
     * @name class constructor
     * @brief setup state
     *
     * @param n string
     */
    TestFind(string n) {
        name = n;
        tree = new ArbitraryKvBtree(false);
    }

    /**
     * @name class destructor
     * @brief destroy state
     */
    ~TestFind() {
        if (tree) delete tree;
    }

    /**
     * @name create_data
     * @brief create a set of nodes in an ArbitraryKvBtree
     *
     * @return bool
     */
    bool create_data() {
        return expect(tree->insert("NodeM", true), ">>Insert node1") &&
               expect(tree->insert("NodeH", true), ">>Insert node1") &&
               expect(tree->insert("NodeS", true), ">>Insert node1") &&
               expect(tree->insert("NodeD", true), ">>Insert node1") &&
               expect(tree->insert("NodeE", true), ">>Insert node1") &&
               expect(tree->insert("NodeA", true), ">>Insert node1") &&
               expect(tree->insert("NodeP", true), ">>Insert node1") &&
               expect(tree->insert("NodeQ", true), ">>Insert node1");
    }

    /**
     * @name test
     * @brief test that nodes exist as expected
     *
     * @return bool
     */
    bool test(){
        return expect(tree->exists("NodeA"), ">>find NodeA works") &&
               expect(tree->exists("NodeM"), ">>find NodeM works") &&
               expect(tree->exists("NodeH"), ">>find NodeH works") &&
               expect(tree->exists("NodeS"), ">>find NodeS works") &&
               expect(tree->exists("NodeE"), ">>find NodeE works") &&
               expect(!tree->exists("NonExistentNode"), ">>find NonExistentNode works");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        debug(name + "::main() starting");
        return expect(create_data(), "create_data() ok") &&
               expect(test(), "test() ok");
    }
};/*end of class*/
