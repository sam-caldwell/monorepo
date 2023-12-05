/**
 * @name networking/TcpCommon/src/PipeResponse/main.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef networking_TcpCommon_PipeResponse_H
#define networking_TcpCommon_PipeResponse_H

#include <string>
#include "projects/networking/TcpCommon/src/messages.h"

namespace networking {
    /**
     * @class PipeResponse
     * @brief This is the response token for the network operations.
     */
    class PipeResponse {
    public:
        /**
         * @name default constructor
         */
        PipeResponse() = default;

        /**
         * @name class constructor
         * @brief define the state of the class.
         *
         * @param successFlag - bool (indicate yes/no whether operation successful)
         * @param msg - string (by reference) (indicates context of the class state).
         */
        PipeResponse(bool successFlag, const std::string &msg);

        /**
         * @name class destructor
         * @brief destroy/cleanup state
         */
         ~PipeResponse();

        /**
         * @name message
         * @brief return the message string.
         *
         * @return string
         */
        inline std::string message() const;

        /**
         * @name isSuccessful
         * @brief returns whether or not an operation is successful.
         *
         * @return bool
         */
        inline bool isSuccessful() const;

        /**
         * @name failure
         * @brief report failure
         * @param msg string reference
         * @return static PipeResponse
         */
        static PipeResponse failure(const std::string &msg = "");

        /**
         * @name success
         * @brief report success
         * @param msg string reference
         * @return static PipeResponse
         */
        static PipeResponse success(const std::string &msg = "");

    private:
        /**
         * @name _successFlag
         * @brief boolean flag indicating whether or not operation is successful
         */
        bool _successFlag = false;
        /**
         * @name _msg
         * @brief the message string which adds context to _successFlag value.
         */
        std::string _msg = "";
    };
}

#include "constructor.cpp"
#include "destructor.cpp"
#include "failure.cpp"
#include "isSuccessful.cpp"
#include "message.cpp"
#include "success.cpp"


#endif /* networking_TcpCommon_PipeResponse_H */