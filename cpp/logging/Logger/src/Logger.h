/**
 * @name Logger/src/Logger.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Logger_h
#define Logger_h

#include <stack>
#include <string>
#include "projects/application/Configuration/src/Configuration.h"
#include "projects/logging/Logger/src/writer.h"
#include "projects/logging/Logger/src/severity.h"
#include "projects/logging/Logger/src/context.h"
#include "projects/common/Validators/validators.h"

class Logger {
private:
    /**
     * @name context
     * @brief This is the private stack of all Context objects for our logger.
     */
    std::stack <Context> context;

public:
    /**
     * @name Logger (default)
     * @brief simplest stdout logger possible.
     */
    Logger();

    /**
     * @name Logger (by Configuration)
     * @brief a simple method of configuring a Logger object using the Configuration
     *        project.
     */
    Logger(Configuration *cfg);

    /**
     * @name Logger
     * @brief configure the logger with no destination url or tags (stdout, stderr).
     *
     * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
     * @param writer - Log writer Id (e.g. stdout, stderr, file, syslog, https)
     */
    Logger(Logging::Severity sev, Logging::Writer::Id writer);

    /**
     * @name Logger
     * @brief configure the logger with no destination url (for stdout, stderr).
     *
     * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
     * @param writer - Log writer Id (e.g. stdout, stderr, file, syslog, https)
     * @param tags - a comma-delimited set of quoted strings used to map attributes
     *               of logs across our observability tooling.
     */
    Logger(Logging::Severity sev, Logging::Writer::Id writer, string tags);

    /**
     * @name Logger
     * @brief configure the logger with a destination url (for file, syslog, https).
     *
     * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
     * @param writer - Log writer Id (e.g. stdout, stderr, file, syslog, https)
     * @param tags - a comma-delimited set of quoted strings used to map attributes
     *               of logs across our observability tooling.
     * @param destination - a String url in the form <protocol>:://<addr[:<port>]> used to
     *                      identify the log receiver (e.g. file://path/filename or syslog:://<ip:port>).
     */
    Logger(Logging::Severity sev, Logging::Writer::Id writer, string tags, string destination);

    /**
     * @name ~Logger
     * @name Class destructor
     * @brief destroy all logging contexts and cleanup..
     */
    ~Logger();

    /**
     * @name newContext [default]
     * @brief copy-create a context or use defaults
     *
     * @return bool
     */
    bool newContext();

    /**
     * @name newContext
     * @brief Create a new logging context.
     *
     * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
     * @param writer - Log writer Id (e.g. stdout, stderr, file, syslog, https)
     * @param tags - a comma-delimited set of quoted strings used to map attributes
     *               of logs across our observability tooling.
     * @param destination - a String url in the form <protocol>:://<addr[:<port>]> used to
     *                      identify the log receiver (e.g. file://path/filename or syslog:://<ip:port>).
     * @return bool
     */
    bool newContext(Logging::Severity sev, Logging::Writer::Id writer, string *tags, string *destination);

    /**
     * @name popContext
     * @brief remove the current context (which will destroy this context).
     *
     * @return bool
     */
    bool popContext();

    /**
     * @name contextCount
     * @brief return the number of contexts on the stack
     *
     * @return unsigned
     */
    inline unsigned contextCount() { return unsigned(context.size()); }

    /**
     * @name severity  [read-only]
     * @brief return the log severity (level) of the current context
     *
     * @return Logging::Severity
     */
    inline Logging::Severity severity() { return context.top().severity(); }

    /**
     * @name writerId
     * @brief return the writerId of the current context
     *
     * @return Logging::Writer::Id
     */
    inline Logging::Writer::Id writerId() { return context.top().writerId(); }

    /**
     * @name tags
     * @brief return the current tag set (string) for the current context
     *
     * @return string
     */
    inline string tags() { return context.top().tags(); }

    /**
     *  @name destination
     *  @brief return the string destination (url) for the current context.
     */
    inline string destination() { return context.top().destination(); }
    /*                                         *
     * ------------ ALERT WRITERS ------------ *
     *                                         */
    /**
     * @name alert
     * @brief print alert level log message
     * @param *msg string pointer
     * @return bool
     */
    inline bool alert(const string *msg) { return context.top().write(Logging::Severity::alert, msg); }

    /**
     * @name alert
     * @brief print alert level log message
     * @param msg string
     * @return bool
     */
    inline bool alert(const string msg) { return alert(&msg); }

    /**
     * @name alert
     * @brief print alert level log message
     * @param *msg string pointer
     * @return bool
     */
    inline bool critical(const string *msg) { return context.top().write(Logging::Severity::critical, msg); }

    /**
     * @name alert
     * @brief print alert level log message
     * @param msg string
     * @return bool
     */
    inline bool critical(const string msg) { return critical(&msg); }

    /**
     * @name alert
     * @brief print alert level log message
     * @param *msg string pointer
     * @return bool
     */
    inline bool debug(const string *msg) { return context.top().write(Logging::Severity::debug, msg); }

    /**
     * @name alert
     * @brief print alert level log message
     * @param msg string
     * @return bool
     */
    inline bool debug(const string msg) { return debug(&msg); }

    /**
     * @name alert
     * @brief print alert level log message
     * @param *msg string pointer
     * @return bool
     */
    inline bool emergency(const string *msg) { return context.top().write(Logging::Severity::emergency, msg); }

    /**
     * @name alert
     * @brief print alert level log message
     * @param msg string
     * @return bool
     */
    inline bool emergency(const string msg) { return emergency(&msg); }

    /**
     * @name alert
     * @brief print alert level log message
     * @param *msg string pointer
     * @return bool
     */
    inline bool error(const string *msg) { return context.top().write(Logging::Severity::error, msg); }

    /**
     * @name alert
     * @brief print alert level log message
     * @param msg string
     * @return bool
     */
    inline bool error(const string msg) { return error(&msg); }

    /**
     * @name alert
     * @brief print alert level log message
     * @param *msg string pointer
     * @return bool
     */
    inline bool notice(const string *msg) { return context.top().write(Logging::Severity::notice, msg); }

    /**
     * @name alert
     * @brief print alert level log message
     * @param msg string
     * @return bool
     */
    inline bool notice(const string msg) { return notice(&msg); }

    /**
     * @name alert
     * @brief print alert level log message
     * @param *msg string pointer
     * @return bool
     */
    inline bool warning(const string *msg) { return context.top().write(Logging::Severity::warning, msg); }

    /**
     * @name alert
     * @brief print alert level log message
     * @param msg string
     * @return bool
     */
    inline bool warning(const string msg) { return warning(&msg); }

};

#include "projects/logging/Logger/src/Logger/constructors.cpp"
#include "projects/logging/Logger/src/Logger/destructors.cpp"
#include "projects/logging/Logger/src/Logger/newContext.cpp"
#include "projects/logging/Logger/src/Logger/popContext.cpp"


#endif /* Logger_h */
