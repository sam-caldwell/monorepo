/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "../src/ArbitraryKvList.h"

class TestLinkListUnique : public TestBase {
private:
    ArbitraryKvList *root;
public:
    TestLinkListUnique(string n) { name = n; }

    ~TestLinkListUnique() { delete root; }

    bool test_setup() {
        debug(name + "::test_setup() starting");
        root = new ArbitraryKvList;
        return (root);
    }

    bool test_unique_getter() {
        return root->unique();
    }

    bool test_unique_setter(bool n) {
        return root->unique(n);
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect we can create initialize a default state.") &&
               expect(!test_unique_getter(), "Expect unique is false.") &&
               expect(test_unique_setter(true), "Expect unique is properly set to true") &&
               expect(test_unique_getter(), "Expect unique is true.") &&
               expect(!test_unique_setter(false), "Expect unique is properly set to false") &&
               expect(!test_unique_getter(), "Expect unique is false.");
    }
};