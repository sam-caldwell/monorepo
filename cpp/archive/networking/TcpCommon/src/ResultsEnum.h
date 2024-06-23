/**
 * @name networking/TcpCommon/src/ResultsEnum.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef networking_TcpCommon_ResultsEnum_H
#define networking_TcpCommon_ResultsEnum_H

namespace networking {

    namespace FdWait {
        /**
         * @enum Result
         * @brief this is the waitFor() result.
         */
        enum Result {
            FAILURE,
            TIMEOUT,
            SUCCESS
        };
    }
}

#endif /* networking_TcpCommon_ResultsEnum_H */
