/**
 * @name Configuration/src/getType.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/common/types/ValueTypes.h"

/**
 * @name getType
 * @brief return the ValueTypes type of the given node.
 *
 * @throws out_of_range exception on invalid key
 *
 * @param key string
 * @return ValueTypes
 */
inline ValueTypes Configuration::getType(string key) {
    try {
        return tree->getType(key);
    } catch (out_of_range &e) {
        throw out_of_range("Configuration::getType() error:" + string(e.what()));
    }
}
