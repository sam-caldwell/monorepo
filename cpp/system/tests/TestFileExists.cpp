/**
 * @name system/tests/TestFileExists.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <cstdio>
#include <fstream>
#include <iostream>
#include <stdlib.h>

#include "projects/common/system/file.h"

class TestFileExists : public TestBase {
private:
    char const *e = "TEST_ENV_VAR";
    char const *v = "TEST_VALUE";
    const string valid_file = "system.file_exists.txt";
    const string invalid_file = "/foomarnonexistent.file.zzz";

public:
    TestFileExists(string n) {
        name = n;
    }

    ~TestFileExists() {
        std::remove(valid_file.c_str());
    }

    bool create_file() {
        fstream data;
        data.open(valid_file, fstream::out);
        data << "this is a valid file";
        data.close();
        return true;
    }

    bool main() {
        debug(name + "::main()");
        return expect(create_file(), "Expect we can create a test data file.") &&
               expect(File::exists(valid_file), "Expect exists works for file which is present") &&
               expect(!File::exists(invalid_file), "Expect non-existent file returns false");
    }

};/*end of class*/