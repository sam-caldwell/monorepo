#
# configureGolang.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This file is the single point of including functions
# for use in the CMakeLists.txt file.
#
option(LANGUAGE "Language to build (e.g., cpp, go)" cpp)
option(BUILD_GROUP "Build group (cmd/projects)" cmd)

include(cmake/logging/main.cmake)
include(cmake/function/main.cmake)

include(cmake/os/apple.cmake)
include(cmake/os/unix.cmake)
include(cmake/os/windows.cmake)


