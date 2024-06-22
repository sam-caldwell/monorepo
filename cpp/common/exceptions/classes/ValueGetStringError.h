/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 */
#ifndef VALUE_TO_STRING_ERROR_H
#define VALUE_TO_STRING_ERROR_H

#include <iostream>
#include "base.h"

using namespace std;

class ValueGetStringError : public BaseException {
public:
    ValueGetStringError(){};

    ValueGetStringError(string msg){
        message = msg;
    };

};

#endif //VALUE_TO_STRING_ERROR_H