/*
 * tools/monorepo-version/main.cpp
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This command will print the monorepo version.
 */
#include <iostream>
#include "../../../projects/cpp/monorepo.h"

#ifndef MONOREPO_VERSION_MAJOR
#define MONOREPO_VERSION_MAJOR -1
#endif

int main() {
    std::cout << MONOREPO_VERSION_MAJOR << std::endl;
    return 0;
}
