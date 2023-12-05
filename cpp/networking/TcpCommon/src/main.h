/**
 * @name networking/TcpCommon/src/main.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <arpa/inet.h>
#include <atomic>
#include <cstdio>
#include <cerrno>
#include <cstring>
#include <errno.h>
#include <functional>

#include <iostream>
#include <mutex>
#include <netdb.h>
#include <netinet/in.h>

#include <stdexcept>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <thread>
#include <unistd.h>
#include <vector>

#include "projects/networking/TcpCommon/src/constants.h"
#include "projects/networking/TcpCommon/src/clientEvent.h"
#include "projects/networking/TcpCommon/src/messages.h"
#include "projects/networking/TcpCommon/src/ResultsEnum.h"
#include "projects/networking/TcpCommon/src/waitFor.h"
#include "projects/networking/TcpCommon/src/PipeResponse/main.h"
#include "projects/networking/TcpCommon/src/ClientObserver.h"
#include "projects/networking/TcpCommon/src/FileDescriptor.h"
#include "projects/networking/TcpCommon/src/ServerObserver.h"
#include "projects/networking/TcpCommon/src/client/main.h"
