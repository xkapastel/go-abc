# ABC
ABC is a virtual machine for functional programs. ABC is unique in
integrating programming language theory with state of the art machine
learning research.

## Contents
- [Installation](#installation)
- [Hypermedia](#hypermedia)
- [Rewriting](#rewriting)
- [Synthesis](#synthesis)
- [Examples](#examples)
- [FAQ](#faq)

## Installation
`go install -u -v github.com/xkapastel/abc/cmd/abc` will install the `abc` command.

The `abc` command provides many tools for analyzing and synthesizing
ABC bytecode.

`abc print` displays a human-readable representation of the bytecode
stream provided on standard input.

`abc reduce` rewrites the bytecode stream on standard input until
either a normal form is reached or the effort quota is exhausted.

`abc box` wraps the bytecode stream provided on standard input in a
block.

`abc hash` computes the SHA-256 hash of the bytecode stream on
standard input, and writes it as a bytecode link to standard output.

## Hypermedia
ABC bytecode is hyperlinked, based on a content-addressing scheme.

## Rewriting

## Synthesis
A `Set` is an iteratively defined collection of `Block`s. The blocks
you get out are similar to the ones you put in. This can be used to
write a simple hill climber:

1. `Get()` N blocks from the set.
2. Evaluate the fitness of each block.
3. `Put()` the K best blocks in to the set.

The blocks you sample will gradually become more fit. This algorithm,
known as "priority queue training", was described in the Google Brain
paper "Neural Program Synthesis with Priority Queue Training".

A `Map` is an iteratively defined function between sets.

## Examples

```
$ cat box swap box swap cat > pair
$ cat app > unpair
$ cat swap drop swap app | abc box > left-0
$ cat box left-0 cat > left
$ cat pair swap drop unpair swap app | abc box > right-0
$ cat box right-0 cat > right
$ cat app > case
```

## FAQ
