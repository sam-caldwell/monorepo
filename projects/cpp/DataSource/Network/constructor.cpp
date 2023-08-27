/*
 * DataSource::Network::constructor
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#include "Network.h"
/*
 *
 */
DataSource::Network::Network(const std::string &address, int port) {
    // Example code to create a network socket and connect
    socketFd = socket(AF_INET, SOCK_STREAM, 0);
    if (socketFd == -1) {
        throw std::runtime_error("Error creating socket");
    }

    sockaddr_in serverAddress;
    serverAddress.sin_family = AF_INET;
    serverAddress.sin_port = htons(port);
    inet_pton(AF_INET, address.c_str(), &serverAddress.sin_addr);

    if (connect(socketFd, reinterpret_cast<struct sockaddr *>(&serverAddress), sizeof(serverAddress)) == -1) {
        throw std::runtime_error("Error connecting to server");
    }
}