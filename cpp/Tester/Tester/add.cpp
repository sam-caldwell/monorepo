/**
 * @name Tester/add.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

bool Tester::add(TestBase *t, bool skip = false) {
    if (debug_flag) {
        t->set_debug();
    }
    test_count++;
    if (test_count > 254) {
        cout << "ERROR: Test design fails.  You should never have more than 200 tests for a single unit.  "
             << "Consider breaking your unit into smaller pieces so you can fully test each unit with fewer tests."
             << endl;
        exit(1);
    }
    if (skip) {
        skipped++;
        cout << COLOR_BLUE
             << "SKIP: "
             << t->get_name()
             << COLOR_RESET
             << endl;
    } else {
        tests.push(t);
    }
    return true;
}