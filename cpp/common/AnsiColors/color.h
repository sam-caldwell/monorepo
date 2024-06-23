/**
 * @name AnsiColors/color.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "AnsiCodes.h"

#include <iostream>
#include <string>

class Color {
private:
    std::string code;
public:
    Color(const std::string &code) : code(code) {}

    friend std::ostream &operator<<(std::ostream &os, const Color &color);

    friend std::ostream &operator<<(std::ostream &os, const Color &&color);
};

std::ostream &operator<<(std::ostream &os, const Color &color) {
    return os << color.code;
}

std::ostream &operator<<(std::ostream &os, const Color &&color) {
    return os << color.code;
}

// ANSI color codes
const Color reset("\033[0m");
const Color red("\033[31m");
const Color green("\033[32m");
const Color blue("\033[34m");

