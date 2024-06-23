/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include <stdlib.h>
#include <time.h>
#include "../src/ArbitraryKvList.h"

class TestSwapPointers : public TestBase {
private:
    ArbitraryKvNode *A;
    ArbitraryKvNode *B;
public:
    TestSwapPointers(string n) { name = n; }

    ~TestSwapPointers() {
        delete A;
        delete B;
    }

    bool test_setup() {
        A = new ArbitraryKvNode("node1", 1);
        B = new ArbitraryKvNode("node2", 2);
        return (A && B);
    }

    bool test() {
        ArbitraryKvNode *A_init = A;
        ArbitraryKvNode *B_init = B;
        swapPointers(A, B);
        return ((A == A_init) && (B == B_init));
    }

    bool main() {
        return expect(test_setup(), "Expect node1 and node2 can be created easily.") &&
               expect(test(), "Expect node adjacent to node->right");
    }

};