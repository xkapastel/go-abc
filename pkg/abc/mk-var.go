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

type mkVar struct{ name string }

// NewVar creates a new free variable.
func NewVar(name string) Block { return mkVar{name} }
func (block mkVar) String() string {
	return block.name
}
func (lhs mkVar) eq(rhs Block) bool {
	switch rhs := rhs.(type) {
	case mkVar:
		return lhs.name == rhs.name
	default:
		return false
	}
}
func (block mkVar) step(ctx *reduce) bool {
	body, err := readFile(block.name)
	if err != nil {
		ctx.clear(block)
		return false
	}
	ctx.work.push(body)
	return true
}
