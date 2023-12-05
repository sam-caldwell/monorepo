/**
 * @name networking/TcpCommon/tests/TestTcpClientBasics.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "projects/Tester/TestBase/main.h"

#include <iostream>
#include <thread>
#include "projects/networking/TcpClient/src/main.h"

#ifdef _WIN32
#include <Windows.h>
#else
#include <unistd.h>
#endif
#include <iostream>
#include <cstdlib>


void netcat() {
    cout << COLOR_GRAY << "starting netcat listener." << COLOR_RESET << endl;
    system("/usr/bin/nc -lv 127.0.0.1 1337");
}

class TestTcpClientBasics : public TestBase {
private:
    std::thread *listener_thread;
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestTcpClientBasics(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestTcpClientBasics() {}

    /**
     * @name test_constructor
     * @brief verify the client constructor can instantiate with a valid FileDescriptor.
     *
     * @return bool
     */
    bool test_constructor() {
        networking::TcpClient p;
        return expect(!p.isConnected(), "default isConnected() is false.");
    }

    /**
     * @name start_listener
     * @brief start a netcat listener on 127.99.99.1:1337
     *
     * @return bool
     */
    bool start_listen() {
        listener_thread = new std::thread(netcat);
        listener_thread->detach();
        return true;
    }

    bool test_listener(){
        sleep(2);
        cout << COLOR_GRAY << "testing netcat listener." << COLOR_RESET << endl;
        system("echo test | /usr/bin/nc 127.0.0.1 1337");
        cout << COLOR_GRAY << "testing message sent." << COLOR_RESET << endl;
        return true;
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_constructor(), "expect test_constructor() ok") &&
               expect(start_listen(), "expect we can start a listener thread. ok") &&
               expect(test_listener(), "expect we can test our listener. ok");
    }

};