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
copyModulesLocal()
#check prerequisites
#install prerequisites

project(MONOREPO VERSION ${MONOREPO_VERSION} LANGUAGES CXX ASM C GO)

# configure languages
configureCpp()
configureGolang()

get_property(enabled_languages GLOBAL PROPERTY ENABLED_LANGUAGES)
ok("Enabled Languages: [${enabled_languages}]")

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
cppBuildTargets()

#add_executable(monorepo-version cmd/tools/monorepo-version/main.cpp)
#set_target_properties(monorepo-version PROPERTIES RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BIN_DIR} LINKER_LANGUAGE CXX)

#add_executable(monorepo-version cmd/tools/monorepo-version/main.cpp)
#set_target_properties(monorepo-version PROPERTIES RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BUILD_ROOT} LINKER_LANGUAGE CXX)