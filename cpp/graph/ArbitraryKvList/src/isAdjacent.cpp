/**
 * @name ArbitraryKvList/src/isAdjacent.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef isAdjacent_H
#define isAdjacent_H

/**
 * @name isAdjacent
 * @brief Determine if two nodes are logically adjacent
 *        (Note: this does not evaluate physical position)
 *
 * @param *A ArbitraryKvNode pointer
 * @param *B ArbitraryKvNode pointer
 * @return bool
 */
bool isAdjacent(ArbitraryKvNode *A, ArbitraryKvNode *B) {
    return (
            A &&
            B && (
                    (
                            (A->left) &&
                            (B->right) &&
                            (A->left == B) &&
                            (B->right == A)
                    ) || (
                            (A->right) &&
                            (B->left) &&
                            (A->right == B) &&
                            (B->left == A)
                    )
            ));
}

#endif /* isAdjacent_H */