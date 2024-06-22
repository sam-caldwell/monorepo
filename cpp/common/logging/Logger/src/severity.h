/**
 * @name Logger/src/severity.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Logging_Severity_h
#define Logging_Severity_h

#include <map>
#include <string>

namespace Logging {
/**
 * @enum Severity
 * @brief The enumerated values of RFC5424 log severity.
 */
    enum Severity : unsigned char {
        emergency = 0,  // emergency; system is unusable
        alert = 1,      // action must be taken immediately
        critical = 2,   // critical condition
        error = 3,      // error condition
        warning = 4,    // warning condition
        notice = 5,     // normal but significant condition
        info = 6,       // informational message
        debug = 7,      // debug-level messages
    };


/**
 * @name SeverityStrings
 * @brief A constant lookup table (map) for converting string severity levels to a numeric value.
 *
 */
    const static map <string, Logging::Severity> SeverityStrings = {
            {"emergency", Logging::Severity::emergency},
            {"alert",     Logging::Severity::alert},
            {"critical",  Logging::Severity::critical},
            {"error",     Logging::Severity::error},
            {"warning",   Logging::Severity::warning},
            {"notice",    Logging::Severity::notice},
            {"info",      Logging::Severity::info},
            {"debug",     Logging::Severity::debug}
    };

/**
 * @name ReverseSeverity
 * @brief A constant lookup table (map) for converting numeric severity levels to strings.
 */
    const static map <Logging::Severity, string> ReverseSeverity = {
            {Logging::Severity::emergency, "emergency"},
            {Logging::Severity::alert,     "alert"},
            {Logging::Severity::critical,  "critical"},
            {Logging::Severity::error,     "error"},
            {Logging::Severity::warning,   "warning"},
            {Logging::Severity::notice,    "notice"},
            {Logging::Severity::info,      "info"},
            {Logging::Severity::debug,     "debug"}
    };
}
#endif /* Logging_Severity_h */