/**
 * @name ArbitraryKvList/src/MoveRight.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name moveRight
 * @brief Move one node right
 *
 * @return bool
 */
inline bool ArbitraryKvList::moveRight() {
    if (root && root->right)
        return root = root->right;
    return false;
}
