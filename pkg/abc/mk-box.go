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
	"io"
)

type mkBox struct{ body Block }

func (tau *mkBox) Box() Block { return &mkBox{tau} }
func (tau *mkBox) Cat(xs ...Block) Block {
	rest := newCatN(xs...)
	return newCat(tau, rest)
}
func (tau *mkBox) Reduce(quota int) Block { return tau }
func (tau *mkBox) Encode(dst io.ByteWriter) error {
	if err := dst.WriteByte(byteBegin); err != nil {
		return err
	}
	if err := tau.body.Encode(dst); err != nil {
		return err
	}
	return dst.WriteByte(byteEnd)
}
func (tau *mkBox) String() string {
	body := tau.body.String()
	return fmt.Sprintf("[%s]", body)
}
func (lhs *mkBox) Eq(rhs Block) bool {
	switch rhs := rhs.(type) {
	case *mkBox:
		return lhs.body.Eq(rhs.body)
	default:
		return false
	}
}
