/*
 * DataSource::Network::constructor
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */
#include "Network.h"
#include <string>
#include <stdexcept>
#include "../../../../projects/cpp/misc/parseNetConnect.h"

/*
 *
 */
DataSource::Network::Network(const std::string &ConnectionString) {
    try {
        NetworkAddress address(ConnectionString);

        socketFd = socket(AF_INET, SOCK_STREAM, 0);
        if (socketFd == -1) {
            throw std::runtime_error("Error creating socket");
        }

        sockaddr_in serverAddress;
        serverAddress.sin_family = AF_INET;
        serverAddress.sin_port = htons(address.Port());
        inet_pton(AF_INET, address.Address().c_str(), &serverAddress.sin_addr);

        if (connect(socketFd, reinterpret_cast<struct sockaddr *>(&serverAddress), sizeof(serverAddress)) == -1) {
            throw std::runtime_error("Error connecting to server");
        }

    } catch (std::exception e) {
        throw std::runtime_error("DataSource: " + static_cast<std::string>( e.what()));
    }
}