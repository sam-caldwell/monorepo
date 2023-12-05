/**
 * @name LogServer/src/LogServer.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <fstream>
#include "projects/application/Configuration/src/constants.h"
#include "projects/application/Core/src/Application.h"

class LogServer : public Application {
public:

    LogServer() = delete;

    /**
     * @name class constructor
     * @brief Setup the application.
     * @param argc - argument count
     * @param argv - list of arguments
     */
    LogServer(int argc, char *argv[]) {
        ConfigStateMap map;
        map.add("CONFIG_FILE", ConfigFile, String, false, true);
        map.add("LISTEN_IP", ConfigLogServerIp, String, false, true);
        map.add("LISTEN_PORT", ConfigLogServerPort, String, false, true);
        map.add("LOG_TARGET", ConfigLogServerTarget, String, false, true);
        map.add("LOG_LEVEL", ConfigLogLevel, String, false, true);
        map.add("LOG_OUTPUT", ConfigLogOutput, String, false, true);
        map.add("LOG_TAGS", ConfigLogTags, String, false, true);
        map.add("LOG_DESTINATION", ConfigLogDestination, String, false, true);

        SignalHandler *signals = nullptr;

        Application::Application(&map, argc, argv, signals);

        TcpServer server(config, &handlerClass); //configure our TCP network server.
    }

    ~LogServer() { server.shutdown(); };

    bool start() {
        for (;;) server.start(); //Start the tcp server as a thread.
    };

    void logWriter();
};