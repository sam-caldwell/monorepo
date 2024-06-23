/**
 * @name formatter/src/setToString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <set>
#include <sstream>
#include "rtrimChar.cpp"

using namespace std;

/**
 * @name setToString
 * @brief serialize a set of tags to create a comma-delimited, quoted list of tags and return the same as a string.
 * @param tags
 * @return string
 */
string setToString(set <string> *data) {

    std::stringstream buffer;

    for (set<string>::iterator it = data->begin(); it != data->end(); ++it)
        buffer << "\"" << *it << "\",";

    return Strings::rtrimChar(buffer.str(), ',');
}
