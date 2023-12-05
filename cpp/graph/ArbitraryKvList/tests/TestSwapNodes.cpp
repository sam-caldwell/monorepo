/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.

 */
#include <iostream>
#include <stdlib.h>
#include <time.h>
#include "projects/graph/ArbitraryKvList/src/ArbitraryKvList.h"

class TestSwapNodes : public TestBase {
private:
    ArbitraryKvNode *root;
    ArbitraryKvNode *node;
public:
    TestSwapNodes(string n) { name = n; }

    ~TestSwapNodes() {
        delete root;
    }

    bool test_setup() {
        debug(name + "::test_setup() starting");
        for (int i = 0; i < 10; i++) {
            string key = "key_" + to_string(i);
            if (root) {
                root->left = new ArbitraryKvNode(key, true);
                root->left->right = root;
                root = root->left;
            } else {
                root = new ArbitraryKvNode(key, true);
            }
            debug(name + "::test_setup() key=" + key);
        }
        return expect(root->key() == "key_9", ">> Expect current node is key_9.");
    }

    bool test_adjacent_swap() {
        root = resetNodeLeft(root);
        if (root->key() != "key_9") return false;
        node = root->right;
        if (node->key() != "key_8") return false;
        if (!swapNodes(root, node)) return false;
        return ((root->key() != "key_8") && (node->key() != "key_9"));
    }

    bool test_separate_swap() {
        if (root->key() != "key_9") return false;
        if (node->key() != "key_8") return false;
        node = root->right;
        if (node->key() != "key_7") return false;
        if (!swapNodes(root, node)) return false;
        return ((root->key() == "key_9") && (node->key() == "key_7"));
    }

    bool test_end_swap_left() {
        root = resetNodeLeft(root);
        node = resetNodeRight(node);
        if (root->key() != "key_8") return false;
        if (node->key() != "key_0") return false;
        if (!swapNodes(root, node)) return false;
        return ((root->key() == "key_8") && (node->key() == "key_0"));
    }

    bool test_end_swap_right() {
        root = resetNodeLeft(root);
        node = resetNodeRight(node);
        if (root->key() != "key_0") return false;
        if (node->key() != "key_8") return false;
        if (!swapNodes(root, node)) return false;
        return ((root->key() == "key_0") && (node->key() == "key_8"));
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect list setup will work") &&
               expect(test_adjacent_swap(), "Expect swap of adjacent nodes will work.") &&
               expect(test_separate_swap(), "Expect swap of separate nodes will work.") &&
               expect(test_end_swap_left(), "Expect swap of node at left end will work.") &&
               expect(test_end_swap_right(), "Expect swap of node at right end will work.");
    }
};