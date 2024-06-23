/**
 * Parsers/yaml/src/Parsers.h/initIterator.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
* @name initIterator
* @brief reset/init the tokens iterator to the start of the map.
*/
inline void Parser::initIterator() {
    tokenIterator = tokens.begin();
}
