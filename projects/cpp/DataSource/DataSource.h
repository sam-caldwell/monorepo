/*
 * DataSource::Root
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file creates an abstraction class which can allow a consuming program
 * to connect to and interact with either a file, network or other data source.
 */
#ifndef DATASOURCE_H
#define DATASOURCE_H

#include <iostream>
#include <fstream>
#include <cstring>
#include <stdexcept>
#include <arpa/inet.h> // For network-based operations
#include <unistd.h>
#include "../types/byte_t.h"
#include "constants.h"

namespace DataSource {
    /*
     * DataSource::Root is a template for abstracting both
     * file and network data sources.
     */

    template<typename SourceType>
    class Root {

    private:
        SourceType data;

    public:

        byte_t Read(uint64_t position) {
            return this->data.Read(position);
        }

        byte_t Write(uint64_t position, byte_t content) {
            return this->data.Write(position, content);
        }

        uint64_t Size() {
            return SizeNotImplemented;
            /* if not implemented, it will return -1. */
        }
    };
};

#include "File/File.h"
#include "Network/Network.h"

#endif