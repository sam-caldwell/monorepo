/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 */
#ifndef VALUE_ERROR_H
#define VALUE_ERROR_H

#include <iostream>
#include "base.h"

using namespace std;

class ValueError : public runtime_error {
public:
    ValueError(std::string msg) : runtime_error(msg.c_str()) {}
};

#endif //VALUE_ERROR_H