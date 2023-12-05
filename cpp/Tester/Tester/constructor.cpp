/**
 * @name Tester/constructor.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name class constructor
 * @brief setup the tester framework to begin executing tests.
 *
 * @param argc - number of arguments
 * @param argv - argument list
 */
Tester::Tester(int argc, char **argv) {
    if (argc < 3) throw runtime_error("test expects log directory, test name: (" + string(argv[0]) + ")");

    string logDir = argv[1];
    string testName = std::filesystem::path(argv[2]).filename().stem();//Ensure first char is capitalized.


    string logFile = logDir + "/build_log/" + string("test.") + string(testName) + string(".log");

    log.open(logFile, std::fstream::out | std::fstream::trunc);
    if (!log) {
        cout << COLOR_RED
             << "  ERROR: Failed to open log file (" << logFile << ")"
             << COLOR_RESET
             << endl;
        exit(99);
    }

    cout << COLOR_BLUE
         << "INIT: '" << testName << "' tests registering" << endl
         << "\tLogfile '" << logFile << "'" << endl
         << COLOR_RESET
         << endl;

    test_group_name = testName;
    test_count = 0;
    skipped = 0;
    passed = 0;
    failed = 0;

    debug_flag = false;


}
