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

type opCat struct{}

func (block opCat) Reduce(quota int) Block { return block }
func (block opCat) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(CodeOpCat)
}
func (block opCat) String() string { return "c" }
func (lhs opCat) Eq(rhs Block) bool {
	_, ok := rhs.(opCat)
	return ok
}
func (block opCat) Copy() bool { return true }
func (block opCat) Drop() bool { return true }
func (block opCat) Swap() bool { return true }
func (block opCat) step(ctx *reduce) bool {
	if ctx.arity() < 2 {
		ctx.clear(block)
		return false
	}
	var ok bool
	rhs, ok := ctx.peek(0).(*mkBox)
	if !ok {
		ctx.clear(block)
		return false
	}
	lhs, ok := ctx.peek(1).(*mkBox)
	if !ok {
		ctx.clear(block)
		return false
	}
	ctx.pop()
	ctx.pop()
	cat := NewCat(lhs.body, rhs.body)
	box := NewBox(cat)
	ctx.push(box)
	return true
}
