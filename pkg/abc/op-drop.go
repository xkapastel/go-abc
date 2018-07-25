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

import ()

type opDrop struct{}

func (block opDrop) String() string { return "%drop" }
func (lhs opDrop) eq(rhs Block) bool {
	_, ok := rhs.(opDrop)
	return ok
}
func (block opDrop) Copy() bool { return true }
func (block opDrop) Drop() bool { return true }
func (block opDrop) Swap() bool { return true }
func (block opDrop) step(ctx *reduce) bool {
	if ctx.data.len() == 0 {
		ctx.clear(block)
		return false
	}
	lhs := ctx.data.peek(0)
	if !lhs.Drop() {
		ctx.clear(block)
		return false
	}
	ctx.data.pop()
	return true
}
