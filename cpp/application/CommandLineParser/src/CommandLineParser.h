/**
 * @name CommandLineParser/src/CommandLineParser.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 * @brief This file defines the Command-line processor class
 *
 * @details The Command-Line (CLI) parser class takes a run-time
 *          definition of some command-line arguments expected
 *          by an application then processes the command-line
 *          args passed by the caller into the application
 *          program against this definition.  The result is
 *          a key-value configuration state.
 */
#ifndef CommandLineParser_H
#define CommandLineParser_H

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



#include "projects/common/formatter/formatter.h"
#include "projects/common/types/unsigned_int.h"
#include "projects/common/exceptions/exceptions.h"
#include "projects/common/types/ConfigStateMap.h"
#include "projects/graph/ArbitraryKvBtree/src/ArbitraryKvBtree.h"
#include "projects/application/Configuration/src/Configuration.h"

using namespace std;

class CommandLineParser {
public:
    CommandLineParser() = delete; //disable default constructor
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
    CommandLineParser(Configuration *_cfg, ConfigStateMap *_map, int argc, char *argv[]);

    /**
     * @name Class destructor
     * @brief tear down the system.
     */
    ~CommandLineParser();

    /**
     * @name parse
     * @brief execute the parser functionality.
     *
     * @return bool
     */
    bool parse(int argc, char *argv[]);

private:
    ConfigStateMap *map;
    Configuration *cfg;
};/*end of Configuration*/

#include "constructor.cpp"
#include "destructor.cpp"
#include "parse.cpp"

#endif /* CommandLineParser_H */