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

type opApp struct{}

func (block opApp) Box() Block { return &mkBox{block} }
func (block opApp) Cat(xs ...Block) Block {
	rest := newCatN(xs...)
	return newCat(block, rest)
}
func (block opApp) Reduce(quota int) Block { return block }
func (block opApp) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(byteOpApp)
}
func (block opApp) String() string { return "app" }
func (lhs opApp) Eq(rhs Block) bool {
	_, ok := rhs.(opApp)
	return ok
}
func (block opApp) step(ctx *reduce) bool {
	if ctx.arity() == 0 {
		ctx.clear(block)
		return false
	}
	fst, ok := ctx.peek(0).(*mkBox)
	if !ok {
		ctx.clear(block)
		return false
	}
	ctx.pop()
	ctx.queue(fst.body)
	return true
}
