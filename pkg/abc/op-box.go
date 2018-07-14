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

type opBox struct{}

func (tau opBox) Box() Block { return &mkBox{tau} }
func (tau opBox) Cat(xs ...Block) Block {
	rest := newCatN(xs...)
	return newCat(tau, rest)
}
func (tau opBox) Reduce(quota int) Block { return tau }
func (tau opBox) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(byteBox)
}
func (tau opBox) String() string { return "box" }
func (lhs opBox) Eq(rhs Block) bool {
	switch rhs.(type) {
	case opBox:
		return true
	default:
		return false
	}
}
