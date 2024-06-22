/**
 * @name networking/TcpCommon/tests/TestBasics.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "../src/PipeResponse/main.h"

#include <iostream>
#include "../src/PipeResponse/main.h"

class TestBasics : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestBasics(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestBasics() {}

    /**
     * @name test_ClientEvent
     * @brief verify the ClientEvent enum
     *
     * @return bool
     */
    bool test_ClientEvent(networking::ClientEvent evt, int expected) {
        return int(evt) == expected;
    }

    /**
     * @name test_FdWait_Result_enum
     * @brief verify that networking::FdWait::Result enum
     *
     * @return bool
     */
    bool test_FdWait_Result_enum(networking::FdWait::Result evt, int expected) {
        return int(evt) == expected;
    }

    /**
     * @name test_MAX_PACKET_SIZE
     * @brief validate that MAX_PACKET_SIZE is set as expected
     *
     * @return bool
     */
    bool test_MAX_PACKET_SIZE(int t) {
        return MAX_PACKET_SIZE == t;
    }

    /**
     * @name test_MAX_CLIENT_TIMEOUT
     * @brief validate that MAX_CLIENT_TIMEOUT is set as expected
     */
    bool test_MAX_CLIENT_TIMEOUT(int t){
        return MAX_CLIENT_TIMEOUT == t;
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_ClientEvent(networking::ClientEvent::DISCONNECTED, 0),
                      "test_ClientEvent(DISCONNECTED,0) ok") &&
               expect(test_ClientEvent(networking::ClientEvent::INCOMING_MSG, 1),
                      "test_ClientEvent(INCOMING_MSG,1) ok") &&
               expect(test_FdWait_Result_enum(networking::FdWait::Result::FAILURE, 0),
                      "test_FdWait_Result_enum(FAILURE,0) ok") &&
               expect(test_FdWait_Result_enum(networking::FdWait::Result::TIMEOUT, 1),
                      "test_FdWait_Result_enum(TIMEOUT,1) ok") &&
               expect(test_FdWait_Result_enum(networking::FdWait::Result::SUCCESS, 2),
                      "test_FdWait_Result_enum(SUCCESS,2) ok") &&
               expect(test_MAX_PACKET_SIZE(4096), "expect MAX_PACKET_SIZE is 4096") &&
               expect(test_MAX_CLIENT_TIMEOUT(3600), "expect MAX_CLIENT_TIMEOUT is 3600");
    }

};