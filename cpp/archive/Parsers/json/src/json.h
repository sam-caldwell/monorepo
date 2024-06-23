/**
 * @name Parsers/json/src/Parsers.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef MONOREPO_JSON_H
#define MONOREPO_JSON_H

#include <iostream>
#include <string>
#include <map>
#include <vector>
#include <cctype>

enum class JsonValueType {
    STRING,
    NUMBER,
    OBJECT,
    ARRAY,
    BOOL,
    NULL_VALUE
};

struct JsonValue {
    JsonValueType type;
    std::string string_value;
    double number_value;
    std::map<std::string, JsonValue> object_value;
    std::vector<JsonValue> array_value;
    bool bool_value;
};

class JsonParser {
public:
    JsonParser(const std::string& json) : json(json), index(0) {}

    JsonValue parse() {
        skipWhiteSpace();
        return parseValue();
    }

private:
    std::string json;
    size_t index;

    void skipWhiteSpace() {
        while (index < json.length() && std::isspace(json[index])) {
            index++;
        }
    }

    char currentChar() {
        return json[index];
    }

    char nextChar() {
        return json[index + 1];
    }

    JsonValue parseValue() {
        skipWhiteSpace();
        char current = currentChar();
        if (current == '{') {
            return parseObject();
        } else if (current == '[') {
            return parseArray();
        } else if (current == '"') {
            return parseString();
        } else if (std::isdigit(current) || current == '-') {
            return parseNumber();
        } else if (current == 't' || current == 'f') {
            return parseBool();
        } else if (current == 'n') {
            return parseNull();
        } else {
            throw std::runtime_error("Invalid JSON");
        }
    }

    JsonValue parseString() {
        expect('"');
        std::string value;
        while (nextChar() != '"') {
            value += currentChar();
            index++;
        }
        expect('"');
        return {JsonValueType::STRING, value};
    }

    JsonValue parseNumber() {
        std::string value;
        while (std::isdigit(currentChar()) || currentChar() == '.' || currentChar() == '-') {
            value += currentChar();
            index++;
        }
        return {JsonValueType::NUMBER, "", std::stod(value)};
    }

    JsonValue parseObject() {
        expect('{');
        std::map<std::string, JsonValue> obj;
        while (currentChar() != '}') {
            expect('"');
            std::string key = parseString().string_value;
            expect(':');
            JsonValue value = parseValue();
            obj[key] = value;
            if (currentChar() == ',') {
                index++;
            }
        }
        expect('}');
        return {JsonValueType::OBJECT, "", 0.0, obj};
    }

    JsonValue parseArray() {
        expect('[');
        std::vector<JsonValue> arr;
        while (currentChar() != ']') {
            JsonValue value = parseValue();
            arr.push_back(value);
            if (currentChar() == ',') {
                index++;
            }
        }
        expect(']');
        return {JsonValueType::ARRAY, "", 0.0, {}, arr};
    }

    JsonValue parseBool() {
        if (currentChar() == 't') {
            expectString("true");
            return {JsonValueType::BOOL, "", 0.0, {}, {}, true};
        } else {
            expectString("false");
            return {JsonValueType::BOOL, "", 0.0, {}, {}, false};
        }
    }

    JsonValue parseNull() {
        expectString("null");
        return {JsonValueType::NULL_VALUE};
    }

    void expect(char expected) {
        if (currentChar() != expected) {
            throw std::runtime_error("Unexpected character");
        }
        index++;
    }

    void expectString(const std::string& expected) {
        size_t len = expected.length();
        for (size_t i = 0; i < len; i++) {
            if (json[index + i] != expected[i]) {
                throw std::runtime_error("Unexpected character");
            }
        }
        index += len;
    }
};

int main() {
    std::string json_string = "{\"name\": \"John\", \"age\": 30, \"city\": \"New York\", \"is_student\": false, \"grades\": [90, 85, 95]}";
    JsonParser parser(json_string);
    JsonValue parsed_json = parser.parse();

    // Accessing parsed JSON
    std::cout << "Name: " << parsed_json.object_value["name"].string_value << std::endl;
    std::cout << "Age: " << parsed_json.object_value["age"].number_value << std::endl;
    std::cout << "City: " << parsed_json.object_value["city"].string_value << std::endl;
    std::cout << "Is Student: " << std::boolalpha << parsed_json.object_value["is_student"].bool_value << std::endl;
    std::cout << "Grades: ";
    for (const auto& grade : parsed_json.object_value["grades"].array_value) {
        std::cout << grade.number_value << " ";
    }
    std::cout << std::endl;

    return 0;
}


#endif //MONOREPO_JSON_H
