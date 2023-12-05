/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include <stdlib.h>
#include <time.h>
#include "projects/graph/ArbitraryKvList/src/ArbitraryKvList.h"

class TestResetNodeRight : public TestBase {
private:
    ArbitraryKvNode *root;
public:
    TestResetNodeRight(string n) { name = n; }

    ~TestResetNodeRight() {
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

    bool test1() {
        root = resetNodeRight(root);
        if (!root->left) return false; //expect no left node.
        if (root->right) return false; //expect there is a right node.
        return (root && root->key() == "key_0");
    }

    bool test2() {
        do { root = (root->right) ? root->right : root; } while (root->right);
        if (!root->left) return false; //expect no left node.
        if (root->right) return false; //expect there is a right node.
        root = resetNodeRight(root);
        return (root && root->key() == "key_0");
    }

    bool main() {
        return expect(test_setup(), "Expect test setup ok.") &&
               expect(test1(), "Expect that we are at left and move to right will do nothing.") &&
               expect(test2(), "Expect that we are at left and move to right will do nothing.");
    }

};