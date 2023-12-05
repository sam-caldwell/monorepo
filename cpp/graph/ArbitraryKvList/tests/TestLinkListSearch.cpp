/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "projects/graph/ArbitraryKvList/src/ArbitraryKvList.h"

class TestLinkListSearch : public TestBase {
private:
    ArbitraryKvList *root;
public:
    TestLinkListSearch(string n) { name = n; }

    ~TestLinkListSearch() { delete root; }

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

    bool test_search(string s) {
        return root->search(s);
    }

    bool test_position(string s) {
        return root->key() == s;
    }

    bool tests_exists_all() {
        debug(name + "::tests_exists_all()");
        for (int i = 0; i < 10; i++) {
            if (!root->search("key_" + to_string(i))) {
                debug(name+"::test_exists_all() failed on node exists. i="+to_string(i));
                return false;
            };
            if (root->key() !="key_" + to_string(i) ){
                debug(name+"::test_exists_all() failed on root location. i="+to_string(i));
                return false;
            }
        }
        return true;
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect we can create initialize a default state.") &&
               expect(tests_exists_all(), "Expect we can find all initial nodes in set.") &&
               expect(test_search("key_5"), "Expect we can search for key_5 (exists)") &&
               expect(test_position("key_5"), "Expect that search left root at key_5") &&
               expect(!test_search("not_found1"), "Expect we can search for non-existent node") &&
               expect(test_position("key_5"), "Expect that search left root at key_5") &&
               expect(!test_search("not_found2"), "Expect we can search for non-existent node") &&
               expect(test_position("key_5"), "Expect that search left root at key_5") &&
               expect(test_search("key_6"), "Expect we can search for key_6 (exists)") &&
               expect(test_position("key_6"), "Expect that search left root at key_6");
    }
};