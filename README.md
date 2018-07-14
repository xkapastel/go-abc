# ABC
ABC is a virtual machine for functional programs. ABC is unique in
integrating programming language theory with state of the art machine
learning research.

## Contents
- [Installation](#installation)
- [Hypermedia](#hypermedia)
- [Rewriting](#rewriting)
- [Editing](#editing)
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

## Editing

## Analysis

## Synthesis

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
### Why is ABC unityped? Don't types help with program search?
1. There are many type systems and we want them all.

Typed Clojure and Typed Racket demonstrate that with a sufficiently
introspective runtime, the distinction between unityped and mulityped
languages is not meaningful.

The only way to embed every possible type system is to make the
foundation unityped. You can always descend in to a subset of the
space of all possible programs, and reject anything that doesn't obey
your domain's type theory.

2. Probability theory may complement type theory.

Machine learning presents a unique opportunity for static analysis of
programs that are intractable with traditional methods. Type
information is certainly important for pruning the search space of
programs, but it seems premature to declare programs that cannot be
typed with brittle human logical analysis to be illegal right off the
bat. A unityped core with means of restricting programs to subsets of
the space of all possible programs seems like the right approach.
