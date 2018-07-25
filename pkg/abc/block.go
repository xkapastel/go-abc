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

import ()

// A block is executable code. ABC is a universal combinator
// calculus, with six primitives:
//
//         [A] app    = A
//         [A] box    = [[A]]
//     [A] [B] cat    = [A B]
//         [A] copy   = [A] [A]
//         [A] drop   =
//     [A] [B] swap   = [B] [A]
//
// Additionally, code is hyperlinked with a content-based addressing
// scheme. A block may refer to another block by its SHA-256 hash.
// This allows compression and an easy opportunity for acceleration.
type Block interface {
	// String returns a block's source code.
	String() string
	// step attempts to perform a rewrite, returning whether or not
	// any work was actually done.
	step(*reduce) bool
	// eq predicates structurally equivalent blocks.
	eq(Block) bool
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

// Equals predicates structurally equivalent blocks.
func Equals(fst, snd Block) bool {
	return fst.eq(snd)
}
