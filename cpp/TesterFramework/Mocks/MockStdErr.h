/**
 * @name TesterFramework/MockStderr.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Tester_Mock_Stderr_H
#define Tester_Mock_Stderr_H

#include <iostream>
#include <sstream>

namespace Mock {
    using namespace std;

    /**
     * @namespace Mock
     * @class Mock::StdErr
     * @brief A mock stderr interceptor class.
     */
    class StdErr {
    private:
        stringstream redirect;
        streambuf *oldBuffer;
        string lastMessage;
    public:
        StdErr() {
            lastMessage = "";
        }

        ~StdErr() {
            oldBuffer = NULL;
        }

        inline bool enable() {
            oldBuffer = std::cerr.rdbuf(redirect.rdbuf());
            cerr.rdbuf(redirect.rdbuf());
            return true;
        }

        inline bool capture() {
            string buffer;
            lastMessage = "";
            while (getline(redirect, buffer)) {
                lastMessage += buffer;
            }
            if (oldBuffer) cerr.rdbuf(oldBuffer);
            return true;
        }

        string last() {
            return lastMessage;
        }
    };
}
#endif /* Tester_Mock_Stderr_H */