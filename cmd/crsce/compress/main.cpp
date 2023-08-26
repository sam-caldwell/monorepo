/*
 * cmd/crsce/compress/main.cpp
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This is the CRSCE file compressor command.
 */
#include <iostream>
#include <string>

using namespace std;

#define COMMAND_USAGE endl \
    << "CRSCE Compression Tool\n" << endl \
    << "compress -h | --help                show this message" << endl<< endl \
    << "compress --in <file> --out <file> [options]" << endl<< endl \
    << "\tOptions:" << endl \
    << "\t\t--timings   print timings for various operations" << endl \
    << "\t\t--color     use ANSI color output" << endl \
    << "\t\t--debug     print detailed messages"<< endl

bool argRequestsHelp(int argc, char *argv) {
    for (int i = 0; i < argc; i++) {
        string s(argv[i]);
        if ((s == "-h") || (s == "--help"))
        {
            return true;
        }
    }
    return false;
}

int main(int argc, char *argv[]) {
    if (argRequestsHelp(argc, argv[0])) {
        cout << COMMAND_USAGE << endl << endl;
        return 1;
    }
}

