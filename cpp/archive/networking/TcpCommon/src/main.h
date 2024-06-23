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

#include "../../../../old/application/Configuration/src/constants.h"
#include "clientEvent.h"
#include "messages.h"
#include "ResultsEnum.h"
#include "waitFor.h"
#include "PipeResponse/main.h"
#include "ClientObserver.h"
#include "FileDescriptor.h"
#include "ServerObserver.h"
#include "PipeResponse/main.h"
