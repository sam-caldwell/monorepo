/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include <stdlib.h>
#include <time.h>
#include "projects/common/formatter/src/intToString.cpp"
#include "projects/graph/ArbitraryKvList/src/ArbitraryKvList.h"

class TestLinkListSort : public TestBase {
private:
    ArbitraryKvList *root;
    string expected_asc;
    string expected_desc;
public:
    TestLinkListSort(string n) {
        name = n;
        srand((unsigned) time(NULL));
    }

    ~TestLinkListSort() { delete root; }

    bool test_setup() {
        debug(name + "::test_setup() starting");
        const int start = 0;
        const int end = 20;
        const int sz = end - start;
        root = new ArbitraryKvList(true, true);
        expected_asc = "";
        expected_desc = "";
        if (root->count() != 0) return false;
        for (int i = start; i < end; i++) {
            string key = "key_" + intToString(i, 4);
            expected_asc += key + ",";
            expected_desc = key + "," + expected_desc;
            root->insert(key, true);
            debug(name + "::test_setup() key=" + root->key());
        }
        debug(name + "::test_setup() "
              + to_string(sz) + "-nodes expected. created "
              + to_string(root->count()) + " nodes.");
        debug(name + "::test_setup()" + root->dumpKeys());

        return expect(root->count() == sz,
                      ">> Expect count: " + to_string(sz) + " actual:" + to_string(root->count())) &&
               expect(root->key() == "key_" + intToString(end - 1, 4),
                      ">> Expect current node is key_" + intToString(end - 1, 4));
    }

    bool test_sort(SortOrder Order) {
        string expected_outcome = "";
        switch (Order) {
            case Ascending: {
                expected_outcome = expected_asc;
                break;
            }
            case Descending: {
                expected_outcome = expected_desc;
                break;
            }
            default:
                return false;
        }
        return expect(root->sort(Order), ">> Expect sort(" + SortOrderToString(Order) + ") works.") &&
               expect(root->dumpKeys() == expected_outcome,
                      ">> Expect sorted output matches expected output. actual: " + root->dumpKeys() +
                      " expected: " + expected_outcome);
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect we can create a list.") &&
               expect(test_sort(Ascending), "Expect ascending sort works(1).") &&
               expect(test_sort(Descending), "Expect ascending sort works(1).");
    }
};