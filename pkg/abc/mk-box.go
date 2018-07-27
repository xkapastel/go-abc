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
	"fmt"
)

type mkBox struct{ body Object }

func newBox(object Object) Object { return &mkBox{object} }
func (object *mkBox) String() string {
	body := object.body.String()
	return fmt.Sprintf("[%s]", body)
}
func (lhs *mkBox) eq(rhs Object) bool {
	switch rhs := rhs.(type) {
	case *mkBox:
		return lhs.body.eq(rhs.body)
	default:
		return false
	}
}
func (object *mkBox) step(ctx *rewrite) bool {
	ctx.data.push(object)
	return false
}
