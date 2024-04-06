EDR Tool
========

## Objective

* Design and implement basic EDR functionality.

## Assumptions
1. Zero Trust
    1. Neither client nor server trusts one another.
        1. The attacker is already on the system.
        2. Clients do not execute arbitrary queries (only predefined ThreatQL)
        3. Data provided by the client at any point is suspect. We can only trust
           data over time to eventually reveal the truth.
2. Ad hoc queries would require JIT compilation which would be slow on the client.
3. There are two types of data:
    1. continuously collected data.
    2. specifically requested data.
4. All queries are "select" or read queries. We do not need create/update/delete.
5. We should not be able to read a file's content in a query, but we should be able to...
    1. Get the file's name and attributes (including permissions).
    2. Get the file's hash.
    3. Get a hash of a range of bytes in the file's contents.

## Pipeline

```text
--------------------------------------------------------------------------------------
                                    Client      ||    Server
------------------------------------------------||------------------------------------
                                                ||
main()                                          ||      main()
   go ProcMon()   |                             ||
   go NetMon()    |  channel                  network
   go HwMon()     |>-------+-> go Emitter() ====||=======> go Collector()
   go FsMon()     |        ^                    ||              |
   go SysMon()    |        |                    ||              |
                  |        |                    ||              +---> real-time
                  |        |                    ||              |     threat detection  
                  |        |                    ||              V
                  |>-------+                    ||            (logs)               
   go Query()     | channel                     ||   
                  |<------< go GetCheckIn()=====||========> go QuerySender()
                  |          (polls server)     ||                ^    
                                                ||                +---> (logs)
                                                ||                V
                                                ||            (query queue)
                                                          
```

## Data Collector GoRoutines

### ProcMon

* Runs as a goroutine to collect process table information and report the same to the event stream.
* Collects the following process information:
    * PID
    * USER
    * PRIORITY
    * NICE
    * CPU
    * MEM
    * Command
    * State (running, sleeping, etc.)
    * Terminal / network connections

### NetMon

* Runs as a goroutine to collect network connections and listening sockets and report the same to an event stream
* Collects the following network information:
    * Interface enumeration
    * Network connections
        * source address/port
        * target address/port
        * protocol (tcp, udp, tcp6, udp6, icmp, etc.)
        * associated process (PID)
        * bytes in
        * bytes out
        * packets/datagrams in/out

### HwMon

* Runs as a goroutine to collect hardware state changes and report the same to the event stream.
* Collects the following hardware information:
    * Hardware device enumeration (lshw/lspci/lsusb/lscpu/lsmem/lspcmcia).
    * Firmware versions and hashes
        * Main firmware (e.g. BIOS/UEFI)
        * hard disk firmware
        * changes in /sys/firmware/dmi/tables/DMI, for example.

### FsMon

* Runs as a goroutine to collect file system / disk changes and activity and report the same to the event stream.
* Collects the following fs information:
    * disk inventory
        * disk device identity
        * disk partition tables / geometry
        * mount points and mount tables (/etc/fstab and actual mounts) and mount attributes
    * File system activity
        * file properties (stat $file)
        * file hashes
        * file I/O events

### SysMon

* Runs as a goroutine to collect operating system state and changes
    * kernel version and hashes
    * driver versions and changes
    * software inventory (packages, names and versions)
    * environment variables
    * cron tabs.
    * firewall changes
    *
* 