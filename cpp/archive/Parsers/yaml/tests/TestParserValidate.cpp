/**
 * @name Parsers/tests/TestParserValidate.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../../../../Tester/Tester/main.h"
#include "../../../../types/src/FileFormats.h"
#include "../src/validate.cpp"
#include "../../../../system/file.h"

class TestParserValidate : public TestBase {
private:
    const string GoodYamlFile = "data_good.yaml";
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParserValidate(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserValidate() {}

    /**
     * @name test_validation[FileFormats, bool]
     * @brief Test filename validation with known good format
     *
     * @param s string
     * @param onException bool
     * @return bool
     */
    bool test_validation(string s, bool onException) {
        try {
            bool result = validate(s, Yaml);
            cout << "validate(" << s << "): " << (result ? "true" : "false") << endl;
            return result;
        } catch (BadFilename &e) {
            cout << "BadFilename exception on " << s << " e:" << string(e.what()) << endl;
            return onException;
        } catch (UnsupportedFileFormat &e){
            cout << "UnsupportedFileFormat exception on " << s << " e:" << string(e.what()) << endl;
        }
        return false;
    }

    /**
     * @name test_validation[FileFormats, bool]
     * @brief Test format validation with known good filename
     *
     * @param s
     * @param onException
     * @return bool
     */
    bool test_validation(FileFormats s, bool onException) {
        try {
            return validate("/etc/passwd", s);
        } catch (UnsupportedFileFormat) {
            return onException;
        }
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_validation("../etc/passwd", true), "Expect test_validation() ok[1]") &&
               expect(test_validation("./etc/passwd", true), "Expect test_validation() ok[2]") &&
               expect(test_validation(GoodYamlFile, false), "Expect test_validation() ok[3]") &&
               expect(test_validation("/etc/../passwd", true), "Expect test_validation() ok[4]") &&
               expect(test_validation("./etc/..", true), "Expect test_validation() ok[5]") &&
               expect(!test_validation("non.existent.bad.file", false), "Expect test_validation() ok[6]") &&
               expect(test_validation(Binary, true), "Expect test_validation(Binary) ok") &&
               expect(test_validation(CSV, true), "Expect test_validation(CSV) ok") &&
               expect(test_validation(Json, true), "Expect test_validation(Json) ok") &&
               expect(test_validation(Text, true), "Expect test_validation(Text) ok") &&
               expect(test_validation(Xml, true), "Expect test_validation(Xml) ok") &&
               expect(test_validation(Yaml, false), "Expect test_validation(Yaml) ok");
    }
};
