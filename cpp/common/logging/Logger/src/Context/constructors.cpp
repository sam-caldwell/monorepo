/**
 * @name Logger/src/Context/constructors.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Logger_Context_constructors_cpp
#define Logger_Context_constructors_cpp

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
Context::Context(ContextId id,
                 Logging::Severity sev,
                 Logging::Writer::Id writerId,
                 string *tags,
                 string *destination) {

    contextId = id;

    currentSeverity = sev;

    currentWriterId = writerId;

    currentTags = (tags) ? *tags : ""; //ToDo: validate the tags

    switch (writerId) {
        case Logging::Writer::null:
            currentWriter = new writerNull;
            break;
        case Logging::Writer::stdout:
            currentWriter = new writerStdout;
            break;
        case Logging::Writer::stderr:
            currentWriter = new writerStderr;
            break;
        case Logging::Writer::file:
            currentWriter = new writerFile(destination);
            break;
        case Logging::Writer::syslog:
            currentWriter = new writerSyslog(destination);
            break;
        case Logging::Writer::https:
            currentWriter = new writerHttps(destination);
            break;
        default:
            throw out_of_range("Context encountered unsupported writer");
    }
}

#endif /* Logger_Context_constructors_cpp */