/**
 * @name networking/TcpCommon/src/PipeResponse/main.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

namespace networking {
    /**
     * @name message
     * @brief return the message string.
     *
     * @return string
     */
    inline std::string PipeResponse::message() const {
        return _msg;
    }
}