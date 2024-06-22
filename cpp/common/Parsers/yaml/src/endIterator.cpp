/**
 * Parsers/yaml/src/Parsers.h/endIterator.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name endIterator
 * @brief return whether the tokens iterator is at the end of the map
 *
 * @return bool
 */
inline bool Parser::endIterator() {
    return tokenIterator == tokens.end();
}
