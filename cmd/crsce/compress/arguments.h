#ifndef COMPRESS_ARGUMENTS_H
#define COMPRESS_ARGUMENTS_H

#include <iostream>
#include <cstring>
#include <stdexcept>
#include <unistd.h>

const std::string argShort = "-";
const std::string argLong = "--";
const std::string argColor = argLong + "color";
const std::string argDebug = argLong + "debug";
const std::string argInput = argLong + "in";
const std::string argHelp = argLong + "help";
const std::string argOutput = argLong + "out";
const std::string emptyString = "";

/*
 * An Argument class for the compress command.
 */
class Argument {
private:
    bool __color;
    bool __debug;
    bool __help;
    std::string __in;
    std::string __out;
public:
    Argument() {
        this->__color = false;
        this->__debug = false;
        this->__help = false;
        this->__in = emptyString;
        this->__out = emptyString;
    }

    void inline Parse(int argc, char *argv[]);

    bool inline color() { return this->__color; };

    bool inline debug() { return this->__debug; };

    bool inline help() { return this->__help; };

    std::string inline in() { return this->__in; };

    std::string inline out() { return this->__out; };

};


#include "arguments.cpp"

#endif
