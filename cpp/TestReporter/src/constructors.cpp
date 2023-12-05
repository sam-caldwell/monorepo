/**
 * @name TestReporter/src/constructors.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <fstream>
#include <glob.h>
#include <iostream>
#include "projects/AnsiColors/AnsiColors.h"
#include <stdexcept>
#include <sstream>
#include <string>
#include <string.h>
#include <vector>

using namespace std;

/**
 *
 * @param argc
 * @param argv
 */
TestReporter::TestReporter(string logDirectory, string summaryFile) {
    int totalCount = 0;
    int totalPassed = 0;
    int totalFailed = 0;
    int totalSkipped = 0;
    has_failures = false;
    exit_code = 0;

    log.open(summaryFile, std::ios_base::out);
    if (!log) throw ("failed to open log file: " + summaryFile);
    /**
     * glob the log files
     */
    string pattern = logDirectory + string("/*.log");
    cout << COLOR_BLUE << "globbing pattern: " << pattern << COLOR_RESET << endl;
    glob_t glob_result;
    memset(&glob_result, 0, sizeof(glob_result));
    if (glob(pattern.c_str(), GLOB_TILDE, NULL, &glob_result) != 0) {
        globfree(&glob_result);
        throw std::runtime_error("glob() failed");
    }
    /**
     * iterate over the fileList queue.  For each file, processFile()
     */

    cout << COLOR_BLUE << "Globbed " << glob_result.gl_pathc << " files" << COLOR_RESET << endl;
    for (size_t i = 0; i < glob_result.gl_pathc; ++i) {
        int count = 0;
        int passed = 0;
        int failed = 0;
        int skipped = 0;
        string projectName;
        string fileName = string(glob_result.gl_pathv[i]);

        cout << COLOR_BLUE << "current file: " << fileName << COLOR_RESET << endl;

        source.open(fileName, std::ios_base::in);
        if (!source) throw runtime_error("failed to open source file");

        string line;
        while (std::getline(source, line)) {
            std::istringstream data(line);
            char delim = ',';
            if (!(data >> count >> delim
                       >> passed >> delim
                       >> failed >> delim
                       >> skipped >> delim
                       >> projectName))
                throw runtime_error("failed to parse line");

            totalCount += count;
            totalPassed += passed;
            totalFailed += failed;
            totalSkipped += skipped;

            log << count << ","
                << passed << ","
                << failed << ","
                << skipped << ","
                << projectName
                << endl;
            log.flush();
        }
        source.close();
    }/* end of for() */
    globfree(&glob_result);

    has_failures = totalFailed > 0;

    cout << endl
         << "-----------------------------" << endl
         << "count  : " << totalCount << endl
         << "passed : " << totalPassed << endl
         << "failed : " << totalFailed << endl
         << "skipped: " << totalSkipped << endl
         << "-----------------------------" << endl
         << endl;

    log << endl
        << "-----------------------------" << endl
        << "count  : " << totalCount << endl
        << "passed : " << totalPassed << endl
        << "failed : " << totalFailed << endl
        << "skipped: " << totalSkipped << endl
        << "-----------------------------" << endl
        << endl;
}
