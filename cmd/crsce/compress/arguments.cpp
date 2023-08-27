/*
 *
 */
#include "arguments.h"

/*
 * Parse the command line arguments and configure
 * the Argument class.
 */
void Argument::Parse(int argc, char *argv) {
    std::string currArg;
    std::string prevArg;
    bool expectValue = false;

    for (int i = 0; i < argc; i++) {
        currArg = argv[i];
        if (expectValue) {
            if (currArg.substr(0, 1) == argShort) {
                throw std::runtime_error("expected value");
            }
            if (prevArg == argInput) {
                this->__in = currArg;
                prevArg = emptyString;
            }
            if (prevArg == argOutput) {
                this->__out = currArg;
                prevArg = emptyString;
            }
        } else {
            if ((currArg == argInput) || (currArg == argOutput)) {
                expectValue = true;
                prevArg = currArg;
            } else {
                expectValue = false;
                if (currArg == argColor) {
                    this->__color = true;
                    continue;
                }
                if (currArg == argDebug) {
                    this->__debug = true;
                    continue;
                }
                if (currArg == argHelp) {
                    this->__help = true;
                    continue;
                }
            }
        }
    }
};
