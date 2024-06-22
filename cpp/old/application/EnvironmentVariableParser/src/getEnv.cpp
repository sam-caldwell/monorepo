/**
 * @name EnvironmentVariableParser/src/getEnv.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name getEnv
 * @brief return an environment variable and throw exception if
 *        required and the env var is not found.
 *
 * @warning throws runtime_error exception
 *
 * @param n string
 * @return string
 */
string EnvironmentVariableParser::getEnv(string *n) {

    string result = "";

    try {

        result = Environment::get(*n);

    } catch (runtime_error e) {

        if (this->map->getRequired(*n))
            throw runtime_error(e.what());

    }

    return result;

}
