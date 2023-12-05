/**
 * @name Parser/src/fetchNextLine.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Parser_fetchNextLine_H
#define Parser_fetchNextLine_H

#include "projects/common/formatter/formatter.h"
#include "projects/common/Parser/src/stripComment.cpp"
#include "projects/common/Parser/src/getIndentation.cpp"

inline bool isEmpty(string *inp) {
    string s(*inp);
    Strings::trim(&s);
    return s.empty();
}


/**
 * @name fetchNextLine
 * @brief fetch the next non-comment line from our stream, increment our
 *        lineNumber ***UNLESS*** the getNextLine flag is false, in which
 *        case we do not change the lineNumber or fetch a new line.  We
 *        simply reset the getNextLine flag to true so the current line is
 *        re-used.
 *
 * @return bool (true: new line read, false: no new line or same line)
 */
bool Parser::fetchNextLine() {
    if (getNextLine) {
        line = "";
        while (isEmpty(&line) && getline(file, line)) {
            lineNumber++;
            /**
             * We should strip Comments on anything except a line we expect
             * to belong to a multi-line value.
             */
            if (stripComment(&line)) continue;
        }
        unsigned i = getIndentation(&line);
        if (i < indent.current()) {
            index = 0;
            expectedToken = Any;
        }
        indent.push(i);
        return true;
    } else {
        //Do not adjust indent or lastIndent
        getNextLine = true; //we only do this one time.
        return false;
    }
}

#endif /* Parser_fetchNextLine_H */