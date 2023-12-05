/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "projects/graph/ArbitraryKvList/src/ArbitraryKvList.h"

class TestLinkListResetLeft : public TestBase {
private:
    ArbitraryKvList *root;
public:
    TestLinkListResetLeft(string n) { name = n; }

    ~TestLinkListResetLeft() { delete root; }

    bool test_setup() {
        debug(name + "::test_setup() starting");
        root = new ArbitraryKvList(true, false);
        if (root->count() != 0) return false;
        for (int i = 0; i < 10; i++) {
            root->insert("key_" + to_string(i), true);
            debug(name + "::test_setup() key=" + root->key());
        }
        debug(name + "::test_setup() 10-nodes expected. created " + to_string(root->count()) + " nodes.");
        return expect(root->count() == 10, ">> Expect root count is 10.") &&
               expect(root->key() == "key_0", ">> Expect current node is key_0.");
    }

    bool test_reset_left() {
        debug(name + "::test_reset_left()");
        debug(name + "::test_reset_left() current node:" + root->key());
        root->resetLeft();
        debug(name + "::test_reset_left() current node:" + root->key());
        return expect(root->count() == 10, ">> Expect root count is 10.") &&
               expect(root->key() == "key_1", ">> Expect current node is key_1 (found '" + root->key() + "')");
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect we can create initialize a default state.") &&
               expect(test_reset_left(), "Expect resetLeft can be used to change root pointer to left-most node.");
    }
};
