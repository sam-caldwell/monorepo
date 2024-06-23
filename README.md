Monorepo
===========

# What is This Repo?

This is the monorepo for all of my projects.  Why a monorepo?  Good question!  I found myself wanting to work on 
projects in my free time only to spend so much time setting up common stuff for each project (and thus wasting time).
So, I decided to pursue a minimum economy of scale...  This monorepo allows me to provide common tooling and common 
code to all of my projects with a single versioning system.

No, this is not the perfect monorepo.  But it's good for testing new ideas, trying new things and experimenting in
a rapid-development way when you've got only two days a week (Saturday, Sunday) to get things done.

# Is this everything?
No.  Not everything gets published to GitHub.  I run a local git server and publish code to GitHub from that server, 
based on a filter program.  That filter program is intended to prevent some projects from being published to GitHub
(such as exploit proofs-of-concept), etc.

# License?
This is all covered in [License.txt](LICENSE.txt)


# Builds and Tests
To build and test...
```shell
 cmake . && cmake --build . && make test
```
Note: There are some cool reusable 'cmake' functions in [CMake/](./CMake) folks might enjoy.

# A Note on C++ code
For the C++ purists out there, avoid the `cpp` directory.  I've been experimenting with ways to make C++ safer
and doing other things that will make you lose your mind.  Different patterns, different concepts...trying to adapt
what I love about golang to c++.  You've been warned.