/**
 * ArbitraryKvNode/src/ArbitraryKvNode.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 */
ArbitraryKvNode::ArbitraryKvNode() {
    left = right = NULL;
    ArbitraryKvNode::ArbitraryKeyValue();
}

ArbitraryKvNode::ArbitraryKvNode(string strKey, bool value) {
    left = right = NULL;
    key(strKey);
    setBool(value);
}

ArbitraryKvNode::ArbitraryKvNode(string strKey, double value) {
    left = right = NULL;
    key(strKey);
    setDouble(value);
}

ArbitraryKvNode::ArbitraryKvNode(string strKey, float value) {
    left = right = NULL;
    key(strKey);
    setFloat(value);
}

ArbitraryKvNode::ArbitraryKvNode(string strKey, int value) {
    left = right = NULL;
    key(strKey);
    setInt(value);
}

ArbitraryKvNode::ArbitraryKvNode(string strKey, string value) {
    left = right = NULL;
    key(strKey);
    setString(value);
}

ArbitraryKvNode::ArbitraryKvNode(string strKey, uint value) {
    left = right = NULL;
    key(strKey);
    setUint(value);
}

ArbitraryKvNode::~ArbitraryKvNode() {
    delete left;
    delete right;
}