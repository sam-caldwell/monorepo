/**
 * @name Application/destructor.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name Class destructor
 * @brief tear down the system.
 */
Application::~Application() {
    if (signal) delete signal;
    if (event) delete event;
    if (log) delete log;
    if (config) delete config;
}