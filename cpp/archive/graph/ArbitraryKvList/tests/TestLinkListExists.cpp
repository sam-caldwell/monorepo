/**
* @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
*/
#include <iostream>
#include "../src/ArbitraryKvList.h"

class TestLinkListExists : public TestBase {
private:
    ArbitraryKvList *root;
public:
    TestLinkListExists(string n) { name = n; }

    ~TestLinkListExists() { delete root; }

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

    bool test_reset_left() {
        debug(name + "::test_reset_left()");
        debug(name + "::test_reset_left() current node:" + root->key());
        root->resetLeft();
        debug(name + "::test_reset_left() current node:" + root->key());
        return expect(root->count() == 10, ">> Expect root count is 10.") &&
               expect(root->key() == "key_1", ">> Expect current node is key_1.");
    }

    bool test_exists(string k) {
        debug(name + "::test_exists(" + k + ")");
        return root->exists(k);
    }

    bool tests_exists_all() {
        debug(name + "::tests_exists_all()");
        for (int i = 0; i < 10; i++) {
            if (!root->exists("key_" + to_string(i))) {
                debug(name+"::test_exists_all() failed on i="+to_string(i));
                return false;
            };
        }
        return true;
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect we can create initialize a default state.") &&
               expect(tests_exists_all(), "Expect all initial nodes exist.") &&
               expect(test_exists("key_3"), "Expect that key_3 exists") &&
               expect(test_reset_left(), "Expect resetLeft will reset root pointer to left-most node.") &&
               expect(test_exists("key_4"), "Expect that key_4 exists");
    }
};