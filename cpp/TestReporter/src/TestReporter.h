/**
 * @name TestReporter/src/TestReporter.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef TestReporter_H
#define TestReporter_H

#include <filesystem>
#include <fstream>
#include <iostream>
#include <queue>

using namespace std;

class TestReporter {
public:
    /**
     * @name Class Constructor
     * @brief initialize the class
     *
     * @param argc int
     * @param **argv char pointer to pointer
     */
    TestReporter(string logDirectory, string summaryFile);

    /**
     * @name Class destructor
     * @brief destroy class state.
     */
    ~TestReporter();

    int result(){
        return exit_code;
    }

private:
    ifstream source;
    ofstream log;
    int exit_code;
    bool has_failures;

};

#include "projects/TestReporter/src/constructors.cpp"
#include "projects/TestReporter/src/destructors.cpp"

#endif /* TestReporter_H */