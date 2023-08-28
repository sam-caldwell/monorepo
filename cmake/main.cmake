#
# configureGolang.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This file is the single point of including functions
# for use in the CMakeLists.txt file.
#
include(cmake/logging/debug.cmake)
include(cmake/logging/debugList.cmake)
include(cmake/logging/debugLine.cmake)
include(cmake/logging/info.cmake)

if (EXISTS ${CMAKE_BINARY_DIR}/cmake/Modules)
    info("Update CMAKE Modules path to ${CMAKE_BINARY_DIR}/cmake/Modules")
    set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${CMAKE_BINARY_DIR}/cmake/Modules/")
endif ()

include(cmake/os/apple.cmake)
include(cmake/os/unix.cmake)
include(cmake/os/windows.cmake)

include(cmake/cpp/configureCpp.cmake)
include(cmake/go/configureGolang.cmake)

include(cmake/function/initBinaryDirectory.cmake)
include(cmake/function/initBuildDirectory.cmake)
include(cmake/function/showMonorepoVersion.cmake)
include(cmake/function/addProjectParameters.cmake)

include(cmake/function/createCppBuildTargets.cmake)
