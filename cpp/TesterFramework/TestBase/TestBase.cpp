/**
 * @name TesterFramework/TestBase.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
TestBase::TestBase() {
    debug_flag = false;
}

TestBase::TestBase(string n) {
    debug_flag = false;
    name = n;
}

TestBase::~TestBase() {
    debug("TestBase destructor executed");
}

bool TestBase::main() {
    debug("TestBase::main() executed");
    return false;
}

bool TestBase::expect(bool condition, string message) {
    if (!condition) {
        cout << "\t[" << COLOR_RED << name << COLOR_RESET << "]:"
             << COLOR_RED << " FAIL" << COLOR_RESET
             << " EXPECTATION: " << COLOR_GRAY << message << COLOR_RESET
             << endl;
        return false;
    }
    cout << "\t[" << COLOR_GREEN << name << COLOR_RESET << "]:"
         << COLOR_GREEN << " PASS" << COLOR_RESET
         << " EXPECTATION: " << COLOR_GRAY << message << COLOR_RESET
         << endl;
    return true;
}

bool TestBase::set_debug() {
    return debug_flag = true;
}

bool TestBase::debug(string msg) {
    if (debug_flag)
        cout << COLOR_BLUE
             << "\t\t("
             << name
             << ") "
             << "[DEBUG]: "
             << msg
             << COLOR_RESET
             << endl;
    return true;
}

bool TestBase::error(string msg) {
    if (debug_flag)
        cout << COLOR_RED
             << "\t\t("
             << name
             << ") "
             << "[ERROR]: "
             << msg
             << COLOR_RESET
             << endl;
    return true;
}

string TestBase::get_name() {
    return name;
}