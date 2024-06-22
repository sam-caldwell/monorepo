/**
 * @name Validators/tests/TestValidateTags.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "../../networking/TcpCommon/src/PipeResponse/main.h"

#include <regex>
#include "../validators.h"

class TestValidateTags : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestValidateTags(string n) { name = n; }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestValidateTags() {}

    /**
     * @name test_valid_single_string_tags
     * @brief expect we can validate a simple two string tag set.
     *
     * @return bool
     */
    bool test_valid_single_string_tags() {
        const string initialState = "\"Foo\"";
        string tags = initialState;
        validateTags(&tags);
        return expect(tags.compare(initialState) == 0, "Expect tags ok");
    }

    /**
     * @name test_valid_single_string_tags
     * @brief expect we can validate a simple two string tag set.
     *
     * @return bool
     */
    bool test_valid_string_tags_pair() {
        const string initialState = "\"Foo\",\"bar\"";
        string tags = initialState;
        validateTags(&tags);
        return expect(tags.compare(initialState) == 0, "Expect tags ok");
    }

    /**
     * @name test_valid_string_tags
     * @brief expect we can validate a simple two string tag set.
     *
     * @return bool
     */
    bool test_valid_string_tags() {
        const string initialState = "\"Foo\",\"bar\",\"m00\",\"m4r\",\"s-o-0\",\"f_a.r1_\"";
        string tags = initialState;
        validateTags(&tags);
        return expect(tags.compare(initialState) == 0, "Expect tags ok");
    }
    /**
     * @name test_valid_key_value_single_tag
     * @brief expect we can validate a single key-value tag.
     *
     * @return bool
     */
    bool test_valid_key_value_single_tag(){
        const string initialState = "\"Foo:bar\"";
        string tags = initialState;
        validateTags(&tags);
        return expect(tags.compare(initialState) == 0, "Expect tags ok");
    }
    /**
     * @name test_valid_key_value_tag_pair
     * @brief expect we can validate a pair of key-value tags.
     *
     * @return bool
     */
    bool test_valid_key_value_tag_pair(){
        const string initialState = "\"Foo:bar\",\"Moo:mar\"";
        string tags = initialState;
        validateTags(&tags);
        return expect(tags.compare(initialState) == 0, "Expect tags ok");
    }
    /**
     * @name test_valid_key_value_tags1
     * @brief expect we can validate a simple two string tag set.
     *
     * @return bool
     */
    bool test_valid_key_value_tags1() {
        const string initialState = "\"Foo:bar\",\"moo:mar\",\"soo:far\"";
        string tags = initialState;
        validateTags(&tags);
        return expect(tags.compare(initialState) == 0, "Expect tags ok");
    }

    /**
     * @name test_valid_key_value_tags2
     * @brief expect we can validate a simple two string tag set.
     *
     * @return bool
     */
    bool test_valid_key_value_tags2() {
        const string initialState = "\"Foo:bar\",\"m00:m4r\",\"s-o-0:f_a.r1_\"";
        string tags = initialState;
        validateTags(&tags);
        return expect(tags.compare(initialState) == 0, "Expect tags ok");
    }


    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_valid_single_string_tags(), "test_valid_single_string_tags() ok") &&
               expect(test_valid_string_tags_pair(), "test_valid_string_tags_pair() ok") &&
               expect(test_valid_key_value_single_tag(), "test_valid_key_value_single_tag() ok") &&
               expect(test_valid_string_tags(), "test_valid_string_tags() ok") &&
               expect(test_valid_key_value_tags1(), "test_valid_key_value_tags1() ok") &&
               expect(test_valid_key_value_tags2(), "test_valid_key_value_tags2() ok");
    }

};