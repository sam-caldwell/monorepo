/**
 * @name types/tests/TestValueTypes.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "../ValueTypes.h"

class TestValueTypes : public TestBase {
public:
    TestValueTypes(string n) {
        name = n;
    }

    ~TestValueTypes() {/*noop*/}

    bool test_Null() {
        debug(name + "::test_Null()");
        ValueTypes o;
        o = Null;
        return expect(ValueTypeToString(o) == "Null", "Expect ValueTypeToString() to return correct result.");
    }
    bool test_Bool() {
        debug(name + "::test_Bool()");
        ValueTypes o;
        o = Bool;
        return expect(ValueTypeToString(o) == "Bool", "Expect ValueTypeToString() to return correct result.");
    }
    bool test_Double() {
        debug(name + "::test_Double()");
        ValueTypes o;
        o = Double;
        return expect(ValueTypeToString(o) == "Double", "Expect ValueTypeToString() to return correct result.");
    }
    bool test_Float() {
        debug(name + "::test_Float()");
        ValueTypes o;
        o = Float;
        return expect(ValueTypeToString(o) == "Float", "Expect ValueTypeToString() to return correct result.");
    }
    bool test_Int() {
        debug(name + "::test_Int()");
        ValueTypes o;
        o = Int;
        return expect(ValueTypeToString(o) == "Int", "Expect ValueTypeToString() to return correct result.");
    }
    bool test_String() {
        debug(name + "::test_String()");
        ValueTypes o;
        o = String;
        return expect(ValueTypeToString(o) == "String", "Expect ValueTypeToString() to return correct result.");
    }
    bool test_Uint() {
        debug(name + "::test_Uint()");
        ValueTypes o;
        o = Uint;
        return expect(ValueTypeToString(o) == "Uint", "Expect ValueTypeToString() to return correct result.");
    }


    bool main() {
        debug(name + "::main()");
        return expect(test_Null(), "Expect test_Null OK") &&
               expect(test_Bool(), "Expect test_Bool OK") &&
               expect(test_Double(), "Expect test_Double OK") &&
               expect(test_Float(), "Expect test_Float OK") &&
               expect(test_Int(), "Expect test_Int OK") &&
               expect(test_Uint(), "Expect test_Uint OK");
    }

};/*end of class*/