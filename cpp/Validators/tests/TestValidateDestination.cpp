/**
 * @name Validators/tests/TestValidateDestination.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "projects/Tester/TestBase/main.h"

#include <regex>
#include "projects/common/Validators/validators.h"

class TestValidateDestination : public TestBase {
private:
    const string validationErrorMessage="Logger destination path failed validation.";
    const string directoryTraversalError="Logger destination path cannot include directory traversal (e.g. '..')";

public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestValidateDestination(string n) { name = n; }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestValidateDestination() {}

    /**
     * @name test_valid_file_path1
     * @brief expect we can validate a simple absolute POSIX file/path destination
     *
     * @return bool
     */
    bool test_valid_absolute_file_path() {
        const string initialState = "/var/log/myLog.log";
        string destination = initialState;
        validateDestination(&destination);
        return expect(destination.compare(initialState) == 0, "Expect destination ok");
    }

    /**
     * @name test_valid_relative_file_path
     * @brief expect we can validate a simple relative POSIX file/path destination
     *
     * @return bool
     */
    bool test_valid_relative_file_path() {
        const string initialState = "./myLog.log";
        string destination = initialState;
        validateDestination(&destination);
        return expect(destination.compare(initialState) == 0, "Expect destination ok");
    }

    /**
     * @name test_directory_traversal_file_path1
     * @brief expect exception is thrown if directory traversal occurs in file path.
     *
     * @return bool
     */
    bool test_directory_traversal_file_path1() {
        const string initialState = "../../myLog.log";
        string destination = initialState;
        try {
            validateDestination(&destination);
            return expect(false, "Expected out_of_range exception.");
        } catch (out_of_range &e) {
            return expect(destination.compare(initialState) == 0, "Expect destination ok") &&
                   expect(string(e.what()).compare(validationErrorMessage) == 0,
                          "Expected error message ok (" + string(e.what()) + ")");
        }
    }

    /**
     * @name test_directory_traversal_file_path2
     * @brief expect exception is thrown if directory traversal occurs in file path.
     *
     * @return bool
     */
    bool test_directory_traversal_file_path2() {
        const string initialState = "/var/log/../../myLog.log";
        string destination = initialState;
        try {
            validateDestination(&destination);
            return expect(false, "Expected out_of_range exception.");
        } catch (out_of_range &e) {
            return expect(destination.compare(initialState) == 0, "Expect destination ok") &&
                   expect(string(e.what()).compare(directoryTraversalError) == 0,
                          "Expected error message ok (" + string(e.what()) + ")");
        }
    }

    /**
     * @name test_directory_traversal_file_url1
     * @brief expect exception is thrown if directory traversal occurs in file url.
     *
     * @return bool
     */
    bool test_directory_traversal_file_url1() {
        const string initialState = "file:///../myLog.log";
        string destination = initialState;
        try {
            validateDestination(&destination);
            return expect(false, "Expected out_of_range exception.");
        } catch (out_of_range &e) {
            return expect(destination.compare(initialState) == 0, "Expect destination ok") &&
                   expect(string(e.what()).compare(directoryTraversalError) == 0,
                          "Expected error message ok (" + string(e.what()) + ")");
        }
    }

    /**
     * @name test_directory_traversal_file_url2
     * @brief expect exception is thrown if directory traversal occurs in file url.
     *
     * @return bool
     */
    bool test_directory_traversal_file_url2() {
        const string initialState = "file:///%2e%2e/myLog.log";
        string destination = initialState;
        try {
            validateDestination(&destination);
            return expect(false, "Expected out_of_range exception.");
        } catch (out_of_range &e) {
            return expect(destination.compare(initialState) == 0, "Expect destination ok") &&
                   expect(string(e.what()).compare(directoryTraversalError) == 0,
                          "Expected error message ok (" + string(e.what()) + ")");
        }
    }

    /**
     * @name test_file_path_url
     * @brief expect we can validate a file url.
     *
     * @return bool
     */
    bool test_file_path_url() {
        const string initialState = "file:///var/log/myLog.log";
        string destination = initialState;
        validateDestination(&destination);
        return expect(destination.compare(initialState) == 0, "Expect destination ok");
    }

    /**
     * @name test_valid_syslog_url
     * @brief expect we can validate a simple two string tag set.
     *
     * @return bool
     */
    bool test_valid_syslog_url() {
        const string initialState = "syslog://127.0.0.1:514";
        string destination = initialState;
        validateDestination(&destination);
        return expect(destination.compare(initialState) == 0, "Expect destination ok");
    }

    /**
     * @name test_valid_https_url
     * @brief expect we can validate a simple two string tag set.
     *
     * @return bool
     */
    bool test_valid_https_url() {
        const string initialState = "https://127.0.0.1:8080";
        string destination = initialState;
        validateDestination(&destination);
        return expect(destination.compare(initialState) == 0, "Expect destination ok");
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_valid_absolute_file_path(), "test_valid_absolute_file_path() ok") &&
               expect(test_valid_relative_file_path(), "test_valid_relative_file_path() ok") &&
               expect(test_directory_traversal_file_path1(), "test_directory_traversal_file_path1() ok") &&
               expect(test_directory_traversal_file_path2(), "test_directory_traversal_file_path2() ok") &&
               expect(test_file_path_url(), "test_file_path_url() ok") &&
               expect(test_valid_syslog_url(), "test_valid_syslog_url() ok") &&
               expect(test_valid_https_url(), "test_valid_https_url() ok");
    }

};