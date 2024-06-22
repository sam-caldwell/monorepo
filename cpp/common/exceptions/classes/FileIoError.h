/**
 * @name projects/exceptions/src/FileIoError.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef FileIoError_H
#define FileIoError_H

#include <iostream>
#include "base.h"

using namespace std;

class FileIoError : public runtime_error {
public:
    FileIoError(std::string msg) : runtime_error(msg.c_str()) {}
};

#endif /* FileIoError_H */
