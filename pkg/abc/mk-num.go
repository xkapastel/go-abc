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

type mkNum struct{ value float64 }

// NewNum creates a new real number.
func NewNum(value float64) Block { return mkNum{value} }
func (block mkNum) String() string {
	return fmt.Sprintf("%.10g", block.value)
}
func (lhs mkNum) eq(rhs Block) bool {
	switch rhs := rhs.(type) {
	case mkNum:
		return lhs.value == rhs.value
	default:
		return false
	}
}
func (block mkNum) step(ctx *reduce) bool {
	ctx.data.push(block)
	return false
}