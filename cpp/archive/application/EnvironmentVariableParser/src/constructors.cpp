/**
 * @name EnvironmentVariableParser/src/constructors.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name Class constructor.
 * @brief Initialize our environment variable parser with a
 *        mapping object by reference.
 *
 * @details This class constructor will take a reference to the ConfigStateMap
 *          and Configuration objects and store them for internal use.
 *
 * @param *_cfg Configuration pointer
 * @param *_map ConfigStateMap pointer
 */
EnvironmentVariableParser::EnvironmentVariableParser(Configuration *_cfg, ConfigStateMap *_map) {
    if (_map)
        map = _map;
    else
        throw out_of_range("null map not allowed");
    if (_cfg)
        cfg = _cfg;
    else
        throw out_of_range("null configuration not allowed");
};
