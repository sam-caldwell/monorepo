cmake_minimum_required(VERSION 3.29)

set(CMAKE_LOGGER True)

if (NOT DEFINED CMAKE_COLORS)
    message(FATAL_ERROR "COLORS must be defined")
endif()

# Print a DEBUG message (if DEBUG is set globally)
function(logDebug msg)
    if (DEFINED DEBUG)
        Blue("DEBUG:   ${msg}")
    endif ()
endfunction()

# Print an OK message
function(logOK msg)
    Green("OK:    ${msg}")
endfunction()

# Print a warning message
function(logWARN msg)
    Yellow("WARN:  ${msg}")
endfunction()

# Print an error message
function(logFAIL msg)
    Red("FAIL:   ${msg}")
endfunction()

# Print a message and terminate
function(logFATAL msg)
    Red("FAIL:   ${msg}")
    message(FATAL_ERROR)
endfunction()

function(testLogger)
    logDebug("test (should not print)")
    set(DEBUG true)
    logDebug("test")
    logOk("test")
    logWARN("test")
    logFAIL("test")
    logFATAL("test")
endfunction()

if (DEFINED DEBUG)
    Blue("DEBUG:  Enabled")
endif ()