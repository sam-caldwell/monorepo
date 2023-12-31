Monorepo
===========

# What is This Repo?

This is the monorepo for all of my projects.  Why a monorepo?  Good question!  I found myself wanting to work on projects in my free time only to spend so much time setting up common stuff for each project (and thus wasting time).  So, I decided to pursue a minimum economy of scale...
1. This is a monorepo that pursues life from a `DevOps` perspective (no, not the job title but from the perspective of automating common tasks through tooling to eliminate the need for that grumpy legacy "ops person" you all know me to have been...freeing me to write code, build things and be happy me.)
2. This monorepo means I spend more time building various projects while still creating code others can use (or that I can share with my employers), allowing a lot of people to benefit from the stuff I do in my free time.
# What is This Repo ***not***?
1. This repo is a mirror of a local repo I keep on my internal network.  Not everything in this repo gets pushed to [Github](https://github.com/sam-caldwell).  What you see publicly is a filtered subset of the local monorepo (git server).  Why?  Well...not everything I build needs to be shared.
2. This repo should NEVER contain any of my proof-of-concept, bug bounty or other similar work.  Sorry, script kiddie, but I just do not want to be responsible for your irresponsibility.
3. Tools published to this repo are tools I believe are reasonably safe for public consumption.  Some tools may have dual purpose, but again, I'm not looking to help someone screw up their lives (or the lives of others).

# How is This Repo Organized?
1. I try to follow these rules...but there are no absolutes.
2. The general pattern is--`./<language>/<projectName>/...`
3. Binaries built through the tooling aren't published to the repo but would be found in `build/`
