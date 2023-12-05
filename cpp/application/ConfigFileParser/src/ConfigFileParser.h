/**
 * @name ConfigFileParser/src/ConfigFileParser.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 * @brief This file defines a configuration file parser tool.
 *
 * @details The ConfigFileParser uses a shared-tree from Configuration
 *          project (similar to the EnvironmentVariableParser and
 *          CommandLineParser) and should be pluggable to support
 *          multiple file formats.
 *
 * @note This first version will support only YAML but the pattern
 *       should be there for expansion to JSON, etc.
 *
 * @todo Support JSON
 *
 */
#ifndef ConfigFileParser_H
#define ConfigFileParser_H

/**
 * Use the following preprocessor directive to disable
 * debugger functionality.
 */
#ifdef TESTER_DEBUGGING
#undef TESTER_DEBUGGING
#endif /**TESTER_DEBUGGING*/

/**
 * Use the following preprocessor directive to enable
 * debugger functionality.
 */


#include "projects/common/types/unsigned_int.h"
#include "projects/common/exceptions/exceptions.h"
#include "projects/common/types/ConfigStateMap.h"
#include "projects/application/Configuration/src/constants.h"
#include "projects/application/Configuration/src/Configuration.h"

class ConfigFileParser {
public:
    ConfigFileParser() = delete; //disable default constructor
    /**
     * @name Class constructor
     * @brief Initialize a the file parser with a given map
     *        and Configuration object to determine how the
     *        data will be parsed and parse the given file.
     *
     * @param *_cfg Configuration
     * @param *_map ConfigStateMap
     */
    ConfigFileParser(Configuration *_cfg, ConfigStateMap *_map);

    /**
     * @name Class destructor
     * @brief tear down the system.
     */
    ~ConfigFileParser();

    /**
     * @name parse
     * @brief Using the configured ConfigStateMap, get the mapped
     *        config file objects, validate them and store them
     *        in the mapped internal objects.
     *
     * @param *filename string by reference
     * @return bool
     */
    bool parse(string *fileName);

private:

    /**
     * @name addMap
     * @brief configure the object with an existing ConfigStateMap by reference.
     *
     * @throws out_of_range exception if null value passed as input
     *
     * @param *p ConfigStateMap pointer
     */
//    inline void addMap(ConfigStateMap *p);

    /**
     * @name addCfg
     * @brief configure the object with an existing Configuration by reference.
     *
     * @throws out_of_range expectation if null value passed as input
     *
     * @param *p Configuration pointer
     */
//    inline void addCfg(Configuration *p);

    ConfigStateMap *map;
    Configuration *cfg;
};/*end of ConfigFileParser*/

#include "constructor.cpp"
#include "destructor.cpp"
#include "parse.cpp"

#endif /* ConfigFileParser_H */