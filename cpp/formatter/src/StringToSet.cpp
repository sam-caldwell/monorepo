/**
 * @name formatter/src/StringToSet.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Formatter_StringToSet_H
#define Formatter_StringToSet_H

#include <set>
#include <iostream>
#include <sstream>
#include<string>

/**
 * @name stringToSet
 * @brief Consume a delimited string and return a set of parsed elements.
 *
 * @param source string
 * @param delimiter char - delimiter - default: ','
 * @return
 */

set <string> stringToSet(string source, char delimiter = ',') {
    set <string> results;
    stringstream s(source);
    string token;
    while (getline(s, token, delimiter))
        results.insert(token);
    return results;
}

#endif /* Formatter_StringToSet_H */