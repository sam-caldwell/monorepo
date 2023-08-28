#
# CMakeList.txt
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This file is the single point of including functions
# for use in the CMakeLists.txt file.
#
cmake_minimum_required(VERSION 3.27)
include(version.cmake)
include(cmake/main.cmake)

project(MONOREPO VERSION ${MONOREPO_VERSION} LANGUAGES CXX)

# Special Signals
set(CMAKE_DEBUG True)
configureCpp()
configureGolang()


# Global Parameters
enable_testing()
set(CTEST_OUTPUT_ON_FAILURE TRUE)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_BUILD_ROOT ${CMAKE_BINARY_DIR}/build)
set(CMAKE_BINARY_ROOT ${CMAKE_BINARY_DIR}/bin)

set(CMAKE_GOLANG_COMPILER ~/go/go1.21.0/bin/go)
set(CMAKE_GOLANG_GOROOT ~/go/go1.21.0)


# --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
# Start...
#
showMonorepoVersion()
initBuildDirectory()
initBinaryDirectory()
createCppBuildTargets()



#add_executable(monorepo-version cmd/tools/monorepo-version/main.cpp)
#set_target_properties(monorepo-version PROPERTIES RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BIN_DIR} LINKER_LANGUAGE CXX)

#add_executable(monorepo-version cmd/tools/monorepo-version/main.cpp)
#set_target_properties(monorepo-version PROPERTIES RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BUILD_ROOT} LINKER_LANGUAGE CXX)