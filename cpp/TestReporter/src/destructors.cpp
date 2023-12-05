/**
 * @name TestReporter/src/destructors.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include <queue>

using namespace std;

TestReporter::~TestReporter() {
    if (log.is_open()) log.close();
    if (source.is_open()) source.close();
}
