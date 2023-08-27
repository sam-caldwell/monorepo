/*
 * DataSource
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file creates an abstraction class which can allow a consuming program
 * to connect to and interact with either a file, network or other data source.
 */
#include <iostream>
#include <fstream>
#include <cstring>
#include <stdexcept>
#include <arpa/inet.h> // For network-based operations
#include <unistd.h>
#include "../types/byte_t.h"

using namespace std;

typedef unsigned char byte_t;

namespace DataSource {
    /*
    * DataSource.
    *      This is a parent class for the data source scheme.
    */
    template<typename SourceType>
    class DataSource {
    private:
        SourceType source;

    public:
        byte_t Read(uint64_t position) {}

        byte_t Write(uint64_t position, byte_t content) {}
    };

    /*
     * Specialization for file-based data source
     */
    class File {
    private:
        fstream file;

    public:
        File(const string &filename) {
            file.open(filename, ios::binary | ios::out | ios::in);
            if (!file.is_open()) {
                throw runtime_error("File not found or couldn't be opened");
            }
        }
        ~File(){
            file.close();
        }

        byte_t Read(uint64_t position) {
            byte_t data;
            file.seekg(position);
            file.read(reinterpret_cast<char *>(&data), sizeof(byte_t));
            return data;
        }

        byte_t Write(uint64_t position, byte_t content) {
            file.seekg(position);
            file.write(reinterpret_cast<char *>(&content), sizeof(byte_t));
            return content;
        }
    };

    /*
     * Specialization for network-based data source
     */
    class Network {
    private:
        int socketFd; // Example network socket descriptor

    public:
        Network(const string &address, int port) {
            // Example code to create a network socket and connect
            socketFd = socket(AF_INET, SOCK_STREAM, 0);
            if (socketFd == -1) {
                throw runtime_error("Error creating socket");
            }

            sockaddr_in serverAddress;
            serverAddress.sin_family = AF_INET;
            serverAddress.sin_port = htons(port);
            inet_pton(AF_INET, address.c_str(), &serverAddress.sin_addr);

            if (connect(socketFd, reinterpret_cast<struct sockaddr *>(&serverAddress), sizeof(serverAddress)) == -1) {
                throw runtime_error("Error connecting to server");
            }
        }
        ~Network(){
            close(socketFd);
        }

        byte_t Read(uint64_t position) {
            byte_t data;
            ssize_t bytesRead = recv(socketFd, &data, sizeof(byte_t), 0);

            if (bytesRead == -1) {
                throw runtime_error("Error reading from network socket");
            } else if (bytesRead == 0) {
                throw runtime_error("Connection closed by the remote host");
            }
            return data;
        }

        byte_t Write(uint64_t position, byte_t content) {
            ssize_t bytesSent = send(socketFd, &content, sizeof(byte_t), 0);

            if (bytesSent == -1) {
                throw runtime_error("Error writing to network socket");
            } else if (bytesSent != sizeof(byte_t)) {
                throw runtime_error("Incomplete write to network socket");
            }
            return content;
        }
    };
};