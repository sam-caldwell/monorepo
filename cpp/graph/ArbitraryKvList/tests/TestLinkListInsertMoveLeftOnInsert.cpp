/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "projects/graph/ArbitraryKvList/src/ArbitraryKvList.h"

class TestLinkListInsertMoveLeftOnInsert : public TestBase {
private:
    ArbitraryKvList *root;
public:
    TestLinkListInsertMoveLeftOnInsert(string n) { name = n; }

    ~TestLinkListInsertMoveLeftOnInsert() { delete root; }

    bool test_setup() {
        debug(name + "::test_setup() starting");
        root = new ArbitraryKvList(true,true); //turn on unique and moveLeftOnInsert.
        if (root->count() != 0) return false;
        for (int i = 0; i < 10; i++) {
            root->insert("key_" + to_string(i), true);
            debug(name + "::test_setup() key=" + root->key());
        }
        debug(name + "::test_setup() 10-nodes expected. created " + to_string(root->count()) + " nodes.");
        return expect(root->count() == 10, ">> Expect root count is 10.") &&
               expect(root->key() == "key_9", ">> Expect current node is key_9.");
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect we can create initialize a default state.");
    }
};