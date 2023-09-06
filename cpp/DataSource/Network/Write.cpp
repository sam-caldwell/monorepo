/*
 * DataSource::Network::Write
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#include "Network.h"
/*
 *
 */
byte_t DataSource::Network::Write(byte_t content) {
    ssize_t bytesSent = send(socketFd, &content, sizeof(byte_t), 0);

    if (bytesSent == -1) {
        throw std::runtime_error("Error writing to network socket");
    } else if (bytesSent != sizeof(byte_t)) {
        throw std::runtime_error("Incomplete write to network socket");
    }
    return content;
}
