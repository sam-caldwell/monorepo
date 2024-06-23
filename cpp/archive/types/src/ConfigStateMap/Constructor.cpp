/**
 * @name types/src/Constructor.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name class constructor [map]
 * @brief initialize the map with an existing map.
 * @param *_map ConfigStateMap pointer
 */
ConfigStateMap::ConfigStateMap(ConfigStateMap *_map) {
    StringMap = _map->StringMap;
    ValueTypeMap = _map->ValueTypeMap;
    ValidatorMap = _map->ValidatorMap;
}