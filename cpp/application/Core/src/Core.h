/**
 * @name Application/src/Application.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 * @brief This is the Application Core class used as a starting point
 *        for building applications within the monorepo.
 *
 * @details This class provides core services in the monorepo for developing
 *          applications.  Core functionality includes observability tooling,
 *          signal handling, configuration, etc.
 */
#ifndef Application_Core_H
#define Application_Core_H

//#include "projects/Event/src/Event.h"
#include "projects/Signals/src/Signals.h"
#include "projects/logging/Logger/src/Logger.h"
#include "projects/common/types/ConfigStateMap.h"
#include "projects/application/Configuration/src/Configuration.h"

namespace Application {
    class Core {
        /**
         * @class Application
         * @details This is the base class for c++ applications in this repository.
         *          The Application presupposes a ConfigStateMap to describe the
         *          internal configuration schema for the application, a table of
         *          signal handlers and the configuration fo any loggers and event
         *          services.
         */
    public:
        Application() = delete;

        /**
         * @name Default Class constructor.
         * @brief Initialize class state, configuring and executing any configuration options.
         *        The end-state should be a loaded Configuration (*config) object, an application
         *        logging system (*log), and event-emitter (*event) and a signal handler (*signal)
         *        which can be used by the Application.
         *
         * @param *map ConfigStateMap pointer
         * @param argc - int - count of arguments from commandline
         * @param *argv[] - reference to an array of argument strings.
         * @param *signals SignalHandler table (default: NULL)
         */
        Application(ConfigStateMap
        *map,
        int argc,
        char *argv[], SignalHandler::Table
        *
        signals = NULL
        );

        /**
         * @name Class destructor
         * @brief tear down the system.
         */
        ~Application();

        /**
         * @name start
         * @brief start the application payload (this should be overridden by the application class)
         *
         * @return int exit code
         */
        int start() { return 0; }

    protected:
        /**
         * @name Configuration Store
         * @brief The configuration pointer is the data repository for parsing, validating
         *        and storing the configuration used throughout the lifecycle of an application.
         */
        Configuration *config;
        /**
         * @name Logger Handler
         * @brief The *log pointer allows the application to write logs to a pre-defined target
         *        with syslog-compliant logging levels and other parameters. These logs may be
         *        sent to stdout, stderr or a remote target.
         */
        Logger *log;
        /**
         * @name Event Handler
         * @brief The *event pointer references an application-level event handler used to capture
         *        and emit highly structured event data which can be printed to stdout, stderr or
         *        most commonly sent to a remote event processing target for analysis.
         */
//    Event *event;
    };
}

#include "constructor.cpp"
#include "destructor.cpp"
#include "loggerSetup.cpp"
#include "start.cpp"
//#include "events_setup.cpp"

#endif /* Application_Core_H */
