/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include <stdlib.h>
#include <time.h>
#include "projects/graph/ArbitraryKvList/src/ArbitraryKvList.h"

class TestIsAdjacent : public TestBase {
private:
    ArbitraryKvNode *node;
public:
    TestIsAdjacent(string n) { name = n; }

    ~TestIsAdjacent() {
        delete node;
    }

    bool test_setup(){
        node = new ArbitraryKvNode("key_1", 1);

        node->left = new ArbitraryKvNode("key_2", 2);
        node->left->right = node;

        node->right = new ArbitraryKvNode("key_3", 0);
        node->right->left = node;

        cout << "node [" << node << "]("<< node->key() << ")" << endl
             << "node->left [" << node->left << "]("<< node->left->key() << ")" << endl
             << "node->right [" << node->right << "]("<< node->right->key() << ")" << endl;

        return true;
    }

    bool test1() {
        return isAdjacent(node, node->left);
    }

    bool test2(){
        return isAdjacent(node,node->right);
    }

    bool test3() {
        return !isAdjacent(node->right, node->left);
    }

    bool main() {
        return expect(test_setup(), "Expect test setup ok.") &&
               expect(test1(), "Expect node adjacent to node->left") &&
               expect(test2(), "Expect node adjacent to node->right") &&
               expect(test3(), "Expect that node->right is NOT adjacent to node->left");
    }

};