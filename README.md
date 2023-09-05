Monorepo
===========

## Description

This is a monorepo for my various projects (golang, c++, python, etc.).  The goal of this repo is to reduce the 
amount of time I spend starting new projects or maintaining older ones.

## Structure
This repo follows a simple pattern:
```
  ./<language>/<cmd|projects>/<projectName>/...
```
Binary programs (executables) produced by this repo are found at--

    `<language>/cmd/<projectName>/<programName>`

Reusable source code is found at--

    `<language>/projects/<projectName>/...`

## Current Work
This is a work in progress at this time.

Currently work is being done to integrate multiple supported languages with a single set of build tools.

### Obstacles
* Getting Golang to work in this ecosystem is a problem. While golang is a great language.  But its package 
  management and community is far too opinionated to play nice with other tools.
