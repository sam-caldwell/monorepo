/**
* @name ArbitraryKvList/src/MoveLeft.cpp
* @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
* @author Sam Caldwell <mail@samcaldwell.net>
*/

/**
 * @name moveLeft
 * @brief Move one node left
 *
 * @return bool
 */
inline bool ArbitraryKvList::moveLeft() {
    if (root && root->left)
        return root = root->left;
    return false;
}