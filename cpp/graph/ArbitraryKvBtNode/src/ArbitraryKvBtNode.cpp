/**
 * ArbitraryKvBtNode/src/ArbitraryKvBtNode.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */

ArbitraryKvBtNode::ArbitraryKvBtNode(){
    left = right = root = NULL;
}

ArbitraryKvBtNode::ArbitraryKvBtNode(string strKey, bool value) {
    left = right = root = NULL;
    key(strKey);
    setBool(value);
}

ArbitraryKvBtNode::ArbitraryKvBtNode(string strKey, double value) {
    left = right = root = NULL;
    key(strKey);
    setDouble(value);
}

ArbitraryKvBtNode::ArbitraryKvBtNode(string strKey, float value) {
    left = right = root = NULL;
    key(strKey);
    setFloat(value);
}

ArbitraryKvBtNode::ArbitraryKvBtNode(string strKey, int value) {
    left = right = root = NULL;
    key(strKey);
    setInt(value);
}

ArbitraryKvBtNode::ArbitraryKvBtNode(string strKey, string value) {
    left = right = root = NULL;
    key(strKey);
    setString(value);
}

ArbitraryKvBtNode::ArbitraryKvBtNode(string strKey, uint value) {
    left = right = root = NULL;
    key(strKey);
    setUint(value);
}

ArbitraryKvBtNode::~ArbitraryKvBtNode() {
    delete left;
    delete right;
    delete root;
}
