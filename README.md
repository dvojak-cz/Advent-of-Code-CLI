# Advent of Code CLI (aof)

[![Go Report Card](https://goreportcard.com/badge/github.com/dvojak-cz/Advent-of-Code-CLI)](https://goreportcard.com/report/github.com/dvojak-cz/Advent-of-Code-CLI)

Aoc is a simple CLI client for popular [AdventOfCode](https://adventofcode.com/). This client is ment to be used as a client for interacting with the website. This CLI allows to download puzzles as a `.md` file, allows downloading input data for a puzzle and can be used for submitting solutions of (see `dev` branch).

```txt
Advent of Code CLI is a simple CLI for Advent of Code
    It allows you to:
        to download puzzles
        to download puzzle input
        to submit puzzle answers and see the results

Usage:
  aoc [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         Get input for a day
  help        Help about any command
  login       Login to Advent of Code
  logout      Logout from Advent of Code

Flags:
  -h, --help   help for aoc

Use "aoc [command] --help" for more information about a command.
```

### Problems
- Since AOC does not provide any API and stable endpoints, and the author of *AdventOfCode* does not guarantee that any of the endpoints on the site are stable, there might be possibilities when CLI does not work correctly. In that case please create an issue.