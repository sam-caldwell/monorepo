/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "projects/graph/ArbitraryKvBtNode/src/ArbitraryKvBtNode.h"

class TestArbitraryKvBtNodeCreate : public TestBase {
public:
    TestArbitraryKvBtNodeCreate(string n) { name = n; }

    bool test_node_create_default() {
        debug(name + "::test_node_create_default() starting");
        ArbitraryKvBtNode node;
        return expect(node.left == NULL, ">> expect that node.left is NULL") &&
               expect(node.right == NULL, ">> expect that node.right is NULL") &&
               expect(node.root == NULL, ">> expect that node.root is NULL");
    }

    bool test_node_create_bool(bool n) {
        debug(name + "::test_node_create_default() starting (n=" + to_string(n) + ")");
        ArbitraryKvBtNode node("test_node_create_bool", n);
        return expect(node.key() == "test_node_create_bool", ">>Expect boolean node has expected key") &&
               expect(node.isType(Bool), ">>Expect node type is boolean.") &&
               expect(node.getBool() == n, ">>Expect value is matches boolean input") &&
               expect(node.left == NULL, ">> expect that node.left is NULL") &&
               expect(node.right == NULL, ">> expect that node.right is NULL") &&
               expect(node.root == NULL, ">> expect that node.root is NULL");
    }

    bool test_node_create_double(double n) {
        debug(name + "::test_node_create_double() starting (n=" + to_string(n) + ")");
        ArbitraryKvBtNode node("test_node_create_double", n);
        return expect(node.key() == "test_node_create_double", ">>Expect double node has expected key") &&
               expect(node.isType(Double), ">>Expect node type is double.") &&
               expect(node.getDouble() == n, ">>Expect value is matches double input") &&
               expect(node.left == NULL, ">> expect that node.left is NULL") &&
               expect(node.right == NULL, ">> expect that node.right is NULL") &&
               expect(node.root == NULL, ">> expect that node.root is NULL");
    }

    bool test_node_create_float(float n) {
        debug(name + "::test_node_create_float() starting (n=" + to_string(n) + ")");
        ArbitraryKvBtNode node("test_node_create_float", n);
        return expect(node.key() == "test_node_create_float", ">>Expect float node has expected key") &&
               expect(node.isType(Float), ">>Expect node type is float.") &&
               expect(node.getFloat() == n, ">>Expect value is matches float input") &&
               expect(node.left == NULL, ">> expect that node.left is NULL") &&
               expect(node.right == NULL, ">> expect that node.right is NULL") &&
               expect(node.root == NULL, ">> expect that node.root is NULL");
    }

    bool test_node_create_int(int n) {
        debug(name + "::test_node_create_int() starting (n=" + to_string(n) + ")");
        ArbitraryKvBtNode node("test_node_create_int", n);
        return expect(node.key() == "test_node_create_int", ">>Expect int node has expected key") &&
               expect(node.isType(Int), ">>Expect node type is int.") &&
               expect(node.getInt() == n, ">>Expect value is matches int input") &&
               expect(node.left == NULL, ">> expect that node.left is NULL") &&
               expect(node.right == NULL, ">> expect that node.right is NULL") &&
               expect(node.root == NULL, ">> expect that node.root is NULL");
    }

    bool test_node_create_string(string n) {
        debug(name + "::test_node_create_string() starting (n=" + n + ")");
        ArbitraryKvBtNode node("test_node_create_string", n);
        return expect(node.key() == "test_node_create_string", ">>Expect string node has expected key") &&
               expect(node.isType(String), ">>Expect node type is string.") &&
               expect(node.getString() == n, ">>Expect value is matches string input") &&
               expect(node.left == NULL, ">> expect that node.left is NULL") &&
               expect(node.right == NULL, ">> expect that node.right is NULL") &&
               expect(node.root == NULL, ">> expect that node.root is NULL");
    }

    bool test_node_create_uint(uint n) {
        debug(name + "::test_node_create_uint() starting (n=" + to_string(n) + ")");
        ArbitraryKvBtNode node("test_node_create_uint", n);
        return expect(node.key() == "test_node_create_uint", ">>Expect uint node has expected key") &&
               expect(node.isType(Uint), ">>Expect node type is uint.") &&
               expect(node.getUint() == n, ">>Expect value is matches uint input") &&
               expect(node.left == NULL, ">> expect that node.left is NULL") &&
               expect(node.right == NULL, ">> expect that node.right is NULL") &&
               expect(node.root == NULL, ">> expect that node.root is NULL");
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_node_create_default(), "Expect we can create a default node.") &&
               expect(test_node_create_bool(true), "Expect we can create a node with a boolean true value") &&
               expect(test_node_create_bool(false), "Expect we can create a node with a boolean false value") &&
               expect(test_node_create_double(3.141529), "Expect we can create a node with a double float.") &&
               expect(test_node_create_float(3.141529), "Expect we can create a node with a single float.") &&
               expect(test_node_create_int(3), "Expect we can create a node with an int.") &&
               expect(test_node_create_string("test_value"), "Expect we can create a node with an string.") &&
               expect(test_node_create_uint(3), "Expect we can create a node with an uint.");
    }
};