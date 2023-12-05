/**
 * @name Application/constructor.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <map>
#include <vector>
//#include "events_setup.cpp"

/**
 * @name Default Class constructor.
 * @brief Initialize class state, configuring and executing any configuration options.
 *
 * @throws out_of_range exception where a configuration is expected but not found or
 *         where the value fails validation.
 *
 * @param *map ConfigStateMap pointer
 * @param argc - int - count of arguments from commandline
 * @param *argv[] - reference to an array of argument strings.
 * @param *signals SignalHandler table (default: NULL)
 */
Application::Application(ConfigStateMap *map,
                         int argc,
                         char *argv[],
                         SignalHandler::Table *signals = NULL
) {
    config = new Configuration;
    /**
     * Configure the application using various sources (e.g. Environment variables,
     * command-line arguments, configuration files) and present the result as a single,
     * centralized Configuration object referenced by *config.
     */
    configure(config, map, argc, argv);
    /**
     * Using the application configuration (*config) object, configure the logger
     * which will present a logger object referenced by this->*log for use
     * by the application.
     */
    log = new Logger(config);
    log->debug("Logging is enabled now.");
    /**
     * Using the application configuration (*config) object, configure an events
     * processor which will make available this->*event for use by the application
     * to emit events to a given event receiver.
     */
//    event = eventSetup(config);
    /**
     * Configure signal handling by passing in a table of signal handler
     * functions associated with specific signals to be caught and handled.
     */
    if (signals) SignalHandler::init(signals);
}