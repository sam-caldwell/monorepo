/**
 * @name projects/exceptions/src/UnsupportedFileFormat.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef UnsupportedFileFormat_H
#define UnsupportedFileFormat_H

#include <iostream>
#include "base.h"

using namespace std;

class UnsupportedFileFormat : public runtime_error {
public:
    UnsupportedFileFormat(std::string msg):runtime_error(msg.c_str()){}
};

#endif /* UnsupportedFileFormat_H */
