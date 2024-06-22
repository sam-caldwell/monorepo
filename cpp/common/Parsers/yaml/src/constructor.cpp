/**
 * Parsers/yaml/src/Parsers.h/constructor.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "validate.cpp"

/**
 * @name Class Constructor
 * @brief destroy initial state
 *
 * @throws BadFileName exception - if filename invalid
 * @throws UnsupportedFileFormat exception -if format not supported.
 * @throws FileIoError exception if there is an error opening the file/stream.
 *
 * @param fileName string reference
 * @param format FileFormats format
 */
Parser::Parser(string fileName, FileFormats fileFormat = Yaml) {
    //PRINT_MSG("Parsers: filename:" << fileName << " format: " << fileFormat)

    validate(fileName, fileFormat);

    format = fileFormat;

    file.open(fileName, std::ifstream::in);

    if (!file) throw FileIoError(fileName);

    lineNumber = 0;

    index = 0;

    line = "";

    expectedToken = Header; //Expect the Header as our bootstrap token.

    getNextLine = true;


};
