/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "projects/graph/ArbitraryKvList/src/ArbitraryKvList.h"

class TestLinkListMoveLeftRight : public TestBase {
private:
    ArbitraryKvList *root;
public:
    TestLinkListMoveLeftRight(string n) { name = n; }

    ~TestLinkListMoveLeftRight() { delete root; }

    bool test_setup() {
        debug(name + "::test_setup() starting");
        root = new ArbitraryKvList;
        if (root->count() != 0) return false;
        for (int i = 0; i < 10; i++) {
            string key = "key_" + to_string(i);
            root->insert(key, true);
            debug(name + "::test_setup() key=" + key);
        }
        debug(name + "::test_setup() 10-nodes expected. created " + to_string(root->count()) + " nodes.");
        return expect(root->count() == 10, ">> Expect root count is 10.") &&
               expect(root->key() == "key_0", ">> Expect current node is key_0.");
    }

    bool test_move(){
        debug(root->dumpKeys());
        return expect(root->key() == "key_0", ">> Expect current node is key_0.") &&
               expect(root->moveLeft(), ">> Expect move left works.") &&
               expect(root->key() == "key_9", "Expect we are on key_9") &&
               expect(root->moveLeft(), ">> Expect move left works.") &&
               expect(root->key() == "key_8", "Expect we are on key_8") &&
               expect(root->moveLeft(), ">> Expect move left works.") &&
               expect(root->key() == "key_7", "Expect we are on key_7") &&
               expect(root->moveRight(), ">> Expect move right works.") &&
               expect(root->key() == "key_8", "Expect we are on key_8") &&
               expect(root->moveRight(), ">> Expect move right works.") &&
               expect(root->key() == "key_9", "Expect we are on key_9") &&
               expect(root->moveRight(), ">> Expect move right works.") &&
               expect(root->key() == "key_0", "Expect we are on key_0") &&
               expect(!root->moveRight(), ">> Expect move right works when no next node exists.") &&
               expect(root->key() == "key_0", "Expect we are on key_0");
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect we can create initialize a default state.") &&
               expect(test_move(), "Expect moveLeft works");
    }
};