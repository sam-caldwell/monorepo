/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "../src/ArbitraryKvList.h"

class TestLinkListInsert : public TestBase {
private:
    ArbitraryKvList *root;
public:
    TestLinkListInsert(string n) { name = n; }

    ~TestLinkListInsert() { delete root; }

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

    bool test_insert() {
        return expect(root->insert("bool_true", true), "Expect we can insert boolean true") &&
               expect(root->insert("bool_false", false), "Expect we can insert boolean false") &&
               expect(root->insert("double_pi", (double) (3.141529)), "Expect we can insert double") &&
               expect(root->insert("float_pi", (float) (3.141529)), "Expect we can insert float") &&
               expect(root->insert("int_key", (int) (-3)), "Expect we can insert int") &&
               expect(root->insert("string_key", "string_value"), "Expect we can insert string node.") &&
               expect(root->insert("uint_key", (uint) (1)), "Expect we can insert uint node.");
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect we can create initialize a default state.") &&
               expect(test_insert(), "Expect that we can create nodes of various types.");
    }
};