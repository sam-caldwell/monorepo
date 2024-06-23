/**
 * @name ArbitraryKvList/src/Remove.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name remove
 * @brief Delete the current node.
 *
 * @return bool
 */
bool ArbitraryKvList::remove() {
    if (!lock->check()) lock->on();
    bool result = false;
    if (root) {
        //ToDo: add locking mechanism

        ArbitraryKvNode *temp_left;
        ArbitraryKvNode *temp_right;

        temp_left = root->left;
        temp_right = root->right;

        if (temp_left) temp_left->right = temp_right;
        if (temp_right) temp_right->left = temp_left;

        root->left = NULL; //unlink node;
        root->right = NULL;//unlink node;

        { //Create a safe scope.
            ArbitraryKvNode *temp;
            temp = root; //temp points to our to-be-deleted node.
            if (temp) delete temp;
        }
        root = (temp_left) ? temp_left : temp_right; //point root back to list

        _count--;
        result = true;
    }
    lock->off();
    return result;
}