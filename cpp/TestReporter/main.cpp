/**
 * @name TestReporter/main.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "projects/TestReporter/src/TestReporter.h"

using namespace std;

#include "projects/AnsiColors/AnsiColors.h"

int main(int argc, char *argv[]) {
    int exit_code = 0;
    try {
        cout << COLOR_BLUE << "TestReporter starting" << COLOR_RESET << endl;
        /**
         * Required input is our logDirectory
        */
        if (argc < 2) throw runtime_error("Error: Expected log directory path");
        string logDirectory = string(argv[1]);
        cout << COLOR_BLUE << "logDirectory: " << logDirectory << endl;
        /**
         * make sure our log directory exists.
         */
        if (!std::filesystem::exists(logDirectory)) throw runtime_error("log directory not found.");
        /**
         * our output file (log) should be set to our ${logDirectory}/summary.log
         */
        string summaryFile = logDirectory + string("/../build_summary.log");
        cout << COLOR_BLUE << "summaryFile: " << summaryFile << endl;

        TestReporter t(logDirectory, summaryFile);
        exit_code = t.result();
        if (exit_code == 0)
            cout << COLOR_GREEN << "reporting finished [no errors]" << COLOR_RESET << endl;
        else {
            cout << COLOR_RED << "reporting finished [with errors]" << COLOR_RESET << endl;
        }
    } catch (exception &e) {
        cout << COLOR_RED << "reporting finished [with exception]" << COLOR_RESET << endl;
        cout << COLOR_RED << e.what() << COLOR_RESET << endl;
        return 1;
    }
    if (has_failures)
        cout << COLOR_RED << "TESTS Failed with errors" << COLOR_RESET << endl;
    else
        cout << COLOR_GREEN << "TESTS Passed with no errors" << COLOR_GREEN << endl;

    return exit_code;
}