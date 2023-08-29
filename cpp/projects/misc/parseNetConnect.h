/*
 *
 */
#include <string>
#include <arpa/inet.h>
#include <stdexcept>

const std::string connectStringDelimiter = ":";

class NetworkAddress {
private:
    std::string address;
    in_port_t port;

public:
    NetworkAddress(const std::string &ConnectionString,
                   const std::string defaultAddress = "127.0.0.1",
                   const uint16_t defaultPort = 0) {

        size_t delimiterPos = ConnectionString.find(connectStringDelimiter);
        if (delimiterPos != std::string::npos) {
            address = address.substr(0, delimiterPos++);
            int intValue = std::stoi(ConnectionString.substr(delimiterPos));
            if ((intValue < 0) || (intValue > UINT16_MAX)) {
                std::runtime_error("value is out of range (expect 0-65535)");
            }
            port = static_cast<uint16_t>(intValue);
        } else {
            // Default values in case no delimiter is found
            address = defaultAddress;
            port = defaultPort;
        }
    }

    std::string Address() {
        return address;
    }

    in_port_t Port() {
        return port;
    }
};