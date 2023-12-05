/**
 * @name CommandLineParser/src/parse.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name parse
 * @brief execute the parser functionality.
 *
 * @return bool
 */
bool CommandLineParser::parse(int argc, char *argv[]) {
    bool found;
    ValueTypes type;
    string value;
    string thisArg;
    for (auto key = map->getLhsBegin(); key != map->getLhsEnd(); ++key) {
        thisArg = key->first;
        found = false;
        value = "";
        try {
            type = map->getType(thisArg);
            switch (type) {
                case Bool:
                    value = "false";
                    break;
                case Double:
                case Float:
                case Int:
                case Uint:
                    value = "0";
                case String:
                    value = "";
                    break;
                default:
                    throw runtime_error("Unknown/unsupported commandline argument type.");
            }
            /**
             * Scan to see if a value is passed from cli.
             * skip argv[0] by doing ++i not i++
             */
            for (int i = 1; i < argc; i++) {
                if (thisArg.compare(argv[i]) == 0) {
                    found = true;
                    if (type == Bool) {
                        value = "true";
                    } else {
                        if (i < argc) {
                            value = argv[++i];
                        }
                    }
                    break;
                }
            }/*end of for*/
            /**
             * check to see if we are missing any required parameter
             */
            if (!found && map->getRequired(thisArg))
                throw runtime_error("Missing required argument:" + thisArg);
            /**
             * set the key and value (default or from cli)
             */
            if (found && map->rhsExists(thisArg))
                continue; //Do not overwrite with defaults if already exists.
            else
                cfg->loadData(map, thisArg, value);
        } catch (out_of_range) {
            throw runtime_error("Missing argument: " + thisArg);
        } catch (exception &e) {
            throw runtime_error("Error parsing command-line argument (" + thisArg + ") " + string(e.what()));
        }
    }/* end of for */
    return true;
}