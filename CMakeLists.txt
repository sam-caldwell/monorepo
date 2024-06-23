cmake_minimum_required(VERSION 3.29)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

set(DEBUG TRUE)
set(CMAKE_BUILD_DIR "${CMAKE_BINARY_DIR}/build")
set(CPP_ENABLED TRUE)
#set(GO_ENABLED TRUE)

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
include("CMake/cpp_add_projects.cmake")

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
    cpp_add_projects(${CMAKE_SOURCE_DIR}/cpp)
endif ()
if (DEFINED GOLANG_ENABLED)
#    include_if_exists(go)
endif ()
# Add more project types here
