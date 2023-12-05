/**
 * @name networking/TcpCommon/src/PipeResponse/constructor.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
namespace networking {
    /**
     * @name class constructor
     * @brief define the state of the class.
     *
     * @param successFlag - bool (indicate yes/no whether operation successful)
     * @param msg - string (by reference) (indicates context of the class state).
     */
    PipeResponse::PipeResponse(bool successFlag, const std::string &msg) :
            _successFlag{successFlag},
            _msg{msg} {}
}