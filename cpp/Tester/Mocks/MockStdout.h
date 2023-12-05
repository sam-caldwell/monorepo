/**
 * @name Tester/MockStdout.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Tester_Mock_Stdout_H
#define Tester_Mock_Stdout_H

#include <iostream>
#include <sstream>

namespace Mock {
    using namespace std;

    /**
     * @namespace Mock
     * @class Mock::Stdout
     * @brief A mock stdout interceptor class.
     */
    class Stdout {
    private:
        stringstream redirect;
        streambuf *oldBuffer;
        string lastMessage;
    public:
        Stdout() {
            lastMessage = "";
        }

        ~Stdout() {
            oldBuffer = NULL;
        }

        inline bool enable() {
            oldBuffer = std::cout.rdbuf(redirect.rdbuf());
            cout.rdbuf(redirect.rdbuf());
            return true;
        }

        inline bool capture() {
            string buffer;
            lastMessage = "";
            while (getline(redirect, buffer)) {
                lastMessage += buffer;
            }
            if (oldBuffer) cout.rdbuf(oldBuffer);
            return true;
        }

        string last() {
            return lastMessage;
        }
    };
}
#endif /* Tester_Mock_Stdout_H */