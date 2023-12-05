/**
 * @name EnvironmentVariableParser/src/getEnvVar.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getEnvVar
 * @brief Get the environment variable value associated with the given name (string)
 *        then validate the value against the appropriate validator function for the
 *        environment variable name and store it in the ArbitraryKvList.
 *
 * @throws runtime_error exception if a required environment variable is not set in the op/sys.
 *
 * @param lhsName string (environment variable name)
 * @return bool
 */
inline bool EnvironmentVariableParser::getEnvVar(string lhsName) {
    /**
     * If we do not have an environment variable ConfigStateMap, then return.
     * This is a noop and it means that (a) we will not create the shared-tree
     * in the Configuration base class, (b) we will not execute any parser code,
     * and we will not load any environment variables.  We will save memory and
     * time.
     */
    if ((map == NULL) || (map && (map->count() == 0))) return false;

    /**
     * Get the raw environment variable value as a string.
     */
    string value = getEnv(&lhsName);
    /**
     * Validate the lhsName (environment variable)
     * and its value using the validator function
     * configured in the map.
     */
    return cfg->loadData(map, lhsName, value);
}
