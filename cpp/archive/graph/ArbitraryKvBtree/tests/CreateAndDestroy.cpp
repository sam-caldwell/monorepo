/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "../src/ArbitraryKvBtree.h"

class TestCreateAndDestroy : public TestBase {
private:
public:
    /**
     * @name class constructor
     * @brief setup state.
     *
     * @param n string
     */
    TestCreateAndDestroy(string n) { name = n; }

    /**
     * @name class destructor
     * @brief destroy state
     */
    ~TestCreateAndDestroy() {}

    /**
     * @name create_destroy_test0
     * @brief expect we can create dynamic class instance and destroy it without error
     *
     * @return bool
     */
    bool create_destroy_test0() {
        ArbitraryKvBtree *root;
        root = new ArbitraryKvBtree();
        delete root;
        return true;
    }

    /**
     * @name create_destroy_test1
     * @brief expect we can create dynamic class instance and destroy it without error
     *
     * @return bool
     */
    bool create_destroy_test1() {
        ArbitraryKvBtree *root;
        root = new ArbitraryKvBtree(true);
        delete root;
        return true;
    }

    /**
     * @name create_destroy_test2
     * @brief expect we can create dynamic class instance and destroy it without error
     *
     * @return bool
     */
    bool create_destroy_test2() {
        ArbitraryKvBtree *root;
        root = new ArbitraryKvBtree(false);
        delete root;
        return true;
    }
    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        debug(name + "::main() starting...");
        return expect(create_destroy_test0(), "Expect default constructor works.") &&
               expect(create_destroy_test1(), "Expect unique constructor (true)") &&
               expect(create_destroy_test2(), "Expect unique constructor (false)");
    }
};
