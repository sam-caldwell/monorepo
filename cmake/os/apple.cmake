#
# apple.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This file configures C++ compilers, etc.
#
if(APPLE)
    set(CMAKE_CXX_COMPILER /Library/Developer/CommandLineTools/usr/bin/c++)
endif()