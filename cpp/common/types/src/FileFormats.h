/**
 * @name projects/types/src/FileFormats.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef FileFormats_H
#define FileFormats_H

enum FileFormats {
    Binary, //Any file parsed by byte position.
    CSV,    //Any comma-delimited file
    Json,   //Json content
    Text,   //Any free-form text
    Xml,    //XML
    Yaml,   //YAML
};

#endif /* FileFormats_H */
