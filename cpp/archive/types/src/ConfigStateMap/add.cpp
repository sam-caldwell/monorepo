/**
 * @name types/src/add.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name add
 * @brief Add a new mapping record (string->string->ValueTypes)
 *
 * @param lhs string
 * @param rhs string
 * @param type ValueTypes
 * @param required bool (flag to determine if record is required).
 * @param validatorFunc ConfigStateMapValidator (default= NULL)
 * @return bool (result of operation)
 */
bool ConfigStateMap::add(string lhs,
                         string rhs,
                         ValueTypes type,
                         bool required = false,
                         bool allowDuplicateRhs = true,
                         ConfigStateMapValidator validatorFunc = ConfigStateMapNullValidator) {
    /**
     * validate the LHS entry
     */
    if (this->exists(lhs)) throw runtime_error("lhs already exists");
    /**
     * validate the RHS entry
     */
    if (!allowDuplicateRhs && this->rhsExists(rhs)) throw runtime_error("duplicate rhs detected");

    /**
     * validate the type entry.
     */
    if (type == Null) throw runtime_error("ValueType cannot be Null");

    StringMap.insert(pair<string, string>(lhs, rhs));
    ValueTypeMap.insert(pair<string, ValueTypes>(rhs, type));
    RequiredMap.insert(pair<string, bool>(lhs, required));
    ValidatorMap.insert(pair<string, ConfigStateMapValidator>(lhs, validatorFunc));
    return true;
}