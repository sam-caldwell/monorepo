/**
 * @name projects/types/src/FileFormatToString.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef FileFormatToString_H
#define FileFormatToString_H
#pragma GCC diagnostic ignored "-Wnonportable-include-path"

#include "FileFormats.h"
#include "../../exceptions/exceptions.h"

inline string FileFormatToString(FileFormats f) {
    switch (f) {
        case Binary:
            return "Binary";
            break;
        case CSV:
            return "CSV";
            break;
        case Json:
            return "Json";
            break;
        case Text:
            return "Text";
            break;
        case Xml:
            return "Xml";
            break;
        case Yaml:
            return "Yaml";
            break;
        default:
            UnsupportedFileFormat(to_string((int) f));
            break;
    }
}

#endif /* FileFormatToString_H */
