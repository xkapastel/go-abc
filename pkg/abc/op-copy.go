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

type opCopy struct{}

func (block opCopy) Box() Block { return &mkBox{block} }
func (block opCopy) Cat(xs ...Block) Block {
	rest := newCatN(xs...)
	return newCat(block, rest)
}
func (block opCopy) Reduce(quota int) Block { return block }
func (block opCopy) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(CodeOpCopy)
}
func (block opCopy) String() string { return "copy" }
func (lhs opCopy) Eq(rhs Block) bool {
	_, ok := rhs.(opCopy)
	return ok
}
func (block opCopy) Copy() bool { return true }
func (block opCopy) Drop() bool { return true }
func (block opCopy) Swap() bool { return true }
func (block opCopy) step(ctx *reduce) bool {
	if ctx.arity() == 0 {
		ctx.clear(block)
		return false
	}
	lhs := ctx.peek(0)
	if !lhs.Copy() {
		ctx.clear(block)
		return false
	}
	ctx.push(lhs)
	return true
}
