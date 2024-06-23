/**
 * @name networking/TcpServer/src/main.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "../../TcpCommon/src/PipeResponse/main.h"

/**
 * @class TcpServer
 * @brief This is the TcpServer for listening on a given IP address and port.
 */
class TcpServer {
public:
    TcpServer(){
        _subscribers.reserve(10);
        _clients.reserve(10);
        _stopRemoveClientsTask = false;
    }

    ~TcpServer(){
        close();
    }
    /*
     * Bind port and start listening
     * Return PipeResponse
     */
    PipeResponse start(int port, int maxNumOfClients = 5, bool removeDeadClientsAutomatically = true){
        if (removeDeadClientsAutomatically) {
            _clientsRemoverThread = new std::thread(&TcpServer::removeDeadClients, this);
        }
        try {
            initializeSocket();
            bindAddress(port);
            listenToClients(maxNumOfClients);
        } catch (const std::runtime_error &error) {
            return PipeResponse::failure(error.what());
        }
        return PipeResponse::success();
    }

    void initializeSocket(){
        SocketFileDescriptor.set(socket(AF_INET, SOCK_STREAM, 0));
        const bool socketFailed = (SocketFileDescriptor.get() == -1);
        if (socketFailed) {
            throw std::runtime_error(strerror(errno));
        }

        // set socket for reuse (otherwise might have to wait 4 minutes every time socket is closed)
        const int option = 1;
        setsockopt(SocketFileDescriptor.get(), SOL_SOCKET, SO_REUSEADDR, &option, sizeof(option));
    }

    void bindAddress(int port){
        memset(&_serverAddress, 0, sizeof(_serverAddress));
        _serverAddress.sin_family = AF_INET;
        _serverAddress.sin_addr.s_addr = htonl(INADDR_ANY);
        _serverAddress.sin_port = htons(port);

        const int bindResult = bind(SocketFileDescriptor.get(), (struct sockaddr *)&_serverAddress, sizeof(_serverAddress));
        const bool bindFailed = (bindResult == -1);
        if (bindFailed) {
            throw std::runtime_error(strerror(errno));
        }
    }

    void listenToClients(int maxNumOfClients){
        const int clientsQueueSize = maxNumOfClients;
        const bool listenFailed = (listen(SocketFileDescriptor.get(), clientsQueueSize) == -1);
        if (listenFailed) {
            throw std::runtime_error(strerror(errno));
        }
    }
    /*
     * Accept and handle new client socket. To handle multiple clients, user must
     * call this function in a loop to enable the acceptance of more than one.
     * If timeout argument equal 0, this function is executed in blocking mode.
     * If timeout argument is > 0 then this function is executed in non-blocking
     * mode (async) and will quit after timeout seconds if no client tried to connect.
     * Return accepted client IP, or throw error if failed
     */
    std::string acceptClient(uint timeout) {
        const PipeResponse waitingForClient = waitForClient(timeout);
        if (!waitingForClient.isSuccessful()) {
            throw std::runtime_error(waitingForClient.message());
        }

        socklen_t socketSize  = sizeof(_clientAddress);
        const int fileDescriptor = accept(SocketFileDescriptor.get(), (struct sockaddr*)&_clientAddress, &socketSize);

        const bool acceptFailed = (fileDescriptor == -1);
        if (acceptFailed) {
            throw std::runtime_error(strerror(errno));
        }

        auto newClient = new Client(fileDescriptor);
        newClient->setIp(inet_ntoa(_clientAddress.sin_addr));
        using namespace std::placeholders;
        newClient->setEventsHandler(std::bind(&TcpServer::clientEventHandler, this, _1, _2, _3));
        newClient->startListen();

        std::lock_guard<std::mutex> lock(_clientsMtx);
        _clients.push_back(newClient);

        return newClient->getIp();
    }

    void subscribe(const ServerObserver & observer){
        std::lock_guard<std::mutex> lock(_subscribersMtx);
        _subscribers.push_back(observer);
    }
    /*
     * Send message to all connected clients.
     * Return true if message was sent successfully to all clients
     */
    PipeResponse sendToAllClients(const char * msg, size_t size){
        std::lock_guard<std::mutex> lock(_clientsMtx);

        for (const Client *client : _clients) {
            PipeResponse sendingResult = sendToClient(*client, msg, size);
            if (!sendingResult.isSuccessful()) {
                return sendingResult;
            }
        }

        return PipeResponse::success();
    }

    /*
     * Send message to specific client (determined by client IP address).
     * Return true if message was sent successfully
     */
    PipeResponse sendToClient(const std::string &clientIP, const char *msg, size_t size){
        std::lock_guard<std::mutex> lock(_clientsMtx);

        const auto clientIter = std::find_if(_clients.begin(), _clients.end(),
                                             [&clientIP](Client *client) { return client->getIp() == clientIP; });

        if (clientIter == _clients.end()) {
            return PipeResponse::failure("client not found");
        }

        const Client &client = *(*clientIter);
        return sendToClient(client, msg, size);
    }

