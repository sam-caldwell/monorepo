/**
 * @name types/src/ConfigStateMap/ConfigStateMap.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 * This type creates a complex mapping of some arbitrary string (lhs)
 * to another arbitrary string (rhs) and then maps the two to several
 * shared characteristics (ValueTypes, error, validator).
 *
 * The validator function assumes a value is being evaluated against
 * the characteristics.  We assume this value is of type string and
 * it is compared against some arbitrary, consumer-provided logic.
 */

#include <map>
#include <regex>
#include "ConfigStateMapValidator.cpp"
#include "projects/common/types/ValueTypes.h"
#include "projects/common/formatter/formatter.h"

#define MAX_ERROR_LENGTH 1024

typedef string LhsString;
typedef string RhsString;

/**
 * @name ConfigStateMap class
 * @brief Map two strings and a ValueTypes value to one another.
 */
class ConfigStateMap {
public:
    /**
     * @name default constructor
     * @brief initialize an empty state.


     */
    ConfigStateMap() {};

    /**
     * @name class constructor [map]
     * @brief initialize the map with an existing map.
     * @param *_map ConfigStateMap pointer
     */
    ConfigStateMap(ConfigStateMap *_map);

    /**
     * @name class destructor
     * @brief destroy the class.
     */
    ~ConfigStateMap() {}

    /**
     * @name add
     * @brief Add a new mapping record (string->string->ValueTypes)
     *
     * @param lhs string
     * @param rhs string
     * @param type ValueTypes
     * @param required bool (flag to determine if record is required).
     * @param allowDuplicateRhs bool (flag to determine if RHS must be unique, default: true)
     * @param validatorFunc ConfigStateMapValidator (default= NULL)
     * @return bool (result of operation)
     */
    bool add(LhsString lhs,
             RhsString rhs,
             ValueTypes type,
             bool required,
             bool allowDuplicateRhs,
             ConfigStateMapValidator validatorFunc);

    /**
     * @name count
     * @brief Return the number of records in the map.
     *
     * @return int
     */
    int count();

    /**
     * @name erase
     * @brief Remove all mappings for the given LhsString.
     *
     * @param lhs string
     */
    bool erase(LhsString lhs);

    /**
     * @name exists
     * @brief return true/false whether LhsString exists.
     *
     * @param lhs string
     * @return bool (operational result)
     */
    inline bool exists(LhsString lhs);

    /**
     * @name getLhsBegin
     * @brief return an iterator to the LhsString pointing at the beginning of the map
     *
     * @return std::map<std::string, int>::iterator
     */
    inline auto getLhsBegin();

    /**
     * @name getLhsEnd
     * @brief return an iterator to the lhs string pointing at the end of the map
     *
     * @return std::map<std::string, int>::iterator
     */
    inline auto getLhsEnd();

    /**
     * @name getRequired
     * @brief return required field for a given lhs string.
     *
     * @throws out_of_range if key not found.
     *
     * @param lhs string
     * @return bool (required field value)
     */
    inline bool getRequired(LhsString lhs);

    /**
     * @name getRhsString
     * @brief Given a mapping of an left-hand (lhs) string to a right-hand (rhs) string
     *        this method returns the right-hand (rhs) string associated with its lhs.
     *
     * @throws out_of_range if key not found.
     *
     * @param lhs string
     * @return string (rhs)
     */
    inline RhsString getRhsString(LhsString lhs);

    /**
     * @name getType
     * @brief Fetch a ValueType for a given mapping.
     *
     * @throws out_of_range if key not found.
     *
     * @param lhs string
     * @return ValueTypes
     */
    inline ValueTypes getType(LhsString lhs);


    /**
     * @name rhsExists
     * @brief Does the rhs string exist?
     *
     * @param rhs string
     * @return bool (operational result)
     */
    bool rhsExists(string rhs);

    /**
     * @name validate
     * @brief execute the validator for a given lhs string against a given input
     *        and return the boolean result.
     *
     * @warning out_of_range exception will be thrown if lhs string does not exist.
     *
     * @param lhs string (the map record referencing the validator.
     * @param input string (the value to validate)
     * @return bool (operation result)
     */
    inline bool validate(LhsString lhs, string input);

    /**
     * disabled operators (start here)
     */
    void operator=(auto) = delete;

    void operator+=(auto) = delete;

    void operator-=(auto) = delete;

    void operator==(auto) = delete;

    void operator!=(auto) = delete;

    void operator<(auto) = delete;

    void operator<=(auto) = delete;

    void operator>(auto) = delete;

    void operator>=(auto) = delete;

    void operator>>(auto) = delete;

    void operator<<(auto) = delete;
    /**
     * disabled operators (ends here)
     */
private:
    /**
     * Map left string (LhsString) to right string (RhsString).
     */
    map <LhsString, RhsString> StringMap;
    /**
     * Map right string to ValueTypes.
     */
    map <RhsString, ValueTypes> ValueTypeMap;
    /**
     * Map left string (LhsString) to boolean indicating a record is required.
     */
    map<LhsString, bool> RequiredMap;
    /**
     * Map left string (LhsString) to a boolean func (BoolFuncPtr) validator function.
     */
    map <LhsString, ConfigStateMapValidator> ValidatorMap;
}; /* end of ConfigStateMap class */

#include "add.cpp"
#include "Constructor.cpp"
#include "count.cpp"
#include "erase.cpp"
#include "exists.cpp"
#include "getLhsBegin.cpp"
#include "getLhsEnd.cpp"
#include "getRequired.cpp"
#include "getRhsString.cpp"
#include "getType.cpp"
#include "rhsExists.cpp"
#include "validate.cpp"
