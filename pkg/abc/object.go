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

// ABC is a universal combinator calculus, with six primitives:
//
//         [A] a = A
//         [A] b = [[A]]
//     [A] [B] c = [A B]
//         [A] d = [A] [A]
//         [A] e =
//     [A] [B] f = [B] [A]
//
// Additionally, code is hyperlinked with a content-based addressing
// scheme. An object may refer to another object by its SHA-256 hash.
// This allows compression and an easy opportunity for acceleration.
type Object interface {
	String() string
	// step attempts to perform a rewrite, returning whether or not
	// any work was actually done.
	step(*rewrite) bool
	// eq predicates structurally equivalent objects.
	eq(Object) bool
}

// Equals predicates structurally equivalent objects.
func Equals(fst, snd Object) bool {
	return fst.eq(snd)
}
