/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "projects/graph/ArbitraryKvList/src/ArbitraryKvList.h"

class TestLinkListRemoveKey : public TestBase {
private:
    ArbitraryKvList *root;
public:
    TestLinkListRemoveKey(string n) { name = n; }

    ~TestLinkListRemoveKey() { delete root; }

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
               expect(root->key() == "key_0", ">> Expect current node is key_0. found:"+root->key());
    }

    bool test_remove_key() {
        debug(name + "::test_remove_key()");
        for (int i = 0; i < 10; i++) {
            string key = "key_" + to_string(i);
            debug(name + "::test_remove_key(" + key + ") deleting. (keys:" + root->dumpKeys() + ")");
            if (!root->exists(key)) {
                error("key (" + key + ") should have existed");
                return false;
            }
            debug(name + "::test_remove_key() deleting key ("+key+")");
            if (!root->remove(key)) {
                error("key(" + key + ") failed to be deleted.");
                return false;
            }
            debug(name + "::test_remove_key() verifying key deletion ("+key+")");
            if (root) debug(name + "::test_remove_key() non-null root after deleting ("+key+")");
            if (root)
                if (root->exists(key)) {
                    error("Deleted node still exists (" + key + ")");
                    return false;
                }
            debug(name + "::test_remove_key(" + key + ") delete done. (keys:" + root->dumpKeys() + ")");
        }
        debug(name + "::test_remove_key() count is " + to_string(root->count()));
        return (root->count() == 0);
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect we can create initialize a default state.") &&
               expect(test_remove_key(), "Expect that we can remove all keys individually and realize count==0");
    }
};