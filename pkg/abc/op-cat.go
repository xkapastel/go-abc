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

type opCat struct{}

func (object opCat) String() string { return "c" }
func (lhs opCat) eq(rhs Object) bool {
	_, ok := rhs.(opCat)
	return ok
}
func (object opCat) step(ctx *rewrite) bool {
	if ctx.data.len() < 2 {
		ctx.clear(object)
		return false
	}
	var ok bool
	rhs, ok := ctx.data.peek(0).(*mkBox)
	if !ok {
		ctx.clear(object)
		return false
	}
	lhs, ok := ctx.data.peek(1).(*mkBox)
	if !ok {
		ctx.clear(object)
		return false
	}
	ctx.data.pop()
	ctx.data.pop()
	cat := newCats(lhs.body, rhs.body)
	box := newBox(cat)
	ctx.data.push(box)
	return true
}
