CRSCE/Collsion Test (SHA1)
==========================

## Objective
Test whether SHA1 could be used for CRSCE compression rather than SHA2/3 to improve performance.

## Discussion

While it is clear SHA1 is broken for cryptographic purposes in the wake of SHAttered years ago, the
SHAttered attack was about three times the size of the optimal 1MB block size of the current SHA256-based
CRSCE algorithm.  If it can be proven that SHA1 is relatively collision proof at a 1MB message block size, then
SHA1 could still be mathematically useful for CRSCE, improving performance of both compression and decompression 
significantly.

Note: Even assuming collisions are possible, the probability of at least two collisions occurring to allow CSM to 
be altered in two elements (changing at least two row,column positions) is so low that SHA1 should still be useful
in CRSCE.