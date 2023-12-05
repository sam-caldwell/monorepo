/**
 * @name Application/src/configure.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef Application_configure_cpp
#define Application_configure_cpp

#include "projects/common/types/ConfigStateMap.h"
#include "projects/application/Configuration/src/Configuration.h"
#include "projects/application/EnvironmentVariableParser/src/EnvironmentVariableParser.h"
#include "projects/application/CommandLineParser/src/CommandLineParser.h"
#include "projects/application/ConfigFileParser/src/ConfigFileParser.h"

/**
 * @name configure
 * @brief Configure the application using Environment variables, command-line arguments and any optional
 *        YAML configuration file.
 *
 * @param *cfg - Configuration pointer
 * @param *map - ConfigStateMap pointer
 * @param argc - int - count of arguments from commandline
 * @param *argv[] - reference to an array of argument strings.
 * @return bool
 */
bool configure(Configuration *cfg, ConfigStateMap *map, int argc, char *argv[]) {
    if (cfg) {
        if (map) {
            if (argv && (argc > 0)) {

                EnvironmentVariableParser env(cfg, map);

                CommandLineParser cli(cfg, map, argc, argv);

                ConfigFileParser file(cfg, map);

                return true;
            } else {
                throw runtime_error("cli args missing or malformed");
            }
        } else {
            throw runtime_error("ConfigStateMap object is null");
        }
    } else {
        throw runtime_error("Configuration object is null");
    }
}

#endif /* Application_configure_cpp */