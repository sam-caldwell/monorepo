/**
 * Parsers/yaml/src/Parsers.h/TokenType.cpp
 *
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef Parser_TokenType_H
#define Parser_TokenType_H

enum TokenType {
    Any = 0,

    Header = 1,
    KeyValueTag = 2,
    MultiLineTag = 3,
    MultiLineValue = 4,
    SequenceItem = 5,
    SequenceKeyValue = 6,
    SequenceMap = 7,
    TagMap = 8,
    Value = 9,
};

string TokenTypeToString(TokenType t) {
    switch (t) {
        case Any:
            return "Any";
            break;
        case Header:
            return "Header";
            break;
        case KeyValueTag:
            return "KeyValueTag";
            break;
        case MultiLineTag:
            return "MultiLineTag";
            break;
        case MultiLineValue:
            return "MultiLineValue";
            break;
        case SequenceItem:
            return "SequenceItem";
            break;
        case SequenceKeyValue:
            return "SequenceKeyValue";
            break;
        case SequenceMap:
            return "SequenceMap";
            break;
        case TagMap:
            return "TagMap";
            break;
        case Value:
            return "Value";
            break;
        default:
            throw runtime_error("unexpected/unknown TokenType");
    }
}


#endif /* Parser_TokenType_H */