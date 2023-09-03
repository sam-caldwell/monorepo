# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This module is part of github.com/sam-caldwell/monorepo/ and is released under
# the MIT LICENSE (See https://github.com/sam-caldwell/monorepo/LICENSE.txt.  Nothing
# in this file should be interpreted as interfering or conflicting with the license or
# rights conveyed by the wider CMAKE project under its respective license.
#
# To edit this file, edit monorepo/cmake/CustomModules/CMakeGOInformation.cmake.
# The version in the cmake Modules/ directory will be overwritten when cmake is run.
#

# This file sets the basic flags for Golang in cmake and loads available platform
# information where possible.
if(NOT CMAKE_Go_COMPILE_OBJECT)
  set(CMAKE_Go_COMPILE_OBJECT "go tool compile -l -N -o <OBJECT> <SOURCE> ")
endif()

if(NOT CMAKE_Go_LINK_EXECUTABLE)
  set(CMAKE_Go_LINK_EXECUTABLE "go tool link -o <TARGET> <OBJECTS>  ")
endif()