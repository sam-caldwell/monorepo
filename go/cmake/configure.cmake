#
# go/configureGolang.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
#set(CMAKE_GO_COMPILER "${CMAKE_BINARY_DIR}/go/go1.21.0/bin/go")
#set(CMAKE_GO_GOROOT "${CMAKE_BINARY_DIR}/go/go1.21.0")

function(configureGolang)
    debug("configureGolang()")

    set(ENV{CGO_ENABLED} 0)
    set(ENV{CGO_LDFLAGS} "")
    set(ENV{CGO_CFLAGS} "")

    enable_language("GO")

    ok("configureGolang()")
endfunction()