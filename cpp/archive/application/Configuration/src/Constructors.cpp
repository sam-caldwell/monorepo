/**
 * @name Configuration/src/Configuration.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name Default Class constructor.
 * @brief Initialize state for blank state with static arbitrary value binary tree.
 * @details Configuration is a base class which can be inherited by both
 *          the EnvironmentVariableParser, CommandLineParser and
 *          ConfigurationFileParser as well as the consuming application.
 *          This class has a static tree property, shared across the various
 *          consumers.
 */
Configuration::Configuration() {
    tree = new ArbitraryKvBtree(true);
};