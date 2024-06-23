/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "../src/ArbitraryKvList.h"

class TestLinkListSwapSeparateNodes : public TestBase {
private:
    ArbitraryKvList *root;
public:
    TestLinkListSwapSeparateNodes(string n) { name = n; }

    ~TestLinkListSwapSeparateNodes() { delete root; }

    bool test_setup(int sz) {
        debug(name + "::test_setup() starting");
        root = new ArbitraryKvList;
        if (root->count() != 0) return false;
        for (int i = 0; i < sz; i++) {
            root->insert("key_" + to_string(i), true);
            debug(name + "::test_setup() key=" + root->key());
        }
        debug(name + "::test_setup() " + to_string(sz) + "-nodes expected. " +
              "created " + to_string(root->count()) + " nodes.");
        return expect(root->count() == sz, ">> Expect root count is "+ to_string(sz) +".") &&
               expect(root->key() == "key_0", ">> Expect current node is key_0.");
    }

    bool test_swap_separate(string A, string B, string expected) {
        debug("initial state: '" + root->dumpKeys() + "'");
        return expect(root->swap(A, B), "Expect we can swap key_3 and key_6") &&
               expect(root->dumpKeys() == expected,
                      "Expect our key order matches expectation: '" + root->dumpKeys() + "'");
    }

    bool main() {
        debug(name + "::main() starting...");
        string expected1 = "key_1,key_2,key_6,key_4,key_5,key_3,key_7,key_8,key_9,key_0,";
        string expected2 = "key_2,key_1,key_6,key_4,key_5,key_3,key_7,key_8,key_9,key_0,";
        string expected3 = "key_2,key_1,key_6,key_4,key_5,key_3,key_7,key_8,key_0,key_9,";
        return expect(test_setup(5), "Expect we can create initialize a default state.");// &&
//               expect(test_swap_separate("key_3","key_6", expected1),
//                      "expect we can swap two arbitrary, non-adjacent nodes.") &&
//               expect(test_swap_separate("key_1","key_2", expected2),
//                      "expect we can swap two arbitrary, non-adjacent nodes.") &&
//               expect(test_swap_separate("key_0","key_9", expected3),
//                      "expect we can swap two arbitrary, non-adjacent nodes.");
    }
};