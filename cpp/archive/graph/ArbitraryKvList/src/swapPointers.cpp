/**
 * ArbitraryKvList/src/swapPointers.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name swapPointers
 * @brief swap two pointers (A,B) such that the result is (B,A)
 *
 * @param *A ArbitraryKvNode pointer
 * @param *B ArbitraryKvNode pointer
 */
void inline swapPointers(ArbitraryKvNode *A, ArbitraryKvNode *B) {
    ArbitraryKvNode *t = A;
    t = A;
    A = B;
    B = t;
}
