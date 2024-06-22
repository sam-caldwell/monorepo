/**
 * @name ArbitraryKvList/src/DumpKeys.cpp
 * @copyright (c) 2022 Sam Caldwell. All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name dumpKeys
 * @brief Dump List keys as a comma-delimited list.
 *
 * @return string
 */
string ArbitraryKvList::dumpKeys() {
    string result = "";
    if (root) {
        ArbitraryKvNode *curr;
        curr = root;
        while (curr->left) curr = curr->left;
        for (int i = 0; i < count(); i++) {
            result += curr->key() + ",";
            curr = curr->right;
        }
    }
    return result;
}
