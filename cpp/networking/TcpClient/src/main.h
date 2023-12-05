/**
 * @name networking/TcpClient/src/main.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Networking_TcpClient_H
#define Networking_TcpClient_H

#include "projects/networking/TcpCommon/src/main.h"
#include "projects/networking/TcpCommon/src/messages.h"

namespace networking {
/**
 * @class TcpClient
 * @brief This is the TcpClient class for sending packets to a given Ip address/port.
 */
    class TcpClient {
    public:
        /**
         * @name class constructor
         * @brief set initial state.
         */
        TcpClient();

        /**
         * @name class destructor
         * @brief close any open connection.
         */
        ~TcpClient();

        /**
         * @name connectTo
         * @brief
         *
         * @param address string
         * @param port int
         * @return PipeResponse
         */
        PipeResponse connectTo(const std::string &address, int port);

        /**
         * @name sendMsg
         * @brief
         *
         * @param msg
         * @param size
         * @return
         */
        PipeResponse sendMsg(const char *msg, size_t size);

        /**
         * @name subscribe
         * @brief
         *
         * @param observer
         */
        void subscribe(const ClientObserver &observer);

        /**
         * @name isConnected
         * @brief return the connected status of the socket.
         *
         * @return bool
         */
        inline bool isConnected() const;

        /**
         * @name close
         * @brief close the network connection
         *
         * @return PipeResponse
         */
        PipeResponse close();

    private:
        /**
         * @name SocketFileDescriptor
         * @brief This is the socket file descriptor because all things are file handles in *nix systems
         */
        FileDescriptor SocketFileDescriptor;
        /**
         *
         */
        std::atomic<bool> _isConnected;
        /**
         *
         */
        std::atomic<bool> _isClosed;
        /**
         *
         */
        struct sockaddr_in _server;
        /**
         *
         */
        std::vector <ClientObserver> _subscibers;
        /**
         *
         */
        std::thread *_receiveTask = nullptr;
        /**
         *
         */
        std::mutex _subscribersMtx;

        /**
         * @name initializeSocket
         * @brief initialize the network socket
         */
        void initializeSocket();

        /**
         * @name startReceivingMessages
         * @brief spawn the thread for receiving messages.
         */
        inline void startReceivingMessages();

        /**
         * @name setAddress
         * @brief set the Ip address and port
         *
         * @param address string
         * @param port int
         */
        void setAddress(const std::string &address, int port);

        /**
         * @name publishServerMsg
         * @brief publish an IncomingPacketHandler client message to the observer.
         * The observer will get only messages that originated from clients with an IP address
         * identical to the observer requested IP address.
         *
         * @param msg cstring
         * @param msgSize cstring length
         */
        void publishServerMsg(const char *msg, size_t msgSize);

        /**
         * @name publishServerDisconnected
         * @brief Publish a client disconnection to observer, and the observer will only get a notification
         *        where the client Ip address is identical to the specific observer requested Ip.
         *
         * @param result PipeResponse (by reference)
         */
        void publishServerDisconnected(const PipeResponse &result);

        /**
         * @name receiveTask
         * @brief Receive server packets, and notify user
         */
        void receiveTask();

        /**
         * @name terminateReceiveThread
         * @brief terminate the receiver threads         *
         */
        void terminateReceiveThread();
    };
}

#include "projects/networking/TcpClient/src/constructor.cpp"
#include "projects/networking/TcpClient/src/destructor.cpp"
#include "projects/networking/TcpClient/src/connectTo.cpp"
#include "projects/networking/TcpClient/src/sendMsg.cpp"
#include "projects/networking/TcpClient/src/subscribe.cpp"
#include "projects/networking/TcpClient/src/isConnected.cpp"
#include "projects/networking/TcpClient/src/close.cpp"

#include "projects/networking/TcpClient/src/initializeSocket.cpp"
#include "projects/networking/TcpClient/src/startReceivingMessages.cpp"
#include "projects/networking/TcpClient/src/setAddress.cpp"
#include "projects/networking/TcpClient/src/publishServerMsg.cpp"
#include "projects/networking/TcpClient/src/publishServerDisconnected.cpp"
#include "projects/networking/TcpClient/src/receiveTask.cpp"
#include "projects/networking/TcpClient/src/terminateReceiveThread.cpp"

#endif /*  Networking_TcpClient_H */