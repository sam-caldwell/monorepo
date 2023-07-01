Windows Management Instrumentation
==================================

## Description

This document at present details what would be needed to implement a new WMI/OLE interface for windows
using golang.

## Requirements

1. Package wmi provides a WQL interface to Windows WMI.
   > Note: It interfaces with WMI on the local machine, therefore it only runs on Windows.
   > It could interact with a remote Windows machine but that adds significant other issues.
2. This would replace `go-ole` (`projects/ole`) and `projects/wmi` with a pure Golang Equivalent.

## Effort:

1. Creating a pure Golang equivalent of the Go-OLE and WMI projects would require designing modules that are
   capable of interfacing directly with the underlying Windows API in a clean, efficient manner without the
   need for CGO or other C++ code. This could be a complex task as it requires comprehensive knowledge of both
   the Windows API and the Golang language.

## Components:

Here are the main components that would need to be developed:

1. ***OLE Automation/COM Interface:***
   > This will replace the Go-OLE functionality. It will allow Golang to interface with the Component Object
   > Model (COM) in Windows, which is crucial for many types of automation and interaction with the operating
   > system.
2. ***WMI Query Engine:***
   > This will replace the WMI functionality. It would need to send WMI queries and interpret the responses,
   > providing a Golang-native way to access system information and settings.
3. ***Error Handling and Logging:***
   > This component is crucial for diagnosing issues and ensuring smooth operation.
4. ***Data Marshalling/Unmarshalling***
   > This component will ensure that data transferred between Golang and the Windows API is
   > correctly interpreted and safely handled.
5. ***Testing and Validation***
   > This will ensure that the software is working as expected and that it provides the same functionality as 
   > the original projects.
```text
+---------------------------+
|        Application        |
|  +---------------------+  |
|  |   OLE Automation    |  |
|  |   /COM Interface    |  |
|  +---------------------+  |
|  +---------------------+  |
|  |  WMI Query Engine   |  |
|  +---------------------+  |
|  +---------------------+  |
|  |   Error Handling &  |  |
|  |        Logging      |  |
|  +---------------------+  |
|  +---------------------+  |
|  |   Data Marshalling/ |  |
|  |     Unmarshalling   |  |
|  +---------------------+  |
|  +---------------------+  |
|  | Testing & Validation|  |
|  +---------------------+  |
+---------------------------+
```
