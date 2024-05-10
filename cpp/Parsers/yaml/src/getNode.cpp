/**
 * Parsers/yaml/src/Parsers.h/getNode.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getNode
 * @brief read Parsers file into internal state.
 *
 * @return Node* Node pointer
 */
Node *Parser::getNode() {
    return tokens;
}
