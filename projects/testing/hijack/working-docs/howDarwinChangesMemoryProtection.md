How Darwin Changes Memory Protections
=====================================
---------------------------------------------------------
"How I Spent a Weekend Massaging my Desk With My Head"

(c) 2023 Sam Caldwell.  See LICENSE.txt

---------------------------------------------------------

At the lowest level in Darwin, calls to ``Mprotect`` via `syscall` modify the memory protection flags of memory pages. 
Here's a step-by-step overview of how `Mprotect` works at the lowest level in Darwin:

<pre>
  +-------------------+
  |  Calling Program  |
  +-------------------+
          |
          | (1) The `Mprotect` function is called with the memory address of the memory region to 
          |     modify, the length of the region, and the desired memory protection flags.
          V
  +-------------------+
  |   Mprotect        |
  +-------------------+
          |
          | (2) The `Mprotect` function uses the syscall package to make a system call to the `Mprotect` 
          |     function in the underlying operating system.
          V
  +-------------------+
  |   syscall pkg     |
  +-------------------+
          |
          | (3) The syscall package uses the libc_Mprotect_trampoline function, which is a trampoline 
          |     function that bridges the Go code to the dynamic symbol Mprotect in the libc library.
          V
  +-------------------+
  |   libc library    |
  +-------------------+
          |
          | (4) The `Mprotect` system call is invoked with the memory address, length, and protection 
          |     flags as arguments.
          V
  +-------------------+
  |   Kernel          |
  +-------------------+
          |
          | (6) The operating system's kernel receives the `Mprotect` system call and performs the requested 
          |     memory protection changes.
          V
  +-------------------+
  |   Kernel          |
  |   Memory          |
  |   Protection      |
  +-------------------+
          |
          | (7) The kernel checks the validity of the memory region specified by the memory address and length. 
          |     It verifies that the memory region is valid and accessible by the calling process.
          V
  +-------------------+
  |   Kernel          |
  |   Page tables     |
  +-------------------+
          |
          | (9) If the memory region is valid, the kernel modifies the protection flags of the memory pages within 
          |     the specified region according to the requested protection flags.
          | 
          | (10) The kernel updates the page tables and applies the new protection flags to the memory pages. 
          |      This might involve marking the pages as read-only, read-write, or executable based on the 
          |      specified flags.
          | 
          | (11) The syscall package in Go receives the result of the system call and returns it to the `Mprotect` 
          |      function.
          V
  +-------------------+
  |   Mprotect        |
  +-------------------+
          |
          | (12) If the memory protection modification is successful, the `Mprotect` system call returns without 
          |      errors.
          |
          | (13) The `Mprotect` function returns any error encountered during the system call, indicating whether 
          |      the memory protection modification was successful or not.
          |
          V
  +-------------------+
  |   Caller          |
  +-------------------+

</pre>
