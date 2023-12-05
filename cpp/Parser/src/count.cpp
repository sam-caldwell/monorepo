/**
 * @name Parser/src/count.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name count
 * @brief return the token count.
 *
 * @return unsigned
 */
inline unsigned Parser::count() {
    return tokens.size();
}
