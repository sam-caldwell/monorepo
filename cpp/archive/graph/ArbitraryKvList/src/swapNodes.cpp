/**
 * ArbitraryKvList/src/swapNodes.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef swapNodes_H
#define swapNodes_H

#include "isAdjacent.cpp"
#include "detectCircular.cpp"

/**
 * @name swapNodes
 * @brief Swap two nodes by pointer.  NOTE: not thread-safe.  no locking.
 *
 * @param A ArbitraryKvNode pointer
 * @param B ArbitraryKvNode pointer
 * @return bool
 */
bool swapNodes(ArbitraryKvNode *A, ArbitraryKvNode *B) {
    bool result = false;
    if (A && B) {
        if (isAdjacent(A, B)) {
            ArbitraryKvNode *LHS = ((A->left == B) && (B->right = A)) ? A : B;
            ArbitraryKvNode *RHS = ((A->left == B) && (B->right = A)) ? B : A;

            if (RHS->left) RHS->left->right = LHS;
            if (LHS->right) LHS->right->left = RHS;

            LHS->left = RHS->left;
            RHS->right = LHS->right;
            LHS->right = RHS;
            RHS->left = LHS;

        } else {
            ArbitraryKvNode *tempA_left = A->left;
            ArbitraryKvNode *tempA_right = A->right;
            ArbitraryKvNode *tempB_left = B->left;
            ArbitraryKvNode *tempB_right = B->right;

            A->left = tempB_left;
            A->right = tempB_right;

            B->left = tempA_left;
            B->right = tempA_right;

            if (A->left) A->left->right = A;
            if (A->right) A->right->left = B;
            if (B->left) B->left->right = B;
            if (B->right) B->right->left = B;

        }
        result = !(detectCircular(A) || detectCircular(B)); //Fail on circular.

    } else {
        result = false;
    }
    return result;
}

#endif /**swapNodes_H*/