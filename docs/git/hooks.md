Git Hooks
=========

## Description
Document the use of local git hooks in this project.

## What are Git hooks
Githooks are automation that can execute on your dev environment
before/after git operations.  For example, this project provides
git hooks that perform linter and SAST checking of all code before
a commit is accepted.

## How Do we Install Git Hooks?
1. A git hook is part of the local git repository.
2. To install the git hooks, run `make hooks` on your local git repo.