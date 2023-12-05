/**
 * @name ConfigFileParser/src/constructor.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name Class constructor
 * @brief Initialize a blank-state config file parser with
 *        a given ConfigStateMap to determine how a data
 *        source will be parsed.
 *
 * @throws runtime_error exception if null map is passed during call.
 * @throws runtime_error exception if null cfg is passed during call.
 * @throws out_of_range exception if the ConfigFile is not defined in cfg prior to instantiation.
 *
 * @param *_cfg Configuration pointer
 * @param *_map ConfigStateMap
 */
ConfigFileParser::ConfigFileParser(Configuration *_cfg, ConfigStateMap *_map) {
    if (_map)
        map = _map;
    else
        throw runtime_error("null map not allowed");

    if (_cfg)
        cfg = _cfg;
    else
        throw runtime_error("null configuration not allowed");

    try {
        string fileName = cfg->getString(ConfigFile);
        parse(&fileName);
    } catch (out_of_range &e) {
        throw out_of_range(
                "ConfigFileParser encountered error:" + string(e.what()) +
                " be sure " + string(ConfigFile) + " is defined before instantiating " +
                " configFileParser.");
    } catch (exception &e) {
        throw runtime_error("ConfigFileParser encountered error: " + string(e.what()));
    }
};
