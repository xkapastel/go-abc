/**
This file is a part of ABC.
Copyright (C) 2018 Matthew Blount

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but
WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public
License along with this program.  If not, see
<https://www.gnu.org/licenses/>.
**/

package abc

import (
	"io"
)

// A block is executable code. ABC is a universal combinator
// calculus, with seven primitives:
//
//         [A] app  = A
//         [A] box  = [[A]]
//     [A] [B] cat  = [A B]
//         [A] copy = [A] [A]
//         [A] drop =
//     [A] [B] swap = [B] [A]
//     [A] [A] eq   = [A] [A]
//
// Additionally, code is hyperlinked with a content-based addressing
// scheme. A block may refer to another block by its SHA-256 hash.
// This allows compression and an easy opportunity for acceleration.
type Block interface {
	// Box wraps a block in another pair of brackets.
	Box() Block
	// Cat composes many blocks left to right.
	Cat(...Block) Block
	// Eq predicates structurally equivalent blocks.
	Eq(Block) bool
	// Reduce rewrites a block until it either reaches a normal
	// form or the effort quota is exhausted.
	Reduce(int) Block
	// Encode writes a block's bytecode to a bytestream.
	Encode(io.ByteWriter) error
	// String is a human-readable notation for blocks.
	String() string
}

// Id is the identity and does nothing. It's represented by
// whitespace in the notation used throughout this documentation.
//     [A] = [A]
var Id Block

// App executes a block of code.
//     [A] app = A
var App Block

// Box quotes a block of code.
//     [A] box = [[A]]
var Box Block

// Cat composes two blocks of code.
//     [A] [B] cat = [A B]
var Cat Block

// Copy duplicates a block of code.
//     [A] copy = [A] [A]
var Copy Block

// Drop erases a block of code.
//     [A] drop =
var Drop Block

// Swap exchanges two blocks of code.
//     [A] [B] swap = [B] [A]
var Swap Block

func init() {
	Id = opId{}
	App = opApp{}
	Box = opBox{}
	Cat = opCat{}
	Copy = opCopy{}
	Drop = opDrop{}
	Swap = opSwap{}
}
