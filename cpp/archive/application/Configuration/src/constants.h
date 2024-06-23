/**
 * @name Configuration/src/constants.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 * @brief this file contains the constant string identifiers for
 *        standardized configuration objects used by the monorepo.
 */
#ifndef Configuration_Constants_H
#define Configuration_Constants_H

/**
 * @name ConfigFile
 * @brief defines the configuration file.
 */
const string ConfigFile = "config.file";


/**
 * --- projects/Logger ---
 */
/**
 * @name ConfigLogLevel
 * @brief defines the log severity level (see RFC5424)
 */
const string ConfigLogLevel = "log.level";
/**
 * @name ConfigLogOutput
 * @brief defines the string version of the log output target (writer).  See Logging::Writer::Id
 */
const string ConfigLogOutput = "log.output";
/**
 * @name ConfigLogTags
 * @brief defines the base tags used by the logger.
 */
const string ConfigLogTags = "log.tags";
/**
 * @name ConfigLogDestination
 * @brief defines the logging destination string (path/file or URL).
 */
const string ConfigLogDestination = "log.destination";

/**
 * --- projects/LogServer ---
 */
/**
 * @name ConfigLogServerIp
 * @brief a parameter used to define the Ip address/fqdn used to recv logs.
 */
const string ConfigLogServerIp = "logServer.net.ip";
/**
 * @name ConfigLogServerPort
 * @brief a parameter used to define the network port used to recv logs.
 */
const string ConfigLogServerPort = "logServer.net.port";
/**
 * @name ConfigLogServerTarget
 * @brief a parameter used to define the target directory where rec'd logs will be written.
 */
const string ConfigLogServerTarget = "logServer.target";


#endif /*  Configuration_Constants_H */

