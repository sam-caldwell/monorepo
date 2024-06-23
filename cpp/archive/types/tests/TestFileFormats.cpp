/**
 * @name types/tests/TestFileFormats.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#pragma GCC diagnostic ignored "-Wnonportable-include-path"
#include <iostream>
#include "../FileFormats.h"


class TestFileFormats : public TestBase {
public:
    TestFileFormats(string n) {
        name = n;
    }

    ~TestFileFormats() {/*noop*/}

    bool test_format_binary() {
        FileFormats o = Binary;
        return expect(o==Binary, "Expect Binary ok");
    }

    bool test_format_csv() {
        FileFormats o = CSV;
        return expect(o==CSV, "Expect CSV ok");
    }

    bool test_format_Json() {
        FileFormats o = Json;
        return expect(o==Json, "Expect Json ok");
    }

    bool test_format_Text() {
        FileFormats o = Text;
        return expect(o==Text, "Expect Text ok");
    }

    bool test_format_Xml() {
        FileFormats o = Xml;
        return expect(o==Xml, "Expect Xml ok");
    }

    bool test_format_Yaml() {
        FileFormats o = Yaml;
        return expect(o==Yaml, "Expect Yaml ok");
    }

    bool main() {
        debug(name + "::main()");
        return expect(test_format_binary(), "Expect test_format_binary OK") &&
                expect(test_format_csv(), "Expect test_format_csv OK") &&
                expect(test_format_Json(), "Expect test_format_Json OK") &&
                expect(test_format_Text(), "Expect test_format_Text OK") &&
                expect(test_format_Xml(), "Expect test_format_Xml OK") &&
                expect(test_format_Yaml(), "Expect test_format_Yaml OK");
    }

};/*end of class*/