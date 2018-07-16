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

type opDrop struct{}

func (block opDrop) Box() Block { return &mkBox{block} }
func (block opDrop) Cat(xs ...Block) Block {
	rest := newCatN(xs...)
	return newCatN(block, rest)
}
func (block opDrop) Reduce(quota int) Block { return block }
func (block opDrop) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(CodeOpDrop)
}
func (block opDrop) String() string { return "drop" }
func (lhs opDrop) Eq(rhs Block) bool {
	_, ok := rhs.(opDrop)
	return ok
}
func (block opDrop) Copy() bool { return true }
func (block opDrop) Drop() bool { return true }
func (block opDrop) Swap() bool { return true }
func (block opDrop) step(ctx *reduce) bool {
	if ctx.arity() == 0 {
		ctx.clear(block)
	}
	lhs := ctx.peek(0)
	if !lhs.Drop() {
		ctx.clear(block)
		return false
	}
	ctx.pop()
	return true
}
