/**
 * @name AnsiColors/test.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include <sstream>
#include <string>
#include <cassert>
#include <functional>
#include "color.h"

// Function to capture std::cout output
std::string captureOutput(const std::function<void()>& func) {
    std::ostringstream oss;
    std::streambuf *oldCoutBuf = std::cout.rdbuf(oss.rdbuf());
    func();
    std::cout.rdbuf(oldCoutBuf);
    return oss.str();
}

// Test function
void testColorOutput() {
    // Expected ANSI color codes
    const std::string expectedBlue = "\033[34mThis is blue text\033[0m";
    const std::string expectedRed = "\033[31mThis is red text\033[0m";
    const std::string expectedGreen = "\033[32mThis is green text\033[0m";
    const std::string expectedMixed = "\033[34mThis is blue\033[31m and red\033[32m and green\033[0m";

    // Capture output
    std::string output;

    // Test blue color
    output = captureOutput([]() { std::cout << blue << "This is blue text" << reset; });
    assert(output == expectedBlue && "Blue color test failed");

    // Test red color
    output = captureOutput([]() { std::cout << red << "This is red text" << reset; });
    assert(output == expectedRed && "Red color test failed");

    // Test green color
    output = captureOutput([]() { std::cout << green << "This is green text" << reset; });
    assert(output == expectedGreen && "Green color test failed");

    // Test mixed colors
    output = captureOutput([]() {
        std::cout << blue << "This is blue" << red << " and red" << green << " and green" << reset;
    });
    assert(output == expectedMixed && "Mixed color test failed");

    std::cout << green << "All tests passed!" << std::endl << reset;
}

int main() {
    testColorOutput();
    return 0;
}