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

func (object opDrop) String() string { return "e" }
func (lhs opDrop) eq(rhs Object) bool {
	_, ok := rhs.(opDrop)
	return ok
}
func (object opDrop) step(ctx *rewrite) bool {
	if ctx.data.len() == 0 {
		ctx.clear(object)
		return false
	}
	ctx.data.pop()
	return true
}
