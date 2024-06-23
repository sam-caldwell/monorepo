/**
 * @name ArbitraryKvBtree/src/drawTreeRecursively.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * Recursively map out the binary tree.
 * @return string
 */
string drawTreeRecursively(ArbitraryKvBtNode *n, string id = "root") {
    string result = "";
    if (n) {
        result += "\n\t\t\t[" + id + "]: " + n->key() + "\n" +
                  drawTreeRecursively(n->left, id + ".L") +
                  drawTreeRecursively(n->right, id + ".R");
    } else {
        result = "\n\t\t\t[" + id + "]: NULL";
    }
    return result;
}
