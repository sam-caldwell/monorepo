/**
 * @name TesterFramework/run.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

int Tester::run() {
    while (!tests.empty()) {
        TestBase *t = tests.front();
        tests.pop();
        cout << COLOR_BLUE << "Starting: " << t->get_name() << COLOR_RESET << endl;
        bool exit_code;
        try {
            exit_code = t->main();
            cout << COLOR_BLUE << "\tresult_code:" << exitCodeToString(exit_code) << COLOR_RESET << endl;
        } catch (exception &e) {
            exit_code=false;
            cout << COLOR_RED << "\tresult_code:" << exitCodeToString(exit_code) << endl
                 << COLOR_RED << "\tuncaught exception: " << e.what() << COLOR_RESET << endl;
        }
        if (exit_code) {
            cout << "\t"
                 << COLOR_RESET
                 << COLOR_GREEN
                 << "[+]PASSED"
                 << COLOR_RESET
                 << ": "
                 << t->get_name()
                 << endl;
            passed++;
        } else {
            cout << "\t"
                 << COLOR_RESET
                 << COLOR_RED
                 << "[-]FAILED"
                 << ": "
                 << t->get_name()
                 << COLOR_RESET
                 << endl;
            failed++;
        }
        cout.flush();
        cout << endl;
    } /**end of while*/
    cout << endl
         << "Results:" << endl
         << "  test_count  : " << (test_count) << endl
         << "  (+) passed  : " << passed << endl
         << "  (-) failed  : " << failed << endl
         << "  ( ) skipped : " << skipped << endl;

    log << setw(10) << count() << ","
        << setw(10) << number_passed() << ","
        << setw(10) << number_failed() << ","
        << setw(10) << number_skipped() << ","
        << test_group_name
        << endl;

    if (failed == 0) {
        cout << COLOR_GREEN << "PASS: " << test_group_name << " tests passing" << COLOR_RESET << endl;
        return 0;
    } else {
        cout << COLOR_RED << "  FAIL: " << test_group_name << " tests failed" << COLOR_RESET << endl;
        return failed;
    }
}
