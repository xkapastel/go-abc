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

type opSwap struct{}

func (block opSwap) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(CodeOpSwap)
}
func (block opSwap) String() string { return "sp" }
func (lhs opSwap) eq(rhs Block) bool {
	_, ok := rhs.(opSwap)
	return ok
}
func (block opSwap) Copy() bool { return true }
func (block opSwap) Drop() bool { return true }
func (block opSwap) Swap() bool { return true }
func (block opSwap) step(ctx *reduce) bool {
	if ctx.arity() < 2 {
		ctx.clear(block)
		return false
	}
	fst := ctx.peek(0)
	snd := ctx.peek(1)
	if !fst.Swap() {
		ctx.clear(block)
		return false
	}
	if !snd.Swap() {
		ctx.clear(block)
		return false
	}
	ctx.pop()
	ctx.pop()
	ctx.push(fst)
	ctx.push(snd)
	return true
}
