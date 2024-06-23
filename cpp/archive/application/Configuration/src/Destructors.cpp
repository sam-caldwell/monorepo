/**
 * @name Configuration/src/Destructors.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name Class destructor
 * @brief tear down the system.
 */
Configuration::~Configuration() {
    if (tree) delete tree;
}
