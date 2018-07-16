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

type opBox struct{}

func (block opBox) Box() Block { return &mkBox{block} }
func (block opBox) Cat(xs ...Block) Block {
	rest := newCatN(xs...)
	return newCat(block, rest)
}
func (block opBox) Reduce(quota int) Block { return block }
func (block opBox) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(CodeOpBox)
}
func (block opBox) String() string { return "b" }
func (lhs opBox) Eq(rhs Block) bool {
	_, ok := rhs.(opBox)
	return ok
}
func (block opBox) Copy() bool { return true }
func (block opBox) Drop() bool { return true }
func (block opBox) Swap() bool { return true }
func (block opBox) step(ctx *reduce) bool {
	if ctx.arity() == 0 {
		ctx.clear(block)
		return false
	}
	lhs := ctx.pop()
	rhs := lhs.Box()
	ctx.push(rhs)
	return true
}
