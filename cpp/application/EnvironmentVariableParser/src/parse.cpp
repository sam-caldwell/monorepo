/**
 * @name EnvironmentVariableParser/src/parse.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name parse
 * @brief Using the configured ConfigStateMap, get the mapped
 *        environment variables, then validate them and if any error
 *        occurs, throw an exception with the mapped error message.
 *        Otherwise, store an ArbitraryValueKvList of environment
 *        variables for consumption.
 *
 * @throws out_of_range exception if map is not initialized.
 *
 * @return bool
 */
bool EnvironmentVariableParser::parse() {

    if (map == NULL) throw out_of_range("environment variable map must be initialized");

    for (auto i = map->getLhsBegin(); i != map->getLhsEnd(); ++i) {
        /**
         * lhs is the name of an environment variable (i->first).
         *
         * getEnvVar(lhs) will read environment variables,
         * parse the data then store the value to the internal rhs
         * state.
         */
        getEnvVar(i->first);
    }
    return true;
}