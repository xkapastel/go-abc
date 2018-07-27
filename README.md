# ABC
ABC is a virtual machine for functional programs, integrating modern
research in programming languages and artificial intelligence.

- [Introduction](#introduction)
- [Getting Started](#getting-started)
- [Functions](#functions)
- [Numbers](#numbers)
- [Rewriting](#rewriting)
- [Hypermedia](#hypermedia)
- [Modules](#modules)
- [Blobs](#blobs)
- [Annotations](#annotations)
- [Accelerators](#accelerators)

## Introduction

## Getting Started
`go install -u -v github.com/xkapastel/abc/cmd/abc` will install the
`abc` command.

## Functions
Functions are the basic building blocks of computation. ABC functions
are true functions, in the sense that they have no causal dependencies
besides their arguments.

## Numbers
ABC accepts numbers in IEEE 754 syntax.

## Rewriting
When you run an ABC program, the result is another program,
potentially simplified.

```
    [A] a = A
    [A] b = [[A]]
[A] [B] c = [A B]
    [A] d = [A] [A]
    [A] e =
[A] [B] f = [B] [A]
```

## Hypermedia
ABC programs are hyperlinked, based on a content-addressing scheme.

## Modules

## Blobs
Blobs are lists of bytes, allowing ABC programs to refer to foreign
objects by their content address.

## Annotations

## Accelerators
