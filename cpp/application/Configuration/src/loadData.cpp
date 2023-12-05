/**
 * @name Configuration/src/Configuration.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/common/types/ConfigStateMap.h"




/**
 * @name loadData
 * @brief Given the ConfigStateMap (by reference), the source name (lhsName) in this
 *        ConfigStateMap, used by the caller to get the data object (value), the target
 *        name (rhsName) to which we will store the loaded (type-aware) value, lookup and
 *        validate the string value then translate it into the appropriate ArbitraryKvBtree
 *        object in the Configuration binary tree structure.
 *
 * @throws runtime_error exception on any type, validation or storage issues.
 *
 * @param *map ConfigStateMap pointer
 * @param lhs string
 * @param value string
 *
 * @return bool (true: success, false: operation failed)
 */
inline bool Configuration::loadData(ConfigStateMap *map, string lhs, string value) {
    /**
     * Validate the lhsName (environment variable)
     * and its value using the validator function
     * configured in the map.
     */
    bool result = false;
    if (map->validate(lhs, value)) {
        /**
         * validate() has passed.  Now we can use the
         * datatype associated with the map to determine
         * how we will store the environment variable value.
         */
        const RhsString rhs = map->getRhsString(lhs);

        switch (map->getType(lhs)) {
            case Bool: {
                result = this->insert(rhs, parseBool(value, false));
                break;
            }
            case Double: {
                result = this->insert(rhs, parseDouble(value, (double) (0.0)));
                break;
            }
            case Float: {
                result = this->insert(rhs, parseFloat(value, (float) (0.0)));
                break;
            }
            case Int: {
                result = this->insert(rhs, parseInt(value, (int) (0)));
                break;
            }
            case String: {
                result = this->insert(rhs, (string)(value));
                break;
            }
            case Uint: {
                result = this->insert(rhs, parseUint(value, (uint)(0)));
                break;
            }
            default:
                throw runtime_error("unsupported Null type for environment variable");
                break;
        }; /* end of switch */
        /**
         * done
         */
    } else {
        /**
         * Validation failed.  Throw an exception using the error
         * from the map.
         */
        throw runtime_error("Configuration::loadData() error");
    } /* end of if...else... */
    return result;
}