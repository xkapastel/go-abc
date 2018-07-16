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

type opNeq struct{}

func (block opNeq) Box() Block { return &mkBox{block} }
func (block opNeq) Cat(xs ...Block) Block {
	rest := newCatN(xs...)
	return newCatN(block, rest)
}
func (block opNeq) Reduce(quota int) Block { return block }
func (block opNeq) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(CodeOpNeq)
}
func (block opNeq) String() string { return "nq" }
func (lhs opNeq) Eq(rhs Block) bool {
	_, ok := rhs.(opNeq)
	return ok
}
func (block opNeq) Copy() bool { return true }
func (block opNeq) Drop() bool { return true }
func (block opNeq) Swap() bool { return true }
func (block opNeq) step(ctx *reduce) bool {
	if ctx.arity() < 2 {
		ctx.clear(block)
		return false
	}
	lhs := ctx.peek(0)
	rhs := ctx.peek(1)
	if lhs.Eq(rhs) {
		ctx.clear(block)
		return false
	}
	return true
}
