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

type opCopy struct{}

func (object opCopy) String() string { return "d" }
func (lhs opCopy) eq(rhs Object) bool {
	_, ok := rhs.(opCopy)
	return ok
}
func (object opCopy) step(ctx *rewrite) bool {
	if ctx.data.len() == 0 {
		ctx.clear(object)
		return false
	}
	lhs := ctx.data.peek(0)
	ctx.data.push(lhs)
	return true
}
