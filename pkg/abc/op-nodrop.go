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

type opNoDrop struct{}

func (block opNoDrop) Box() Block { return &mkBox{block} }
func (block opNoDrop) Cat(xs ...Block) Block {
	rest := newCatN(xs...)
	return newCat(block, rest)
}
func (block opNoDrop) Reduce(quota int) Block { return block }
func (block opNoDrop) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(CodeOpNoDrop)
}
func (block opNoDrop) String() string { return "nd" }
func (lhs opNoDrop) Eq(rhs Block) bool {
	_, ok := rhs.(opNoDrop)
	return ok
}
func (block opNoDrop) Copy() bool { return true }
func (block opNoDrop) Drop() bool { return false }
func (block opNoDrop) Swap() bool { return true }
func (block opNoDrop) step(ctx *reduce) bool {
	ctx.stash(block)
	return false
}
