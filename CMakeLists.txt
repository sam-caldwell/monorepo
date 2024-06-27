cmake_minimum_required(VERSION 3.29)
project(monorepo)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

set(DEBUG TRUE)
set(CMAKE_BUILD_DIR "${CMAKE_BINARY_DIR}/build")
set(CPP_ENABLED TRUE)
#set(GO_ENABLED TRUE)
enable_testing()

#
# --------------------------------
# Setup the reusable CMAKE tooling
# --------------------------------
#

include("CMake/colors.cmake")
include("CMake/logger.cmake")
include("CMake/require.cmake")
include("CMake/include_if_exists.cmake")
include("CMake/create_build_dir.cmake")
include("CMake/create_documentation_header.cmake")
include("CMake/create_documentation_line.cmake")
include("CMake/create_documentation_footer.cmake")
include("CMake/add_cpp_projects.cmake")
include("CMake/add_go_projects.cmake")
#
# -----------------------
# Create our build target
# -----------------------
#

create_build_dir()


#
# ---------------------------------------
# include language-specific project files
# ---------------------------------------
#

if (DEFINED CPP_ENABLED)
    add_cpp_projects(${CMAKE_SOURCE_DIR}/cpp)
endif ()
if (DEFINED GOLANG_ENABLED)
#    add_go_projects(${CMAKE_SOURCE_DIR}/go)
endif ()
