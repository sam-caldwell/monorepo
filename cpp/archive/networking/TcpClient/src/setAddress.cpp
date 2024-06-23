/**
 * @name networking/TcpClient/src/setAddress.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

namespace networking {
    /**
     * @name setAddress
     * @brief set the Ip address and port
     *
     * @throws runtime_error exception if hostname cannot be resolved.
     *
     * @param address string
     * @param port int
     */
    void TcpClient::setAddress(const std::string &address, int port) {

        const int inetSuccess = inet_aton(address.c_str(), &_server.sin_addr);

        if (!inetSuccess) {
            /*
             * inet_addr failed to parse the given address, so it must possibly be a hostname/fqdn.
             * We will attempt to resolve the hostname/fqdn.
             */
            struct hostent *host;
            struct in_addr **addrList;
            if ((host = gethostbyname(address.c_str())) == nullptr)
                throw std::runtime_error("Failed to resolve hostname");
            addrList = (struct in_addr **) host->h_addr_list;
            _server.sin_addr = *addrList[0];
        }
        _server.sin_family = AF_INET;
        _server.sin_port = htons(port);
    }
}