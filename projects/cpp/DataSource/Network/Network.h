/*
 * DataSource::Network
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#ifndef DATASOURCE_NETWORK_H
#define DATASOURCE_NETWORK_H

#include <cstring>
#include <stdexcept>
#include <arpa/inet.h> // For network-based operations
#include <unistd.h>
#include "../../types/byte_t.h"
#include "../constants.h"

namespace DataSource {
/*
 * Specialization for network-based data source
 */
    class Network {
    private:
        int socketFd; // Example network socket descriptor

    public:
        Network(const std::string &address, int port);

        ~Network();

        byte_t Read();

        byte_t Write(byte_t content);

        // Not implemented.
        uint64_t Size(){return SizeNotImplemented;};
    };
};

#include "Network.cpp"

#endif