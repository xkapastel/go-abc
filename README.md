# ABC
ABC is a virtual machine for functional programs, integrating modern
research in programming languages and artificial intelligence.

## Contents
- [Installation](#installation)
- [Hypermedia](#hypermedia)
- [Rewriting](#rewriting)
- [Analysis](#analysis)
- [Synthesis](#synthesis)
- [Examples](#examples)
- [FAQ](#faq)

## Installation
`go install -u -v github.com/xkapastel/abc/cmd/abc` will install the
`abc` command.

## Hypermedia
ABC bytecode is hyperlinked, based on a content-addressing scheme.

## Rewriting
```
    P = app  | box  | cat
      | copy | drop | swap
      | [P]  | P P

         [A] app = A
         [A] box = [[A]]
     [A] [B] cat = [A B]
        [A] copy = [A] [A]
        [A] drop =
    [A] [B] swap = [B] [A]
```

## Analysis

## Synthesis

## Examples

## FAQ
