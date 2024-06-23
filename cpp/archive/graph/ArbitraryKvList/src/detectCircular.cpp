/**
 * @name ArbitraryKvList/src/detectCircular.cpp
 * @copyright (c) 2022 Sam Caldwell. All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef detectCircular_H
#define detectCircular_H

/**
 * @name detectCircular
 * @brief This file defines a class which creates/maintains
 *        an ArbitraryKeyValue LinkedList.
 *
 * @param n ArbitraryKvNode pointer
 * @return bool
 */
inline bool detectCircular(ArbitraryKvNode *n) {
    return (
            (n) &&
            (
                    (n->left == n) ||
                    (n->right == n) ||
                    (n->left == n->right)
            )
    );
}

#endif