    PipeResponse close(){
        terminateDeadClientsRemover();
        { // close clients
            std::lock_guard<std::mutex> lock(_clientsMtx);

            for (Client * client : _clients) {
                try {
                    client->close();
                } catch (const std::runtime_error& error) {
                    return PipeResponse::failure(error.what());
                }
            }
            _clients.clear();
        }

        { // close server
            const int closeServerResult = ::close(SocketFileDescriptor.get());
            const bool closeServerFailed = (closeServerResult == -1);
            if (closeServerFailed) {
                return PipeResponse::failure(strerror(errno));
            }
        }

        return PipeResponse::success();
    }

    void printClients(){
        std::lock_guard<std::mutex> lock(_clientsMtx);
        if (_clients.empty()) {
            std::cout << "no connected clients\n";
        }
        for (const Client *client : _clients) {
            client->print();
        }
    }

private:

    FileDescriptor SocketFileDescriptor;

    struct sockaddr_in _serverAddress;

    struct sockaddr_in _clientAddress;

    fd_set _fds;

    std::vector<Client*> _clients;

    std::vector<ServerObserver> _subscribers;

    std::mutex _subscribersMtx;

    std::mutex _clientsMtx;

    std::thread * _clientsRemoverThread = nullptr;

    std::atomic<bool> _stopRemoveClientsTask;
    /*
     * Publish IncomingPacketHandler client message to observer.
     * Observers get only messages that originated
     * from clients with IP address identical to
     * the specific observer requested IP
     */
    void publishClientMsg(const Client & client, const char * msg, size_t msgSize){
        std::lock_guard<std::mutex> lock(_subscribersMtx);

        for (const ServerObserver& subscriber : _subscribers) {
            if (subscriber.wantedIP == client.getIp() || subscriber.wantedIP.empty()) {
                if (subscriber.IncomingPacketHandler) {
                    subscriber.IncomingPacketHandler(client.getIp(), msg, msgSize);
                }
            }
        }
    }
    /*
     * Publish client disconnection to observer.
     * Observers get only notify about clients
     * with IP address identical to the specific
     * observer requested IP
     */
    void publishClientDisconnected(const std::string&, const std::string&) {
        std::lock_guard<std::mutex> lock(_subscribersMtx);

        for (const ServerObserver& subscriber : _subscribers) {
            if (subscriber.wantedIP == clientIP) {
                if (subscriber.DisconnectionHandler) {
                    subscriber.DisconnectionHandler(clientIP, clientMsg);
                }
            }
        }
    }

    PipeResponse waitForClient(uint32_t timeout) {
        if (timeout > 0) {
            const FdWait::Result waitResult = FdWait::waitFor(SocketFileDescriptor, timeout);
            const bool noIncomingClient = (!FD_ISSET(SocketFileDescriptor.get(), &_fds));

            if (waitResult == FdWait::Result::FAILURE) {
                return PipeResponse::failure(strerror(errno));
            } else if (waitResult == FdWait::Result::TIMEOUT) {
                return PipeResponse::failure("Timeout waiting for client");
            } else if (noIncomingClient) {
                return PipeResponse::failure("File descriptor is not set");
            }
        }

        return PipeResponse::success();
    }
    /**
     * Handle different client events. Subscriber callbacks should be short and fast, and must not
     * call other server functions to avoid deadlock
     */
    void clientEventHandler(const Client&, ClientEvent, const std::string &msg){
        switch (event) {
            case ClientEvent::DISCONNECTED: {
                publishClientDisconnected(client.getIp(), msg);
                break;
            }
            case ClientEvent::INCOMING_MSG: {
                publishClientMsg(client, msg.c_str(), msg.size());
                break;
            }
        }
    }

    /**
     * Remove dead clients (disconnected) from clients vector periodically
     */
    void removeDeadClients(){
        std::vector<Client*>::const_iterator clientToRemove;
        while (!_stopRemoveClientsTask) {
            {
                std::lock_guard<std::mutex> lock(_clientsMtx);
                do {
                    clientToRemove = std::find_if(_clients.begin(), _clients.end(),
                                                  [](Client *client) { return !client->isConnected(); });

                    if (clientToRemove != _clients.end()) {
                        (*clientToRemove)->close();
                        delete *clientToRemove;
                        _clients.erase(clientToRemove);
                    }
                } while (clientToRemove != _clients.end());
            }

            sleep(2);
        }
    }

    void terminateDeadClientsRemover(){
        if (_clientsRemoverThread) {
            _stopRemoveClientsTask = true;
            _clientsRemoverThread->join();
            delete _clientsRemoverThread;
            _clientsRemoverThread = nullptr;
        }
    }

    /*
     * Send message to specific client (determined by client IP address).
     * Return true if message was sent successfully
     */
    static PipeResponse sendToClient(const Client & client, const char * msg, size_t size){
        try{
            client.send(msg, size);
        } catch (const std::runtime_error &error) {
            return PipeResponse::failure(error.what());
        }

        return PipeResponse::success();
    }

};