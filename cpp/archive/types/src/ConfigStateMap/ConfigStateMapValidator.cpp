/**
 * @name types/src/ConfigStateMapValidator.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "ConfigStateMapNullValidator.cpp"

/**
 * @name ConfigStateMapValidator
 * @brief function pointer to a validator function which will
 *        evaluate the given string input and return a boolean
 *        response.
 *
 * @param input string
 * @return bool (operation outcome)
 */
typedef bool (*ConfigStateMapValidator)(string input);
