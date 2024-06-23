/**
 * @name Parsers/tests/TestParserTokenType.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../../../../Tester/Tester/main.h"
#include "../src/TokenType.cpp"
#include "../../../exceptions/exceptions.h"

class TestParserTokenType : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParserTokenType(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserTokenType() {}

    bool test_token_type_Undefined() {
        TokenType o = Any;
        return expect(o == Any, "expect Any");
    }

    bool test_token_type_Header() {
        TokenType o = Header;
        return expect(o == Header, "expect Header");
    }

    bool test_token_type_KeyValueTag() {
        TokenType o = KeyValueTag;
        return expect(o == KeyValueTag, "expect KeyValueTag");
    }

    bool test_token_type_MultiLineTag() {
        TokenType o = MultiLineTag;
        return expect(o == MultiLineTag, "expect MultiLineTag");
    }

    bool test_token_type_MultiLineValue() {
        TokenType o = MultiLineValue;
        return expect(o == MultiLineValue, "expect MultiLineValue");
    }

    bool test_token_type_SequenceItem() {
        TokenType o = SequenceItem;
        return expect(o == SequenceItem, "expect SequenceItem");
    }

    bool test_token_type_SequenceKeyValue() {
        TokenType o = SequenceKeyValue;
        return expect(o == SequenceKeyValue, "expect SequenceKeyValue");
    }

    bool test_token_type_SequenceMap() {
        TokenType o = SequenceMap;
        return expect(o == SequenceMap, "expect SequenceMap");
    }

    bool test_token_type_TagMap() {
        TokenType o = TagMap;
        return expect(o == TagMap, "expect TagMap");
    }

    bool test_token_type_Value() {
        TokenType o = Value;
        return expect(o == Value, "expect Value");
    }


    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_token_type_Undefined(), "Expect test_token_type_Undefined() ok") &&
               expect(test_token_type_Header(), "Expect test_token_type_Header() ok") &&
               expect(test_token_type_KeyValueTag(), "Expect test_token_type_KeyValueTag() ok") &&
               expect(test_token_type_MultiLineTag(), "Expect test_token_type_MultiLineTag() ok") &&
               expect(test_token_type_MultiLineValue(), "Expect test_token_type_MultiLineValue() ok") &&
               expect(test_token_type_SequenceKeyValue(), "Expect test_token_type_SequenceTag() ok") &&
                expect(test_token_type_SequenceItem(), "Expect test_token_type_SequenceItem() ok") &&
                expect(test_token_type_SequenceMap(), "Expect test_token_type_SequenceMap() ok") &&
                expect(test_token_type_TagMap(), "Expect test_token_type_TagMap() ok") &&
               expect(test_token_type_Value(), "Expect test_token_type_Value() ok");
    }
};
