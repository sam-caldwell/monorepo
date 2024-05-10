/**
 * Parsers/yaml/src/Parsers.h/extractTagOnly.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef Parser_extractTagOnly_H
#define Parser_extractTagOnly_H

/**
 * @name extractTagOnly
 * @brief Remove any non-tag characters (e.g. trailing ":" or ": |"), etc.
 *        from the current line and return only the tag identifier.
 *
 * @note we know that any | char must follow a : char, so on detecting : we
 *       just nuke everything from the : to the end of the string and all cases
 *       are happy.  Likewise for a "- value" we just detect - and delete it then
 *       trim the spaces.  Because left spaces are removed - should be the first
 *       character so it's rather easy to change to " " then trim again.
 *
 * @param *line string pointer
 * @return string
 */
string extractTagOnly(string *line) {
    string s = Strings::ltrim(*line);
    s[0] = (s[0] == '-') ? ' ' : s[0];
    Strings::ltrim(s);
    for (int i = 0; i < s.length(); i++)
        if (s[i] == ':')
            s.erase(i, s.length() - 1);
    Strings::ltrim(&s);
    return s;
}

#endif /* Parser_extractTagOnly_H */