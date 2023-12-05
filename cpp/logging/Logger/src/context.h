/**
 * @name Logger/src/context.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Logger_Context_h
#define Logger_Context_h

#include <ctime>
#include "projects/logging/Logger/src/writer.h"
#include "projects/logging/Logger/src/severity.h"

typedef unsigned short int ContextId;

class Context {
private:
    ContextId contextId;
    Logging::Severity currentSeverity;
    writerBase *currentWriter = NULL; //The writer for our current context.
    string currentTags;
    Logging::Writer::Id currentWriterId;
public:
    /**
     * @name Context
     * @brief Class constructor for configuring a context object.
     *
     * @param id - contextId
     * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
     * @param writerId - Log writer Id (e.g. stdout, stderr, file, syslog, https)
     * @param tags - a comma-delimited set of quoted strings used to map attributes
     *               of logs across our observability tooling.
     * @param destination - a String url in the form <protocol>:://<addr[:<port>]> used to
     *                      identify the log receiver (e.g. file://path/filename or syslog:://<ip:port>).
     */
    Context(ContextId id,
            Logging::Severity sev,
            Logging::Writer::Id writerId,
            string *tags,
            string *destination);

    /**
     * @name Class destructor
     * @brief Clean up any context resources.
     */
    ~Context();

    /**
     * @name tags
     * @brief return the current tag set (string) for the current context
     *
     * @return string
     */
    inline string tags() { return currentTags; }

    /**
     *  @name destination
     *  @brief return the string destination (url) for the current context.
     */
    inline string destination() { return currentWriter->getDestination(); }

    /**
     * @name severity [read-only]
     * @brief return the log severity (level) of the current context
     *
     * @return Logging::Severity
     */
    inline Logging::Severity severity() { return currentSeverity; }

    /**
     * @name writerId
     * @brief return the writerId of the current context
     *
     * @return Logging::Writer::Id
     */
    inline Logging::Writer::Id writerId() { return currentWriterId; }

    /**
     * @name write
     * @brief Provides an interface between logging endpoints (e.g. debug(), error(), etc.) and the
     *        underlying writer function for the current context.
     * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
     * @param msg
     * @return bool
     */
    inline bool write(Logging::Severity sev, const string *msg);

    /**
     * @name id
     * @brief return the current contextId
     *
     * @return unsigned char
     */
    inline unsigned char id() { return contextId; }
};

#include "projects/logging/Logger/src/Context/constructors.cpp"
#include "projects/logging/Logger/src/Context/destructors.cpp"
#include "projects/logging/Logger/src/Context/write.cpp"

#endif /* Logger_Context_h */