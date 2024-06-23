/**
 * @name ArbitraryKvBtree/tests/TestInitialState.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "../src/ArbitraryKvBtree.h"

class TestInitialState : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup state
     *
     * @param n string
     */
    TestInitialState(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy state
     */
    ~TestInitialState() {/*noop*/}

    /**
     * @name test0
     * @brief Expect that the default state of an ArbitraryKvBtree is unique=true,
     *        node count is zero and balance is zero.
     *
     * @return bool
     */
    bool test0() {
        debug(name + "::test0()");
        ArbitraryKvBtree tree;
        return expect(tree.unique(), "Expect unique is true. Found " + to_string(tree.unique())) &&
               expect(tree.count() == 0, "Expect count is 0") &&
               expect(tree.balance() == 0, "Expect balance is 0");
    }

    /**
     * @name test1
     * @brief Expect that the state of an ArbitraryKvBtree can be configured with unique bool true and initial state
     *        is zero node count and zero balance
     *
     * @return bool
     */
    bool test1() {
        debug(name + "::test1()");
        ArbitraryKvBtree tree(true);
        return expect(tree.unique(), "Expect unique is true.") &&
               expect(tree.count() == 0, "Expect count is 0") &&
               expect(tree.balance() == 0, "Expect balance is 0");
    }

    /**
     * @name test2
     * @brief Expect that the state of an ArbitraryKvBtree can be configured with unique bool false and initial state
     *        is zero node count and zero balance
     *
     * @return bool
     */
    bool test2() {
        debug(name + "::test2()");
        ArbitraryKvBtree tree(false);
        return expect(!tree.unique(), "Expect unique is false.") &&
               expect(tree.count() == 0, "Expect count is 0") &&
               expect(tree.balance() == 0, "Expect balance is 0");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test0(), "Expect default constructor to work") &&
               expect(test1(), "Expect unique false constructor to work") &&
               expect(test2(), "Expect unique true constructor to work");
    }

};/*end of class*/