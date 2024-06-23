/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 */
#ifndef VALUE_ASSIGNMENT_ERROR_H
#define VALUE_ASSIGNMENT_ERROR_H

#include <iostream>

using namespace std;

class ValueAssignmentError : public BaseException {
public:
    ValueAssignmentError() {};

    ValueAssignmentError(string &msg) {
        message = msg;
    }

    ValueAssignmentError(string msg) {
        message = msg;
    };
};

#endif //VALUE_ASSIGNMENT_ERROR_H