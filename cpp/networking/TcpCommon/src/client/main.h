/**
 * @name networking/TcpCommon/src/client/main.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef networking_TcpCommon_Client_H
#define networking_TcpCommon_Client_H

#include <atomic>
#include <cstdio>
#include <cstring>
#include <cerrno>
#include <iostream>
#include <functional>
#include <mutex>
#include <stdexcept>
#include <string>
#include <sys/socket.h>
#include <thread>
#include <unistd.h>
#include "projects/networking/TcpCommon/src/messages.h"
#include "projects/networking/TcpCommon/src/constants.h"
#include "projects/networking/TcpCommon/src/PipeResponse/main.h"
#include "projects/networking/TcpCommon/src/clientEvent.h"
#include "projects/networking/TcpCommon/src/FileDescriptor.h"

namespace networking {
    /**
     * @class Client
     * @brief this is the common network client component.
     */
    class Client {
        /**
         * @name ClientEventHandler
         * @brief is
         */
        using ClientEventHandler = std::function<void(const Client &, ClientEvent, const std::string &)>;

    public:
        /**
         * @name class constructor
         * @brief set up state.
         *
         * @param fileDescriptor
         */
        Client(int fileDescriptor);

        /**
         * @name class destructor
         * @brief cleanup state
         *
         * @param other
         */
        ~Client();

        /**
         * @name operator ==
         * @brief Create an operator to test lhs==rhs
         * @param other Client reference
         * @return bool
         */
        bool operator==(const Client &other) const;

        /**
         * @name operator !=
         * @brief Create an operator to test lhs!=rhs
         * @param other Client reference
         * @return bool
         */
        bool operator!=(const Client &other) const;

        /**
         * @name setIp
         * @brief set the ip address.
         * @param ip string
         */
        inline void setIp(const std::string &ip);

        /**
         * @name getIp
         * @brief return the current Ip address string
         * @return string
         */
        inline std::string getIp() const;

        /**
         * @name isConnected
         * @brief determines if the connection is connected
         * @return bool
         */
        inline bool isConnected() const;

        /**
         * @name setEventsHandler
         * @brief set the event handler callback function.
         *
         * @param eventHandler event handler function pointer.
         */
        inline void setEventsHandler(const ClientEventHandler &eventHandler);

        /**
         * @name publishEvent
         * @brief publish an event to the event handler
         *
         * @param clientEvent
         * @param msg
         */
        void publishEvent(ClientEvent clientEvent, const std::string &msg = "");

        /**
         * @name startListen
         * @brief span a listener thread.
         */
        void startListen();

        /**
         * @name send (c-string)
         * @brief send a message of a given size to the destination server.
         *
         * @param msg - const char (null-terminated c string)
         * @param msgSize message size (length)
         */
        void send(const char *msg, size_t msgSize) const;

        /**
         * @name send (string)
         * @brief send a message of a given size to the destination server.
         *
         * @param msg - std::string by reference
         */
        inline void send(std::string *msg) const;

        /**
         * @name close
         * @brief close the connection
         */
        void close();

        /**
         * @name print
         * @brief print a status of our current connection.
         */
        void print() const;

    private:
        /**
         * @name SocketFileDescriptor
         * @brief our network socket file descriptor
         */
        FileDescriptor SocketFileDescriptor;
        /**
         * @name _ip
         * @brief a string ip address or hostname
         */
        std::string _ip = "";
        /**
         * @name _isConnected
         * @brief our connection-state flag.
         */
        std::atomic<bool> _isConnected;
        /**
         * @name *_receiveThread
         * @brief a pointer to our receiver thread.
         */
        std::thread *_receiveThread = nullptr;
        /**
         * @name _eventHandlerCallback
         * @brief a function pointer to our event handler callback.
         */
        ClientEventHandler _eventHandlerCallback;

        /**
         * @name setConnected
         * @brief set the connected status.
         *
         * @param flag
         */
        void setConnected(bool flag) {
            _isConnected = flag;
        }

        /**
         * @name receiveTask
         * @brief receive a message.
         */
        void receiveTask();

        /**
         * @name terminateReceiveThread
         * @brief terminate the receiver thread and set the connection state to false.
         */
        void terminateReceiveThread();
    }; /* end of class */
}

#include "close.cpp"
#include "constructor.cpp"
#include "destructor.cpp"
#include "equalityOperator.cpp"
#include "getIp.cpp"
#include "isConnected.cpp"
#include "print.cpp"
#include "publishEvent.cpp"
#include "receiveTask.cpp"
#include "send.cpp"
#include "setEventsHandler.cpp"
#include "setIp.cpp"
#include "startListen.cpp"
#include "terminateReceiveThread.cpp"

#endif /* networking_TcpCommon_Client_H */