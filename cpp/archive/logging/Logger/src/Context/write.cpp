/**
 * @name Logger/src/Context/write.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Logger_Context_write_cpp
#define Logger_Context_write_cpp
using namespace std;

/**
 * @name write
 * @brief Provides an interface between logging endpoints (e.g. debug(), error(), etc.) and the
 *        underlying writer function for the current context.
 * @param sev - Log severity (e.g. emergency...debug) consistent with RFC5424
 * @param msg
 * @return bool
 */
inline bool Context::write(Logging::Severity sev, const string *msg) {
    if (currentSeverity >= sev) {
        const auto now = std::chrono::system_clock::now();
        stringstream buffer;
        buffer << "{"
               << "\"n\":" << chrono::duration_cast<chrono::seconds>(now.time_since_epoch()).count() << ","
               << "\"c\":" << contextId << ","
               << "\"s\":\"" << Logging::ReverseSeverity.at(sev) << "\","
               << "\"m\":\"" << *currentWriter->encode(msg) << "\","
               << "\"t\":[" << currentTags << "]"
               << "}";
        currentWriter->write(&buffer);
        return true;
    }
    return false;
}

#endif /* Logger_Context_write_cpp */
