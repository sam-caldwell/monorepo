/**
 * @name Parser/src/getToken.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getToken
 * @brief given a token name (key), return the associated value (string).
 *
 * @throws out_of_range exception is thrown if no such key is found.
 *
 * @param key string
 * @return string
 */
string Parser::getToken(string key) {
    try {
        return tokens.at(key);
    } catch (out_of_range &e) {
        throw out_of_range("key (" + key + ") is missing.");
    }
}