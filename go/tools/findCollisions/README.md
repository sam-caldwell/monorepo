find_collisions (CRSCE 1024-byte block)
========================================

## Objectives
To test the feasibility of SHA-1 in the narrow use-case of CRSCE.

## Discussion
This program is an experiment to test the possible use of SHA-1 in the CRSCE project.
By design as of August 2023, CRSCE is implemented using SHA-256.  But SHA-256 is a
significant performance obstacle for CRSCE.  We instead consider using SHA-1 which 
would perform faster and provide the same row and column level integrity protections.
However, this approach will only work if no collisions exist in this specific use case.

The use-case for SHA-1 in CRSCE, assumes at most a 1MB (1024x1024 byte) CSM message block.
This means that we must confirm that there exist no collisions when we have a 1MB message
block.  Specifically where our hashes are a hash of row or column, we only need to know
that there are no collisions for [1024]byte using SHA-1
