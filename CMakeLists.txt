#
# CMakeList.txt
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This file is the single point of including functions
# for use in the CMakeLists.txt file.
#
cmake_minimum_required(VERSION 3.27.4)
include(version.cmake)
include(cmake/main.cmake)
set(CMAKE_DEBUG True)
#check prerequisites
#install prerequisites

project(MONOREPO VERSION ${MONOREPO_VERSION} LANGUAGES CXX)

# configure languages
configureCpp()
configureGolang()


# Global Parameters
enable_testing()
set(CTEST_OUTPUT_ON_FAILURE TRUE)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)



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