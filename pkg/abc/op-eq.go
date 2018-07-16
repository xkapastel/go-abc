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

type opEq struct{}

func (block opEq) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(CodeOpEq)
}
func (block opEq) String() string { return "eq" }
func (lhs opEq) Eq(rhs Block) bool {
	_, ok := rhs.(opEq)
	return ok
}
func (block opEq) Copy() bool { return true }
func (block opEq) Drop() bool { return true }
func (block opEq) Swap() bool { return true }
func (block opEq) step(ctx *reduce) bool {
	if ctx.arity() < 2 {
		ctx.clear(block)
		return false
	}
	lhs := ctx.peek(0)
	rhs := ctx.peek(1)
	if !lhs.Eq(rhs) {
		ctx.clear(block)
		return false
	}
	return true
}
