/**
 * @name EnvironmentVariableParser/src/EnvironmentVariableParser.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 * @brief This file defines the EnvironmentVariableParser class
 *
 * @details The EnvironmentVariableParser class will take a list
 *          of expected environment variables and their respective
 *          mappings to internal states then read and validate
 *          the system environment variables and return the mapped
 *          list of configuration state.
 */
#ifndef EnvironmentVariableParser_H
#define EnvironmentVariableParser_H
/**
 *
 */
#include <cstdlib>
#include <stdexcept>
#include <string>
#include "../../../../system/environment.h"
#include "../../../../formatter/formatter.h"
#include "../../../../types/unsigned_int.h"
#include "../../../../exceptions/exceptions.h"
#include "../../../../types/src/ConfigStateMap/ConfigStateMap.h"
#include "../../Configuration/src/Configuration.h"

using namespace std;

class EnvironmentVariableParser {
public:
    EnvironmentVariableParser() = delete; //disable default constructor
    /**
     * @name Class constructor.
     * @brief Initialize our environment variable parser with a
     *        mapping object by reference.
     * @details This class constructor will take a reference
     *          to the a ConfigStateMap associating an
     *          environment variable identifier with a
     *          configuration object and to a ValueType
     *          for that environment variable.
     *
     * @param *_cfg Configuration pointer
     * @param *_map ConfigStateMap pointer
     */
    EnvironmentVariableParser(Configuration *_cfg, ConfigStateMap *_map);

    /**
     * @name Class destructor
     * @brief tear down the system.
     */
    ~EnvironmentVariableParser();

    /**
     * @name addMap
     * @brief configure the object with an existing ConfigStateMap by reference.
     *
     * @throws out_of_range exception when null input encountered.
     *
     * @param *p ConfigStateMap pointer
     */
    inline void addMap(ConfigStateMap *_map);

    /**
     * @name parse
     * @brief Using the configured ConfigStateMap, get the mapped
     *        environment variables, then validate them and if any error
     *        occurs, throw an exception with the mapped error message.
     *        Otherwise, store an ArbitraryValueKvList of environment
     *        variables for consumption.
     *
     * @return bool
     */
    bool parse();

protected:
    /**
     * @name getEnvVar
     * @brief Get the environment variable value associated with the given name (string)
     *        then validate the value against the appropriate validator function for the
     *        environment variable name and store it in the ArbitraryKvList.
     *
     * @warning if validation fails or if a required environment variable value is missing
     *          an exception will be thrown.
     *
     * @param lhsName string (environment variable name)
     * @return bool
     */
    inline bool getEnvVar(string lhsName);

private:
    /**
     * @name getEnv
     * @brief wrapper for Environment::get() which will throw an exception
     *        if the environment variable is not found but is required.
     *
     * @param *s string pointer (env var name)
     * @return string
     */
    string getEnv(string *n);

    ConfigStateMap *map;
    Configuration *cfg;
}; /*end of EnvironmentVariableParser class*/

#include "constructors.cpp"
#include "destructors.cpp"
#include "getEnv.cpp"
#include "getEnvVar.cpp"
#include "parse.cpp"
#include "../../Configuration/src/parseBool.cpp"
#include "../../Configuration/src/parseDouble.cpp"
#include "../../Configuration/src/parseFloat.cpp"
#include "../../Configuration/src/parseInt.cpp"
#include "../../Configuration/src/parseUint.cpp"


#endif /* EnvironmentVariableParser_H */