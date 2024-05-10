/**
 * Parsers/yaml/src/Parsers.h/routeHeader.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name routeHeader
 * @brief detect the yaml --- header and store as root token
 *
 * @throws runtime_error if header is missing or not the first detected token.
 *
 * @return bool
 */
inline bool Parser::routeHeader() {
    /**
     * Check for (---) Yaml header.
     *     We expect a Yaml file to start with --- (excluding any comments which may appear
     *     above the --- header.  Most linters consider --- to be a warning, but our parser
     *     requires this header....
     */
    if ((std::regex_match(line, std::regex(PATTERN_HEADER))) && (tokens.size() == 0)) {
        expectedToken = Any;
        return true;
    }
    throw runtime_error("missing --- header at the top of YAML file.");
}
