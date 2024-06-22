/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "../src/ArbitraryKvList.h"

class TestLinkListSwapAdjacent : public TestBase {
private:
    ArbitraryKvList *root;
public:
    TestLinkListSwapAdjacent(string n) { name = n; }

    ~TestLinkListSwapAdjacent() { delete root; }

    bool test_setup() {
        debug(name + "::test_setup() starting");
        root = new ArbitraryKvList;
        if (root->count() != 0) return false;
        for (int i = 0; i < 10; i++) {
            root->insert("key_" + to_string(i), true);
            debug(name + "::test_setup() key=" + root->key());
        }
        debug(name + "::test_setup() 10-nodes expected. created " + to_string(root->count()) + " nodes.");
        return expect(root->count() == 10, ">> Expect root count is 10.") &&
               expect(root->key() == "key_0", ">> Expect current node is key_0.");
    }

    bool test_swap_adjacent() {
        string expected = "key_1,key_2,key_4,key_3,key_5,key_6,key_7,key_8,key_9,key_0,";
        debug("initial state: '" + root->dumpKeys() + "'");
        return expect(root->swap("key_3", "key_4"), "Expect we can swap key_2 and key_6") &&
               expect(root->dumpKeys()==expected, "Expect the swap was successful.");
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect we can create initialize a default state.") &&
               expect(test_swap_adjacent(), "expect we can swap two arbitrary, non-adjacent nodes.");
    }
};