/**
 * @name networking/TcpCommon/src/PipeResponse/main.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

namespace networking {
    /**
     * @name isSuccessful
     * @brief returns whether or not an operation is successful.
     *
     * @return bool
     */
    inline bool PipeResponse::isSuccessful() const {
        return _successFlag;
    }
}