/*
 * DataSource::Network::Read
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#include "Network.h"
/*
 *
 */
byte_t DataSource::Network::Read() {
    byte_t data;
    ssize_t bytesRead = recv(socketFd, &data, sizeof(byte_t), 0);

    if (bytesRead == -1) {
        throw std::runtime_error("Error reading from network socket");
    } else if (bytesRead == 0) {
        throw std::runtime_error("Connection closed by the remote host");
    }
    return data;
}