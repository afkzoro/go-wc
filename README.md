# go-wc

This is a Go implementation of the Unix `wc` command, which counts lines, words, and characters in a text file. I took on this project because of [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc/), but it was simply built for fun and learning purposes.

## Features

- Count lines (-l)
- Count bytes (-c)
- Count words (-w)
- Count characters (-m)

## Installation

```bash
$ go get github.com/afkzoro/go-wc
```

## Usage

```bash
# Count all (lines, words, bytes)
$ ./go-wc filename.txt

# Count only lines
$ ./go-wc -l filename.txt

# Count only bytes
$ ./go-wc -c filename.txt

# Count only characters
$ ./go-wc -m filename.txt

# Use with pipes
$ ./go-wc cat filename.txt | wc -l
```

## Project structure

- `cmd/wc`: Main application entry point
- `internal/counter`: Counting logic
- `internal/reader`: Input handling
- `internal/printer`: Output formatting

## Development

I wrote a Makefile to ease things up. You can easily automate stuff:

```bash
# Building
$ make build

# Testing
$ make test

# Cleaning
$ make clean
```
