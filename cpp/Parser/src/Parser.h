/**
 * @name Parser/src/Parser.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef Parser_H
#define Parser_H

#include <algorithm>
#include <cctype>
#include <fstream>
#include <iostream>
#include <locale>
#include <map>
#include <queue>
#include <sstream>
#include <string>
#include "tokens.h"
#include "TokenType.cpp"
#include "projects/common/system/file.h"
#include "projects/common/types/FileFormats.h"
#include "projects/common/formatter/formatter.h"
#include "projects/common/Breadcrumbs/src/Breadcrumbs.h"

using namespace std;


class Parser {
public:
    /**
     * @name Class Constructor
     * @brief destroy initial state
     *
     * @throws BadFileName exception - if filename invalid
     * @throws UnsupportedFileFormat exception -if format not supported.
     *
     * @param &fileName string reference
     * @param format FileFormats format
     */
    Parser(string fileName, FileFormats format);

    /**
     * @name Class Destructor
     * @brief destroy initial state
     */
    ~Parser() {
        if (file.is_open()) file.close();
    };

    /**
     * @name fetchNextLine
     * @brief fetch the next line from our stream, increment our lineNumber
     *        UNLESS the getNextLine flag is false, in which case we do not
     *        change the lineNumber or fetch a new line.  We simply reset the
     *        getNextLine flag to true so the current line is re-used.
     *
     * @return bool
     */
    bool fetchNextLine();

    /**
     * @name getToken
     * @brief return the value associated to a given token (key)
     *
     * @throws out_of_range exception is thrown if no such key is found.
     *
     * @param key string
     * @return string
     */
    string getToken(string key);

    /**
     * @name initIterator
     * @brief reset/init the tokens iterator to the start of the map.
     */
    inline void initIterator();

    /**
     * @name endIterator
     * @brief return whether the tokens iterator is at the end of the map
     *
     * @return bool
     */
    inline bool endIterator();

    /**
     * @name next
     * @brief return the next key:value pair from the map using an iterator.
     *
     * @throws out_of_range if the iterator is at the end of the map.
     *
     * @return pair<string,string>
     */
    pair <string, string> next();

    /**
     * @name count
     * @brief return the token count.
     *
     * @return unsigned
     */
    inline unsigned count();

#ifdef TESTING_ONLY
    /**
     * @name getLine
     * @brief return the current lineNumber.
     *
     * @return int
     */
    inline int getLine() { return lineNumber; }
#endif /* TESTING_ONLY */

#ifdef TESTING_ONLY
    /**
     * @name getType
     * @brief get the expected token type
     *
     * @return TokenType
     */
    TokenType getType() { return expectedToken; }
#endif /* TESTING_ONLY */

#ifdef TESTING_ONLY
    /**
     * @name nextLine
     * @brief set getNextLine to true
     */
    void nextLine() { getNextLine = true; }
#endif /* TESTING_ONLY */

    /**
     * @name parse
     * @brief Parse the source file into the Node graph.
     *
     * @throws MalformedFile when there is a tokenizing error.
     */
    void parse();

#ifdef TESTING_ONLY
    /**
     * @name preserveLine
     * @brief set getNextLine to false
     */
    inline void preserveLine() { getNextLine = false; }
#endif /* TESTING_ONLY */

    /**
     * @name routeHeader
     * @brief
     *
     * @return bool
     */
    inline bool routeHeader();

    /**
     * @name routeAny
     * @brief
     *
     * @return bool
     */
    inline bool routeAny();

    /**
     * @name routeMultiLineTag
     * @brief
     *
     * @param *line string pointer
     * @return bool
     */
    inline bool routeMultiLineTag(string *line);

    /**
     * @name routeSequenceKeyValue
     * @brief
     *
     * @return bool
     */
    inline bool routeSequenceKeyValue();

    /**
     * @name routeSequenceMap
     * @brief
     *
     * @return bool
     */
    inline bool routeSequenceMap();

    /**
     * @name routeSequenceItem
     * @brief
     *
     * @return bool
     */
    inline bool routeSequenceItem();

    /**
     * @name routeKeyValueTag
     * @brief
     *
     * @return bool
     */
    inline bool routeKeyValueTag();

    /**
     * @name routeKeyTagMap
     * @brief
     *
     * @return bool
     */
    inline bool routeTagMap();

private:

    /**
     * @name expectedToken
     * @brief The TokenType expected in the next parser pass.
     */
    TokenType expectedToken;

    /**
     * @name _file
     * @brief file stream object.
     */
    ifstream file;

    /**
     * @name _format
     * @brief identifies the file format (e.g. Yaml)
     */
    FileFormats format;

    /**
     * @name getNextLine
     * @brief a flag to signal whether fetchNextLine() will read a
     *        new line when called.
     */
    bool getNextLine;

    /**
     * @name indent
     * @brief a set of all indentations associated with tags.
     */
    Breadcrumbs<unsigned> indent;

    /**
     * @name index
     * @brief index of a list.
     */
    unsigned index;

    /**
     * @name line
     * @brief the current line read from the source _file stream.
     */
    string line;

    /**
    * @name lineNumber
     * @brief the current line number of the line read from file.
     */
    unsigned lineNumber;

    /**
     * @name tags
     * @brief The set of all tags discovered at a given time.
     */
    Breadcrumbs <string> tags;

    /**
     * @name tagIndent
     * @brief this is the indentation of a tag
     */
    Breadcrumbs<unsigned> tagIndent;

    /**
     * @name tokens
     * @brief a map of tokens to their value.
     */
    map <string, string> tokens;

    /**
     * @name tokenIterator
     * @brief iterator for the tokens map.
     */
    map<string, string>::iterator tokenIterator;
};

#include "constructor.cpp"
#include "count.cpp"
#include "endIterator.cpp"
#include "fetchNextLine.cpp"
#include "getToken.cpp"
#include "initIterator.cpp"
#include "next.cpp"
#include "parse.cpp"
#include "routeAny.cpp"
#include "routeHeader.cpp"
#include "routeKeyValueTag.cpp"
#include "routeMultiLineTag.cpp"
#include "routeSequenceKeyValue.cpp"
#include "routeSequenceItem.cpp"
#include "routeSequenceMap.cpp"
#include "routeTagMap.cpp"

#endif /* Parser_H */