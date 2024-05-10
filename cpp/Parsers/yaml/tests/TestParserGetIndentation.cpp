/**
 * @name Parsers/tests/TestParserGetIndentation.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../../../Tester/Tester/main.h"
#include "../src/getIndentation.cpp"

class TestParserGetIndentation : public TestBase {
private:
    const string GoodYamlFile = "data_good.yaml";
    Parser *p;
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParserGetIndentation(string n) {
        name = n;
        p = new Parser(GoodYamlFile, Yaml);
    }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserGetIndentation() {
        if (p) delete p;
    }

    /**
     * @name test_check_indent
     * @brief check getIndentation() against the given input
     *
     * @param text string
     * @param i int
     * @return bool
     */
    bool test_check_indent(string text, int i) {
        return expect(getIndentation(&text) == i, "Expect getIndentation ok (" + text + "): " + to_string(i));
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_check_indent("foo", 0), "01.Expect test_check_indent() ok") &&
               expect(test_check_indent("  foo", 2), "02.Expect test_check_indent() ok") &&
               expect(test_check_indent("foo bar", 0), "03.Expect test_check_indent() ok") &&
               expect(test_check_indent("  foo bar", 2), "04.Expect test_check_indent() ok") &&
               expect(test_check_indent("foo  ", 0), "05.Expect test_check_indent() ok") &&
               expect(test_check_indent("  foo  ", 2), "06.Expect test_check_indent() ok");
    }
};
