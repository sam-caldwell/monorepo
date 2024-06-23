/**
 * @name Logger/src/Logger/destructors.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Logger_Logger_constructors_cpp
#define Logger_Logger_constructors_cpp

/**
 * @name Logger (default)
 * @brief simplest stdout logger possible.
 */
Logger::Logger() {
    newContext();
}

/**
 * @name Logger (by Configuration)
 * @brief a simple method of configuring a Logger object using the Configuration
 *        project.
 *
 * @throws out_of_range exception if the configuration input is bad.
 */
Logger::Logger(Configuration *cfg) {
    if (cfg) {
        Logging::Severity sev;
        Logging::Writer::Id writerId;
        string tags;
        string destination;

        try {

            sev = Logging::SeverityStrings.at(cfg->getString(ConfigLogLevel));

            writerId = Logging::Writer::StringsTable.at(cfg->getString(ConfigLogOutput));

            tags = cfg->getString(ConfigLogTags);
            validateTags(&tags);

            switch (writerId) {
                case Logging::Writer::Id::file:
                case Logging::Writer::Id::syslog:
                case Logging::Writer::Id::https:
                    destination = cfg->getString(ConfigLogDestination);
                    validateDestination(&destination);
                    break;
                default:
                    destination = "";
                    break; //destination not needed.
            }
        } catch (out_of_range &e) {
            throw out_of_range("Logger configuration failed(1): " + string(e.what()));
        }
        newContext(sev, writerId, &tags, &destination);

    } else {
        throw runtime_error("Null configuration object passed to Logger");
    }
}

/**
 * @name Logger
 * @brief configure the logger with no destination url or tags (stdout, stderr).
 *
 * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
 * @param writer - Log writer Id (e.g. stdout, stderr, file, syslog, https)
 */
Logger::Logger(Logging::Severity sev, Logging::Writer::Id writer) {
    newContext(sev, writer, NULL, NULL);
}


/**
 * @name Logger
 * @brief configure the logger with no destination url (for stdout, stderr).
 *
 * @throws out_of_range exception if the configuration input is bad.
 *
 * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
 * @param writer - Log writer Id (e.g. stdout, stderr, file, syslog, https)
 * @param tags - a comma-delimited set of quoted strings used to map attributes
 *               of logs across our observability tooling.
 */
Logger::Logger(Logging::Severity sev, Logging::Writer::Id writer, string tags) {
    try {
        validateTags(&tags);
    } catch (out_of_range &e) {
        throw out_of_range("Logger Configuration failed(2): " + string(e.what()));
    }
    newContext(sev, writer, &tags, NULL);

}


/**
 * @name Logger
 * @brief configure the logger with a destination url (for file, syslog, https).
 *
 * @throws out_of_range exception if the configuration input is bad.
 *
 * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
 * @param writer - Log writer Id (e.g. stdout, stderr, file, syslog, https)
 * @param tags - a comma-delimited set of quoted strings used to map attributes
 *               of logs across our observability tooling.
 * @param destination - a String url in the form <protocol>:://<addr[:<port>]> used to
 *                      identify the log receiver (e.g. file://path/filename or syslog:://<ip:port>).
 */
Logger::Logger(Logging::Severity sev, Logging::Writer::Id writer, string tags, string destination) {
    try {
        validateTags(&tags);
        validateDestination(&destination);
    } catch (out_of_range &e) {
        throw out_of_range("Logger Configuration failed(3): " + string(e.what()));
    }
    newContext(sev, writer, &tags, &destination);
}


#endif /* Logger_Logger_constructors_cpp */
