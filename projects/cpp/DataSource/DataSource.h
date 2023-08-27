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
#include <mutex>

namespace DataSource {
    /*
     * DataSource::Root is a template for abstracting both
     * file and network data sources.
     */

    template<typename SourceType>
    class Abstract {

    private:
        SourceType *data;
        std::mutex dataMutex; // Mutex for synchronization

    public:
        Abstract(std::string connect){
            data = new SourceType(connect);
        };
        ~Abstract(){
            delete data;
        };

        byte_t Read(uint64_t position) {
            std::lock_guard<std::mutex> lock(dataMutex); // Lock during Read
            return data->Read(position);
        }

        byte_t Write(uint64_t position, byte_t content) {
            std::lock_guard<std::mutex> lock(dataMutex); // Lock during Read
            return data->Write(position, content);
        }

        uint64_t Size() {
            std::lock_guard<std::mutex> lock(dataMutex); // Lock during Read
            return SizeNotImplemented;
            /* if not implemented, it will return -1. */
        }
    };
};

#include "File/File.h"
#include "Network/Network.h"

#endif