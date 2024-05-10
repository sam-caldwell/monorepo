/**
 * Parsers/yaml/src/Parsers.h/next.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name next
 * @brief return the next key:value pair from the map using an iterator.
 *
 * @throws out_of_range if the iterator is at the end of the map.
 *
 * @return pair<string,string>
 */
pair <string, string> Parser::next() {
    if (endIterator()) {
        throw out_of_range("no tokens.");
    } else {
        pair <string, string> v(tokenIterator->first, tokenIterator->second);
        tokenIterator++;
        return v;
    }
}
