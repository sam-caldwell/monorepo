/**
 * ArbitraryKvBtNode/src/ArbitraryKvBtNode.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */

#ifndef ArbitraryKvBtNode_H
#define ArbitraryKvBtNode_H

#include <regex>
#include <list>
#include <stdexcept>
#include "../../../../types/unsigned_int.h"
#include "../../ArbitraryKeyValue/src/ArbitraryKeyValue.h"

class ArbitraryKvBtNode : public ArbitraryKeyValue {
private:
protected:
public:
    ArbitraryKvBtNode *left;
    ArbitraryKvBtNode *right;
    ArbitraryKvBtNode *root;

    ArbitraryKvBtNode();

    ArbitraryKvBtNode(string strKey, bool value);

    ArbitraryKvBtNode(string strKey, double value);

    ArbitraryKvBtNode(string strKey, float value);

    ArbitraryKvBtNode(string strKey, int value);

    ArbitraryKvBtNode(string strKey, string value);

    ArbitraryKvBtNode(string strKey, uint value);

    ~ArbitraryKvBtNode();
};

#include "ArbitraryKvBtNode.cpp"

#endif /** ArbitraryKvBtNode_H */