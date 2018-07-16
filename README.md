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
   P = app    | box    | cat
     | copy   | drop   | swap
     | nocopy | nodrop | noswap
     | eq     | neq    | tag
     | [P]    | P P

      [A] app    = A
      [A] box    = [[A]]
  [A] [B] cat    = [A B]
      [A] copy   = [A] [A]    if A does not contain nocopy
      [A] drop   =            if A does not contain nodrop
  [A] [B] swap   = [B] [A]    if A and B do not contain noswap
      [A] nocopy = nocopy [A]
      [A] nodrop = nodrop [A]
      [A] noswap = noswap [A]
  [A] [B] eq     = [A] [B]    if A == B
  [A] [B] neq    = [A] [B]    if A != B
  [A] [B] tag    = [A] [B]    invoke the runtime with A and B
```

## Analysis
### Tags

## Synthesis

## Examples

## FAQ
