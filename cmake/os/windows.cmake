#
# configure.windows.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This file configures C++ compilers, etc.
#

# For Windows
if (MSVC)
    # Use the latest C++ standard.
    add_definitions("/std:c++latest")
    # We need to report the correct value for the __cplusplus macro.
    add_definitions("/Zc:__cplusplus")
endif()

