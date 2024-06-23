/**
 * ArbitraryKvNode/src/ArbitraryKvNode.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 */

#ifndef ARBITRARY_KV_NODE_H
#define ARBITRARY_KV_NODE_H
#include <regex>
#include <list>
#include <stdexcept>
#include "../../../../types/unsigned_int.h"
#include "../../ArbitraryKeyValue/src/ArbitraryKeyValue.h"

class ArbitraryKvNode : public ArbitraryKeyValue {
public:
    ArbitraryKvNode *left;
    ArbitraryKvNode *right;

    ArbitraryKvNode();

    ArbitraryKvNode(string strKey, bool value);

    ArbitraryKvNode(string strKey, double value);

    ArbitraryKvNode(string strKey, float value);

    ArbitraryKvNode(string strKey, int value);

    ArbitraryKvNode(string strKey, string value);

    ArbitraryKvNode(string strKey, uint value);

    ~ArbitraryKvNode();
};

#include "ArbitraryKvNode.cpp"

#endif //ARBITRARY_KV_NODE_H