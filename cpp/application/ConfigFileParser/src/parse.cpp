/**
 * @name ConfigFileParser/src/parse.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef ConfigFileParser_parse_H
#define ConfigFileParser_parse_H

#include "projects/common/Parser/src/Parser.h"

/**
 * @name parse
 * @brief Parse the configured file and store results in the shared tree
 *
 * @throws exception if there are problems parsing the data.  See projects/Parsers
 * @throws exception if there are problems accessing the source file.  See projects/Parsers.
 *
 * @param *filename string by reference
 * @return bool
 */
inline bool ConfigFileParser::parse(string *filename) {
    Parser parser(*filename, Yaml);
    parser.parse(); //The parser now has key:value tokens where our key is dot-delimited.
    parser.initIterator();
    while (!parser.endIterator()) {
        pair <string, string> node = parser.next();
        string lhs = node.first;
        string value = node.second;
        cfg->loadData(map, lhs, value);
    }
    return true;
}

#endif /* ConfigFileParser_parse_H */