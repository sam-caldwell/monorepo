#
# cpp/configureCpp.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# Configure C++
#
function(configureCpp)
    debug("configureCpp()")
    set(CMAKE_CXX_STANDARD 20)
    set(CMAKE_CXX_STANDARD_REQUIRED True)
    set(CMAKE_CXX_EXTENSIONS OFF)
    debug("   [CMAKE_CXX_STANDARD          : ${CMAKE_CXX_STANDARD}]")
    debug("   [CMAKE_CXX_STANDARD_REQUIRED : ${CMAKE_CXX_STANDARD_REQUIRED}]")
    debug("   [CMAKE_CXX_EXTENSIONS        : ${CMAKE_CXX_EXTENSIONS}]")
    enable_language("CXX")
    ok("configureCpp()")
endfunction()