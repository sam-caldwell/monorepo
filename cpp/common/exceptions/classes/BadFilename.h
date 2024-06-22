/**
 * @name projects/exceptions/src/BadFilename.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef BadFilename_H
#define BadFilename_H

#include <iostream>
#include "base.h"

using namespace std;

class BadFilename : public runtime_error {
public:
    BadFilename(std::string msg):runtime_error(msg.c_str()){}
};

#endif /* BadFilename_H */
