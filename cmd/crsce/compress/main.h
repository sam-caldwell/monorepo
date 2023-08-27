/*
 * cmd/crsce/compress/main.cpp
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This is the CRSCE file compressor command.
 */
#include <iostream>
#include <string>
#include "usage.h"
#include "arguments.h"
#include "../../../projects/cpp/DataSource/DataSource.h"

using namespace std;

int main(int argc, char *argv[]) {
    Argument arg;
    try {
        arg.Parse(argc, argv);
    } catch (exception e) {
        cout << "Error parsing arguments. " << e.what() << endl;
        return 1;
    }

    DataSource::Abstract<DataSource::File> source(arg.in());
    DataSource::Abstract<DataSource::File> target(arg.out());
    //ToDo: Compress in -> out with flags.

}

