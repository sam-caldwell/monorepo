/**
* @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
*/
#include <iostream>
#include "../src/ArbitraryKvList.h"

class TestLinkListCountMatches : public TestBase {
private:
    ArbitraryKvList *root;
public:
    TestLinkListCountMatches(string n) { name = n; }

    ~TestLinkListCountMatches() { delete root; }

    bool test_setup() {
        debug(name + "::test_setup() starting");
        root = new ArbitraryKvList(false); // unique is turned off
        if (root->count() != 0) return false;
        for (int i = 0; i < 10; i++)
            root->insert("key_" + to_string(i), true);
        return expect(root->count() == 10, ">> Expect root count is 10.") &&
               expect(!root->unique(), ">>Expect unique is false") &&
               expect(root->key() == "key_0", ">> Expect current node is key_0.");
    }

    bool test_reset_left() {
        debug(name + "::test_reset_left()");
        root->resetLeft();
        return expect(root->count() == 10, ">> Expect root count is 10.") &&
               expect(root->key() == "key_1", ">> Expect current node is key_1.");
    }

    bool test_count_matches(string k, int count) {
        return root->count(k) == count;
    }

    bool test_insert(string k){
        debug(name+"::test_insert("+k+") starting");
        return root->insert(k,true);
    }
    bool tests_count_all() {
        debug(name + "::tests_count_all()");
        for (int i = 0; i < 10; i++) {
            if (root->count("key_" + to_string(i))!=1) {
                debug(name + "::tests_count_all() failed on node exists. i="+to_string(i)
                           + "count: "+to_string(root->count("key_" + to_string(i))));
                return false;
            };
        }
        return true;
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect we can create initialize a default state.") &&
               expect(tests_count_all(), "Expect we can find all initial nodes in set with count == 1.") &&
               expect(test_count_matches("key_3",1), "Expect that key_3 exists one time") &&
               expect(test_reset_left(), "Expect resetLeft will reset root pointer to left-most node.") &&
               expect(test_count_matches("key_3",1), "Expect that key_3 exists one time") &&
               expect(test_insert("key_3"), "Expect we can create duplicate key") &&
               expect(test_count_matches("key_3",2), "Expect that key_1 exists one time");
    }
};