/**
 * @name Parsers/tests/TestParserInitialState.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/Tester/TestBase/main.h"
#include "projects/common/Parser/src/Parser.h"
#include "projects/common/exceptions/exceptions.h"

class TestParserInitialState : public TestBase {
private:
    const string GoodYamlFile = "data_good.yaml";
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParserInitialState(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserInitialState() {}

    /**
     * @name test_parser_default_file_format
     * @brief Can we create an object with the default file format.
     *
     * @return bool
     */
    bool test_parser_default_file_format() {
        Parser p(GoodYamlFile);
        Parser q(GoodYamlFile, Yaml);
        return true;
    }

    /**
     * @name test_parser_supported_file_formats
     * @brief Can we create good objects ok and catch exceptions on bad objects?
     * @return bool
     */
    bool test_parser_supported_file_formats() {
        try {
            Parser r(GoodYamlFile, Binary);
            return false;
        } catch (UnsupportedFileFormat) {
            return true;
        }
        try {
            Parser r(GoodYamlFile, CSV);
            return false;
        } catch (UnsupportedFileFormat) {
            return true;
        }
        try {
            Parser r(GoodYamlFile, Json);
            return false;
        } catch (UnsupportedFileFormat) {
            return true;
        }
        try {
            Parser r(GoodYamlFile, Text);
            return false;
        } catch (UnsupportedFileFormat) {
            return true;
        }
        try {
            Parser r(GoodYamlFile, Xml);
            return false;
        } catch (UnsupportedFileFormat) {
            return true;
        }
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_parser_default_file_format(), "Expect test_parser_default_file_format() ok") &&
               expect(test_parser_supported_file_formats(), "Expect test_parser_supported_file_formats() ok");
    }
};
