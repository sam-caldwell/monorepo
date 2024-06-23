/**
 * @name Logger/src/Logger/newContext.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Logger_Logger_newContext_cpp
#define Logger_Logger_newContext_cpp


/**
 * @name newContext [default]
 * @brief copy-create a context or use defaults
 *
 * @return bool
 */
bool Logger::newContext() {
    Logging::Severity sev;
    Logging::Writer::Id writerId;
    string *tags;
    string *destination;

    if (context.size() > 255)
        throw out_of_range("no more than 256 logging contexts are allowed");

    if (context.size() == 0) {
        sev = Logging::Severity::info;
        writerId = Logging::Writer::stdout;
        tags = NULL;
        destination = NULL;
    } else {
        string temp;
        sev = context.top().severity();
        writerId = context.top().writerId();
        temp = string(context.top().tags());
        tags = &temp;
        temp = string(context.top().destination());
        destination = &temp;
    }
    volatile ContextId id = ContextId(context.size() - 1);
    Context c(id, sev, writerId, tags, destination);
    context.push(c);
    return true;
}

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
bool Logger::newContext(Logging::Severity sev, Logging::Writer::Id writer, string *tags, string *destination) {

    if (context.size() > 255)
        throw out_of_range("no more than 256 logging contexts are allowed");

    volatile ContextId id = ContextId(context.size() - 1);
    Context c(id, sev, writer, tags, destination);
    context.push(c);
    return true;
}

#endif /* Logger_Logger_newContext_cpp */
