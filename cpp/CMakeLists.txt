cmake_minimum_required(VERSION 3.29)

set(CPP_SOURCE_DIR ${CMAKE_SOURCE_DIR}/cpp)

include_if_exists(cpp/TesterFramework)
include_if_exists(cpp/common)
include_if_exists(cpp/services)
include_if_exists(cpp/tools)

logDebug("C++ projects are included")
