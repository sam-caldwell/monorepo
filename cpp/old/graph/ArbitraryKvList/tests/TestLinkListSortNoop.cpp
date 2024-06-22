/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include <stdlib.h>
#include <time.h>
#include "../src/ArbitraryKvList.h"

class TestLinkListSortNoop : public TestBase {
private:
    ArbitraryKvList *root;
public:
    TestLinkListSortNoop(string n) {
        name = n;
        srand((unsigned) time(NULL));
    }

    ~TestLinkListSortNoop() { delete root; }

    bool test_setup() {
        debug(name + "::test_setup() start");
        root = new ArbitraryKvList(true, true); //Accept duplicates
        if (root->count() != 0) return false;
        for (int i = 0; i < 10; i++) {
            string key = "key_" + to_string(i);
            root->insert(key, i);
            debug(name + "::test_setup() key=" + key);
        }
        debug(name + "::test_setup(): " + root->dumpKeys());
        return expect(root->count() == 10, ">> Expect root count is 10.");
    }

    bool test_noop() {
        debug(name + "::test_noop()");
        return expect(!root->sort(Undefined), ">> Expect noop on undefined sort order.");
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect we can create a list with sample set 1.") &&
               expect(test_noop(), "Expect noop works");
    }
};