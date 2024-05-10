/**
 * Parsers/yaml/src/Parsers.h/validate.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef parser_validate_h
#define parser_validate_h

#include "../../../exceptions/exceptions.h"

/**
 * @name validate
 * @brief validate inputs and throw exceptions where needed.
 *
 * @throws BadFileName exception - if filename invalid
 * @throws UnsupportedFileFormat exception -if format not supported.
 *
 * @param fileName string
 * @param format ParserFileFormat format
 * @return bool
 */
bool validate(string fileName, FileFormats format) {

    if (!File::exists(fileName))
        throw BadFilename("path/filename not found");

    if (File::detectDirectoryTraversal(fileName))
        throw BadFilename("directory traversal not allowed");

    if (fileName.length() == 0)
        throw BadFilename("path/filename cannot be empty.");

    if (fileName.length() > 2048)
        throw BadFilename("path/filename cannot exceed 2048 chars");

    switch (format) {
        case Yaml:
            break; //Test passes.  Go on to the next...

        default:
            throw UnsupportedFileFormat(FileFormatToString(format));
    }
    return true;
}

#endif /* parser_validate_h */