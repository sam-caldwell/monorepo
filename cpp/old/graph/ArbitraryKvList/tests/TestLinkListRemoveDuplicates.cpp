/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "../src/ArbitraryKvList.h"

class TestLinkListRemoveDuplicates : public TestBase {
private:
    ArbitraryKvList *root;
public:
    TestLinkListRemoveDuplicates(string n) { name = n; }

    ~TestLinkListRemoveDuplicates() { delete root; }

    bool test_setup() {
        debug(name + "::test_setup() starting");
        root = new ArbitraryKvList(false);
        if (root->count() != 0) return false;
        for (int i = 0; i < 10; i++) {
            string key = "duplicate_key";
            root->insert(key, i);
            debug(name + "::test_setup() key(" + key + ") value:" + to_string(i));
        }
        root->insert("unique_key", 1337);
        return expect(root->count() == 11, ">> Expect root count is 11.") &&
               expect(root->key() == "duplicate_key", ">> Expect current node is duplicate_key. found:" + root->key());
    }

    bool test_remove_duplicates() {
        debug(name + "::test_remove_key()");
        return expect(root->remove("duplicate_key"), "Expect we can remove the given key") &&
               expect(root->count("duplicate_key") == 0,
                      "Expect that all matching keys are removed. "
                      "count=" + to_string(root->count("duplicate_key"))) &&
               expect(root->count() == 1, "Expect count ==1") &&
               expect(root->key() == "unique_key", "Expect we are pointing at unique_key node");
    }

    bool test_remove_last_node() {
        return expect(root->remove("unique_key"), "Expect last node can be removed.") &&
               expect(root->count() == 0, "Expect count is zero (0)");
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect we can create initialize a default state.") &&
               expect(test_remove_duplicates(), "Expect one call will remove all duplicate keys.") &&
               expect(test_remove_last_node(), "Expect all nodes are removed.");
    }
};