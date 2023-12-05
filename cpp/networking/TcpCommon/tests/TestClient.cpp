/**
 * @name networking/TcpCommon/tests/TestClient.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "projects/Tester/TestBase/main.h"

#include <iostream>
#include "projects/networking/TcpCommon/src/main.h"

namespace testFixture {
    static networking::ClientEvent ce;
    static std::string msg;


    void dummy(const networking::Client &a, networking::ClientEvent b, const std::string &c) {
        ce = b;
        msg = c;
    }
}

class TestClient : public TestBase {
private:
    networking::ClientEvent stateCe;
    std::string stateMsg;
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestClient(string n) {
        name = n;
        stateMsg = "";
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestClient() {}


    /**
     * @name test_constructor
     * @brief verify the client constructor can instantiate with a valid FileDescriptor.
     *
     * @return bool
     */
    bool test_constructor() {
        networking::Client p(1);
        return expect(!p.isConnected(), "default isConnected() is false.");
    }

    /**
     * @name test_get_ip_default
     * @brief verify getIp work default (empty) works.
     *
     * @return bool
     */
    bool test_get_ip_default() {
        networking::Client p(1);
        return expect(p.getIp() == "", "Expect default value is empty");
    }

    /**
     * @name test_set_get_ip
     * @brief verify setIp and getIp work
     *
     * @params s - string
     * @return bool
     */
    bool test_set_get_ip(string s) {
        networking::Client p(1);
        p.setIp(s);
        return p.getIp().compare(s) == 0;
    }

    /**
     * @name test_setEventHandler
     * @brief test we can set an event handler (noop) function
     *
     * @return bool
     */
    bool test_setEventHandler() {
        networking::Client p(1);
        p.setIp("127.0.0.1");
        p.setEventsHandler(testFixture::dummy);
        return true;
    }

    /**
     * @name test_publishEvent()
     * @brief verify the publishEvent method works.
     *
     * @return bool
     */
    bool test_publishEvent(networking::ClientEvent ce, std::string msg) {
        networking::Client p(1);
        p.setIp("127.0.0.1");
        p.setEventsHandler(testFixture::dummy);
        p.publishEvent(ce, msg);
        return expect(testFixture::ce == ce, "Expect clientEvent ok") &&
               expect(testFixture::msg.compare(msg) == 0, "Expect msg ok");
    }

    /**
     * @name test_operator_equality
     * @brief verify we the == operator works as expected
     *
     * @return bool
     */
    bool test_operator_equality() {
        networking::Client a(1);
        a.setIp("1");
        networking::Client b(1);
        b.setIp("1");
        networking::Client c(2);
        c.setIp("2");
        return (a == b) && (a != c) && (b != c);
    }

    /**
     * @name test_close
     * @brief verify we can safely run close() without having opened anything really.
     *
     * @return bool
     */
    bool test_close() {
        networking::Client a(1);
        a.close();
        return true;
    }

    /**
     * @name test_startListen
     * @brief test that we can spawn a listener thread without error.
     *
     * @return bool
     */
//    bool test_startListen() {
//        networking::Client a(2);
//            a.startListen();
//        a.close();
//        return true;
//    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        networking::ClientEvent ce1 = networking::ClientEvent::DISCONNECTED;
        networking::ClientEvent ce2 = networking::ClientEvent::INCOMING_MSG;
        return expect(test_constructor(), "expect test_constructor() ok") &&
               expect(test_get_ip_default(), "expect test_get_ip_default() ok") &&
               expect(test_set_get_ip("127.0.0.1"), "test_set_get_ip(127) ok") &&
               expect(test_set_get_ip("192.168.22.1"), "test_set_get_ip(192) ok") &&
               expect(test_set_get_ip("172.16.10.11"), "test_set_get_ip(172) ok") &&
               expect(test_operator_equality(), "expect test_operator_equality() ok") &&
               expect(test_setEventHandler(), "expect test_setEventHandler() ok") &&
               expect(test_publishEvent(ce1, "test1"), "expect test_publishEvent(1) ok") &&
               expect(test_publishEvent(ce2, "test2"), "expect test_publishEvent(2) ok") &&
               expect(test_close(), "expect test_close() ok");// &&
//               expect(test_startListen(), "expect test_startListen() ok");
    }

};