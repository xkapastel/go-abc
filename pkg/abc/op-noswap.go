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

type opNoSwap struct{}

func (block opNoSwap) Box() Block { return &mkBox{block} }
func (block opNoSwap) Cat(xs ...Block) Block {
	rest := newCatN(xs...)
	return newCat(block, rest)
}
func (block opNoSwap) Reduce(quota int) Block { return block }
func (block opNoSwap) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(CodeOpNoSwap)
}
func (block opNoSwap) String() string { return "noswap" }
func (lhs opNoSwap) Eq(rhs Block) bool {
	_, ok := rhs.(opNoSwap)
	return ok
}
func (block opNoSwap) Copy() bool { return true }
func (block opNoSwap) Drop() bool { return true }
func (block opNoSwap) Swap() bool { return false }
func (block opNoSwap) step(ctx *reduce) bool {
	ctx.stash(block)
	return false
}
