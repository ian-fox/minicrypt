# MiniCrypt

[![CircleCI](https://circleci.com/gh/ian-fox/minicrypt.svg?style=svg)](https://circleci.com/gh/ian-fox/minicrypt)

MiniCrypt is a personal interest project which aims to generate cryptic crosswords in a bite-sized form factor, like the [miniature crosswords](https://www.nytimes.com/crosswords/game/mini) of the New York Times.

I will be blogging about its development at [ianfox.dev](https://ianfox.dev).

It is currently in major version 0 and thus may have breaking changes introduced at any time.

## Installation

```go get -u github.com/ian-fox/minicrypt```

## Usage

```minicrypt path/to/wordlist.txt```

## Structure

Subdirectory | Purpose
-------------|--------
`cmd` | Command line interface
`grid` | Grid generation
`wordutill` | Utilities such as prefix trees and anagrams

## Links

[KanBan board](https://trello.com/b/HUiJylhk/minicrypt)
