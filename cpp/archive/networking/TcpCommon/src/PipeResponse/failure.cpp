/**
 * @name networking/TcpCommon/src/PipeResponse/main.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <string>

namespace networking {
    /**
     * @name failure
     * @brief report failure
     * @param msg string reference
     * @return static PipeResponse
     */
    PipeResponse PipeResponse::failure(const std::string &msg) {
        return PipeResponse(false, msg);
    }
}