/**
 * @name projects/exceptions/src/TypeError.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef TypeError_H
#define TypeError_H

#include <iostream>
#include "base.h"

using namespace std;

class TypeError : public runtime_error {
public:
    TypeError(std::string msg):runtime_error(msg.c_str()){}
};

#endif /* TypeError_H */
