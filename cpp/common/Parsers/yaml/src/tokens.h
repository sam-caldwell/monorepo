/**
 * Parsers/yaml/src/Parsers.h/tokens.cpp
 *
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef Parser_Tokens_H
#define Parser_Tokens_H

const char SPACE = ' ';

const string PATTERN_HEADER = "^---$";

const string PATTERN_TAG_MAP = "^( )*[a-zA-Z0-9][a-zA-Z0-9 ]*:$";
const string PATTERN_TAG_KEYVALUE = "^( )*[a-zA-Z0-9][a-zA-Z0-9 ]*:[ ].+$";
const string PATTERN_TAG_MULTILINE = "^( )*([-]( )){0,1}[a-zA-Z0-9][a-zA-Z0-9]*:( )\\|$";

const string PATTERN_SEQUENCE_ELEMENTS = "^( )*[-]( ).+$";
const string PATTERN_SEQUENCE_MAP_ITEM = "^( )*[-]( )[a-zA-Z0-9][a-zA-Z0-9]*:[ ]*$";
const string PATTERN_SEQUENCE_KV_ITEM = "^( )*[-]( )[a-zA-Z0-9][a-zA-Z0-9]*:.+$";

#endif /* Parser_Tokens_H */