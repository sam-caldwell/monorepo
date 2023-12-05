/**
* @name Tester/MockFile.h
* @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
* @author Sam Caldwell <mail@samcaldwell.net>
*/
#ifndef Tester_Mock_File_H
#define Tester_Mock_File_H

#include <iostream>
#include <sstream>

namespace Mock {
    using namespace std;

    /**
     * @namespace Mock
     * @class Mock::File
     * @brief A mock stderr interceptor class.
     */
    class File {
    private:
        stringstream redirect;
        streambuf *oldBuffer;
        string lastMessage;
        string testFileName;
    public:
        File(string fileName) {
            testFileName = fileName;
            lastMessage = "";
        }

        ~File() {
            oldBuffer = NULL;
        }

        inline bool enable() {
            return true;
        }

        inline bool captureLastLine() {
            string line = "";
            std::ifstream log;
            log.open(testFileName, std::ifstream::in);
            while (getline(log, line)) {
                lastMessage = line;
            } //Read to the last line.
            log.close();
            return true;
        }

        inline string captureLine(unsigned lineNo) {
            //ToDo: Implement this feature.
            return "";
        }

        string last() {
            return lastMessage;
        }
    };
}
#endif /* Tester_Mock_File_H */