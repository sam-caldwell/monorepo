/**
 * @name CommandLineParser/src/constructor.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name Class constructor
 * @brief Initialize a the command-line parser with a given map
 *        and Configuration object to determine how the data will
 *        be parsed and parse the given file.
 *
 * @param *_cfg Configuration
 * @param *_map ConfigStateMap
 * @param argc int
 * @param *argv[] pointer to argument list.
 */
CommandLineParser::CommandLineParser(Configuration *_cfg, ConfigStateMap *_map, int argc, char *argv[]) {
    if (_map)
        map = _map;
    else
        throw runtime_error("null map not allowed");

    if (_cfg)
        cfg = _cfg;
    else
        throw runtime_error("null configuration not allowed");

    if (argc < 0) throw runtime_error("Programming error: CommandLineParser expects argc >0");
    if (argv == NULL) throw runtime_error("Programming error: CommandLineParser");

    if (!parse(argc, argv)) throw runtime_error("command-line argument parsing failed.");

    try {
        //parse();
    } catch (out_of_range &e) {
        throw out_of_range("CommandLineParser encountered error:" + string(e.what()));
    } catch (exception &e) {
        throw runtime_error("CommandLineParser encountered error: " + string(e.what()));
    }
}